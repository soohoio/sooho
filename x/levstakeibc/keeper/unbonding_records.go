package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	"github.com/soohoio/stayking/v2/utils"
	epochstypes "github.com/soohoio/stayking/v2/x/epochs/types"
	lendingpooltypes "github.com/soohoio/stayking/v2/x/lendingpool/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	recordstypes "github.com/soohoio/stayking/v2/x/records/types"
	"github.com/spf13/cast"
)

func (k Keeper) CreateEpochUnbondingRecord(ctx sdk.Context, epochNumber uint64) bool {
	k.Logger(ctx).Info(fmt.Sprintf("Creating Epoch Unbonding Records for Epoch %d", epochNumber))

	hostZoneUnbondings := []*recordstypes.HostZoneUnbonding{}

	for _, hostZone := range k.GetAllHostZone(ctx) {
		k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Creating Epoch Unbonding Record"))

		hostZoneUnbonding := recordstypes.HostZoneUnbonding{
			NativeTokenAmount: sdk.ZeroInt(),
			StTokenAmount:     sdk.ZeroInt(),
			Denom:             hostZone.HostDenom,
			HostZoneId:        hostZone.ChainId,
			Status:            recordstypes.HostZoneUnbonding_UNBONDING_QUEUE,
		}
		hostZoneUnbondings = append(hostZoneUnbondings, &hostZoneUnbonding)
	}

	epochUnbondingRecord := recordstypes.EpochUnbondingRecord{
		EpochNumber:        cast.ToUint64(epochNumber),
		HostZoneUnbondings: hostZoneUnbondings,
	}
	k.RecordsKeeper.SetEpochUnbondingRecord(ctx, epochUnbondingRecord)
	return true
}

func (k Keeper) CleanupEpochUnbondingRecords(ctx sdk.Context) bool {
	k.Logger(ctx).Info("Cleaning Claimed Epoch Unbonding Records...")

	for _, epochUnbondingRecord := range k.RecordsKeeper.GetAllEpochUnbondingRecord(ctx) {
		shouldDeleteEpochUnbondingRecord := true
		hostZoneUnbondings := epochUnbondingRecord.HostZoneUnbondings

		for _, hostZoneUnbonding := range hostZoneUnbondings {
			// if an EpochUnbondingRecord has any HostZoneUnbonding with non-zero balances, we don't delete the EpochUnbondingRecord
			// because it has outstanding tokens that need to be claimed
			if !hostZoneUnbonding.NativeTokenAmount.Equal(sdk.ZeroInt()) {
				shouldDeleteEpochUnbondingRecord = false
				break
			}
		}
		if shouldDeleteEpochUnbondingRecord {
			k.Logger(ctx).Info(fmt.Sprintf("  EpochUnbondingRecord %d - All unbondings claimed, removing record", epochUnbondingRecord.EpochNumber))
			k.RecordsKeeper.RemoveEpochUnbondingRecord(ctx, epochUnbondingRecord.EpochNumber)
		} else {
			k.Logger(ctx).Info(fmt.Sprintf("  EpochUnbondingRecord %d - Has unclaimed unbondings", epochUnbondingRecord.EpochNumber))
		}
	}

	return true
}

func (k Keeper) InitiateAllHostZoneUnbondings(ctx sdk.Context, dayEpochNumber uint64) (bool, []string, []string) {
	k.Logger(ctx).Info(fmt.Sprintf("Initiating all host zone unbondings for epoch %d...", dayEpochNumber))

	success := true
	successfulUnbondings := []string{}
	failedUnbondings := []string{}

	for _, hostZone := range k.GetAllHostZone(ctx) {
		if dayEpochNumber%hostZone.UnbondingFrequency != 0 {
			k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId,
				"Host does not unbond this epoch (Unbonding Frequency: %d, Epoch: %d)", hostZone.UnbondingFrequency, dayEpochNumber))
			continue
		}

		msgs, totalAmountToUnbond, marshalledCallbackArgs, epochUnbondingRecordIds, err := k.GetHostZoneUnbondingMsgs(ctx, hostZone)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error getting unbonding msgs for host zone %s: %s", hostZone.ChainId, err.Error()))
			success = false
			failedUnbondings = append(failedUnbondings, hostZone.ChainId)
			continue
		}

		if totalAmountToUnbond.IsZero() {
			continue
		}

		err = k.SubmitHostZoneUnbondingMsg(ctx, msgs, totalAmountToUnbond, marshalledCallbackArgs, hostZone)
		if err != nil {
			errMsg := fmt.Sprintf("Error submitting unbonding tx for host zone %s: %s", hostZone.ChainId, err.Error())
			k.Logger(ctx).Error(errMsg)
			success = false
			failedUnbondings = append(failedUnbondings, hostZone.ChainId)
			continue
		}

		err = k.RecordsKeeper.SetHostZoneUnbondings(ctx, hostZone.ChainId, epochUnbondingRecordIds, recordstypes.HostZoneUnbonding_UNBONDING_IN_PROGRESS)
		if err != nil {
			k.Logger(ctx).Error(err.Error())
			success = false
			continue
		}
		successfulUnbondings = append(successfulUnbondings, hostZone.ChainId)
	}

	return success, successfulUnbondings, failedUnbondings
}

func (k Keeper) SweepAllUnbondedTokens(ctx sdk.Context) bool {
	k.Logger(ctx).Info("Sweeping All Unbonded Tokens...")

	success := true
	var successfulSweeps []string
	var sweepAmounts []sdk.Int
	var failedSweeps []string
	hostZones := k.GetAllHostZone(ctx)

	epochUnbondingRecords := k.RecordsKeeper.GetAllEpochUnbondingRecord(ctx)
	for _, hostZone := range hostZones {
		hostZoneSuccess, sweepAmount := k.SweepAllUnbondedTokensForHostZone(ctx, hostZone, epochUnbondingRecords)
		if hostZoneSuccess {
			successfulSweeps = append(successfulSweeps, hostZone.ChainId)
			sweepAmounts = append(sweepAmounts, sweepAmount)
		} else {
			success = false
			failedSweeps = append(failedSweeps, hostZone.ChainId)
		}
	}

	return success
}

func (k Keeper) GetHostZoneUnbondingMsgs(ctx sdk.Context, hostZone types.HostZone) ([]sdk.Msg, sdk.Int, []byte, []uint64, error) {
	k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Preparing MsgUndelegates from the delegation account to each validator"))

	// variable initialization..
	var epochUnbondingRecordIds []uint64
	var marshalledCallbackArgs []byte
	var err error
	var msgs []sdk.Msg
	var totalAmountToUnbond = sdk.ZeroInt()

	for _, epochUnbonding := range k.RecordsKeeper.GetAllEpochUnbondingRecord(ctx) {
		hostZoneRecord, found := k.RecordsKeeper.GetHostZoneUnbondingByChainId(ctx, epochUnbonding.EpochNumber, hostZone.ChainId)
		if !found {
			continue
		}
		k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Epoch %d - Status: %s, Amount: %v",
			epochUnbonding.EpochNumber, hostZoneRecord.Status, hostZoneRecord.NativeTokenAmount))

		if hostZoneRecord.Status == recordstypes.HostZoneUnbonding_UNBONDING_QUEUE && hostZoneRecord.NativeTokenAmount.GT(sdk.ZeroInt()) {
			totalAmountToUnbond = totalAmountToUnbond.Add(hostZoneRecord.NativeTokenAmount)
			epochUnbondingRecordIds = append(epochUnbondingRecordIds, epochUnbonding.EpochNumber)
			k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "  %v%s included in total unbonding", hostZoneRecord.NativeTokenAmount, hostZoneRecord.Denom))
		}
	}

	k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Total unbonded amount: %v%s", totalAmountToUnbond, hostZone.HostDenom))

	if totalAmountToUnbond.IsZero() {
		return nil, sdk.ZeroInt(), nil, nil, nil
	}

	// Determine the desired unbonding amount for each validator based on our target weights
	targetUnbondingsByValidator, err := k.GetTargetValAmtsForHostZone(ctx, hostZone, totalAmountToUnbond)
	if err != nil {
		errMsg := fmt.Sprintf("Error getting target val amts for host zone %s %v: %s", hostZone.ChainId, totalAmountToUnbond, err)
		k.Logger(ctx).Error(errMsg)
		return nil, sdk.ZeroInt(), nil, nil, errorsmod.Wrap(types.ErrNoValidatorAmts, errMsg)
	}

	// Check if each validator has enough current delegations to cover the target unbonded amount
	// If it doesn't have enough, update the target to equal their total delegations and record the overflow amount
	finalUnbondingsByValidator := make(map[string]sdk.Int)
	overflowAmount := sdk.ZeroInt()
	for _, validator := range hostZone.Validators {

		targetUnbondAmount := targetUnbondingsByValidator[validator.Address]
		k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId,
			"  Validator %s - Weight: %d, Target Unbond Amount: %v, Current Delegations: %v", validator.Address, validator.Weight, targetUnbondAmount, validator.DelegationAmt))

		// If they don't have enough to cover the unbondings, set their target unbond amount to their current delegations and increment the overflow
		if targetUnbondAmount.GT(validator.DelegationAmt) {
			overflowAmount = overflowAmount.Add(targetUnbondAmount).Sub(validator.DelegationAmt)
			targetUnbondAmount = validator.DelegationAmt
		}
		finalUnbondingsByValidator[validator.Address] = targetUnbondAmount
	}

	// If there was overflow (i.e. there was at least one validator without sufficient delegations to cover their unbondings)
	//  then reallocate across the other validators
	if overflowAmount.GT(sdk.ZeroInt()) {
		k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId,
			"Expected validator undelegation amount on is greater than it's current delegations. Redistributing undelegations accordingly."))

		for _, validator := range hostZone.Validators {
			targetUnbondAmount := finalUnbondingsByValidator[validator.Address]

			// Check if we can unbond more from this validator
			validatorUnbondExtraCapacity := validator.DelegationAmt.Sub(targetUnbondAmount)
			if validatorUnbondExtraCapacity.GT(sdk.ZeroInt()) {

				// If we can fully cover the unbonding, do so with this validator
				if validatorUnbondExtraCapacity.GT(overflowAmount) {
					finalUnbondingsByValidator[validator.Address] = finalUnbondingsByValidator[validator.Address].Add(overflowAmount)
					overflowAmount = sdk.ZeroInt()
					break
				} else {
					// If we can't, cover the unbondings, cover as much as we can and move onto the next validator
					finalUnbondingsByValidator[validator.Address] = finalUnbondingsByValidator[validator.Address].Add(validatorUnbondExtraCapacity)
					overflowAmount = overflowAmount.Sub(validatorUnbondExtraCapacity)
				}
			}
		}
	}

	// If after re-allocating, we still can't cover the overflow, something is very wrong
	// TODO: 밸리데이터한테 사전에 Weight 에 따라 Delegation 하는 것을 재할당 해버리면 Overflow 가 발생하는 것에 대한 부분 같은데 이 부분이 어떻게 발생되는지 자세히 봐야할 듯
	if overflowAmount.GT(sdk.ZeroInt()) {
		errMsg := fmt.Sprintf("Could not unbond %v on Host Zone %s, unable to balance the unbond amount across validators",
			totalAmountToUnbond, hostZone.ChainId)
		k.Logger(ctx).Error(errMsg)
		return nil, sdk.ZeroInt(), nil, nil, errorsmod.Wrap(sdkerrors.ErrNotFound, errMsg)
	}

	delegationAccount := hostZone.DelegationAccount
	if delegationAccount == nil || delegationAccount.Address == "" {
		errMsg := fmt.Sprintf("host zone (chain id %v) is missing a delegation address!", hostZone.ChainId)
		k.Logger(ctx).Error(errMsg)
		return nil, sdk.ZeroInt(), nil, nil, errorsmod.Wrap(types.ErrHostZoneICAAccountNotFound, errMsg)
	}

	// Construct the MsgUndelegate transaction
	var splitDelegations []*types.SplitDelegation
	for _, validatorAddress := range utils.StringMapKeys(finalUnbondingsByValidator) { // DO NOT REMOVE: StringMapKeys fixes non-deterministic map iteration
		undelegationAmount := sdk.NewCoin(hostZone.HostDenom, finalUnbondingsByValidator[validatorAddress])

		// Store the ICA transactions
		msgs = append(msgs, &stakingtypes.MsgUndelegate{
			DelegatorAddress: delegationAccount.Address,
			ValidatorAddress: validatorAddress,
			Amount:           undelegationAmount,
		})

		// Store the split delegations for the callback
		splitDelegations = append(splitDelegations, &types.SplitDelegation{
			Validator: validatorAddress,
			Amount:    undelegationAmount.Amount,
		})
	}

	// Store the callback data
	undelegateCallback := types.UndelegateCallback{
		HostZoneId:              hostZone.ChainId,
		SplitDelegations:        splitDelegations,
		EpochUnbondingRecordIds: epochUnbondingRecordIds,
	}
	k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Marshalling UndelegateCallback args: %+v", undelegateCallback))
	marshalledCallbackArgs, err = k.MarshalUndelegateCallbackArgs(ctx, undelegateCallback)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return nil, sdk.ZeroInt(), nil, nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, err.Error())
	}

	return msgs, totalAmountToUnbond, marshalledCallbackArgs, epochUnbondingRecordIds, nil
}

func (k Keeper) SubmitHostZoneUnbondingMsg(ctx sdk.Context, msgs []sdk.Msg, totalAmtToUnbond sdk.Int, marshalledCallbackArgs []byte, hostZone types.HostZone) error {
	delegationAccount := hostZone.GetDelegationAccount()

	// safety check: if msgs is nil, error
	if msgs == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no msgs to submit for host zone unbondings")
	}

	_, err := k.SubmitTxsEpoch(ctx, hostZone.GetConnectionId(), msgs, *delegationAccount, epochstypes.DAY_EPOCH, ICACallbackID_Undelegate, marshalledCallbackArgs)
	if err != nil {
		errMsg := fmt.Sprintf("Error submitting unbonding tx: %s", err)
		k.Logger(ctx).Error(errMsg)
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, errMsg)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute("hostZone", hostZone.ChainId),
			sdk.NewAttribute("newAmountUnbonding", totalAmtToUnbond.String()),
		),
	)

	return nil
}

func (k Keeper) SweepAllUnbondedTokensForHostZone(ctx sdk.Context, hostZone types.HostZone, epochUnbondingRecords []recordstypes.EpochUnbondingRecord) (success bool, sweepAmount sdk.Int) {
	k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Sweeping unbonded tokens"))

	totalAmtTransferToRedemptionAcct := sdk.ZeroInt()
	var epochUnbondingRecordIds []uint64

	for _, epochUnbondingRecord := range epochUnbondingRecords {

		// Get all the unbondings associated with the epoch + host zone pair
		hostZoneUnbonding, found := k.RecordsKeeper.GetHostZoneUnbondingByChainId(ctx, epochUnbondingRecord.EpochNumber, hostZone.ChainId)
		if !found {
			continue
		}

		// Get latest blockTime from light client
		blockTime, err := k.GetLightClientTimeSafely(ctx, hostZone.ConnectionId)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("\tCould not find blockTime for host zone %s", hostZone.ChainId))
			continue
		}

		k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Epoch %d - Status: %s, Amount: %v, Unbonding Time: %d, Block Time: %d",
			epochUnbondingRecord.EpochNumber, hostZoneUnbonding.Status.String(), hostZoneUnbonding.NativeTokenAmount, hostZoneUnbonding.UnbondingTime, blockTime))

		inTransferQueue := hostZoneUnbonding.Status == recordstypes.HostZoneUnbonding_EXIT_TRANSFER_QUEUE
		validUnbondingTime := hostZoneUnbonding.UnbondingTime > 0 && hostZoneUnbonding.UnbondingTime < blockTime
		if inTransferQueue && validUnbondingTime {
			k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "  %v%s included in sweep", hostZoneUnbonding.NativeTokenAmount, hostZoneUnbonding.Denom))

			totalAmtTransferToRedemptionAcct = totalAmtTransferToRedemptionAcct.Add(hostZoneUnbonding.NativeTokenAmount)
			epochUnbondingRecordIds = append(epochUnbondingRecordIds, epochUnbondingRecord.EpochNumber)
		}
	}

	// If we have any amount to sweep, then we can send the ICA call to sweep them
	if totalAmtTransferToRedemptionAcct.LTE(sdk.ZeroInt()) {
		k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "No tokens ready for sweep"))
		return true, totalAmtTransferToRedemptionAcct
	}
	k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Batch transferring %v to host zone", totalAmtTransferToRedemptionAcct))

	// Get the delegation account and redemption account
	delegationAccount := hostZone.DelegationAccount
	if delegationAccount == nil || delegationAccount.Address == "" {
		k.Logger(ctx).Error(fmt.Sprintf("Zone %s is missing a delegation address!", hostZone.ChainId))
		return false, sdk.ZeroInt()
	}

	transferCoin := sdk.NewCoin(hostZone.HostDenom, totalAmtTransferToRedemptionAcct)
	timeoutTimestamp := uint64(ctx.BlockTime().UnixNano()) + k.GetParam(ctx, types.KeyIBCTransferTimeoutNanos)
	channelEnd, found := k.IBCKeeper.ChannelKeeper.GetChannel(ctx, ibctransfertypes.PortID, hostZone.TransferChannelId)
	if !found {
		errMsg := fmt.Sprintf("invalid channel id, %s not found", hostZone.TransferChannelId)
		k.Logger(ctx).Error(errMsg)
		return false, sdk.ZeroInt()
	}
	k.Logger(ctx).Info(fmt.Sprintf("[DEBUG] channelEnd.Counterparty.ChannelId : %v", channelEnd.Counterparty.ChannelId))
	k.Logger(ctx).Info(fmt.Sprintf("[DEBUG] hostZoneTransferChannelId : %v", hostZone.TransferChannelId))
	msg := ibctransfertypes.NewMsgTransfer(ibctransfertypes.PortID, channelEnd.Counterparty.ChannelId, transferCoin, delegationAccount.Address, hostZone.Address, clienttypes.Height{}, timeoutTimestamp)
	//msg := ibctransfertypes.NewMsgTransfer(ibctransfertypes.PortID, hostZone.TransferChannelId, transferCoin, delegationAccount.Address, hostZone.Address, clienttypes.Height{}, timeoutTimestamp)
	msgs := []sdk.Msg{msg}

	transferCallback := types.TransferUndelegatedTokensCallback{
		EpochUnbondingRecordIds: epochUnbondingRecordIds,
		HostZoneId:              hostZone.ChainId,
	}

	marshalledCallbackArgs, err := k.MarshalTransferUndelegatedTokensArgs(ctx, transferCallback)
	if err != nil {
		return false, sdk.ZeroInt()
	}

	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return false, sdk.ZeroInt()
	}
	_, err = k.SubmitTxsEpoch(ctx, hostZone.ConnectionId, msgs, *delegationAccount, epochstypes.DAY_EPOCH, ICACallbackID_TransferUndelegatedTokens, marshalledCallbackArgs)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Failed to SubmitTxs, transfer to redemption account on %s", hostZone.ChainId))
		return false, sdk.ZeroInt()
	}
	k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "ICA MsgSend Successfully Sent"))

	err = k.RecordsKeeper.SetHostZoneUnbondings(ctx, hostZone.ChainId, epochUnbondingRecordIds, recordstypes.HostZoneUnbonding_EXIT_TRANSFER_IN_PROGRESS)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return false, sdk.ZeroInt()
	}

	return true, totalAmtTransferToRedemptionAcct
}

func (k Keeper) UnStakeWithLeverage(ctx sdk.Context, _sender string, positionId uint64, chainId, receiver string) error {

	sender, _ := sdk.AccAddressFromBech32(_sender)
	hostZone, found := k.GetHostZone(ctx, chainId)
	if !found {
		return errorsmod.Wrapf(types.ErrHostZoneNotFound, "host zone not found by chain id %s", chainId)
	}

	epochTracker, found := k.GetEpochTracker(ctx, epochstypes.DAY_EPOCH)
	if !found {
		return errorsmod.Wrapf(types.ErrEpochNotFound, "epoch tracker (%s) not found", epochstypes.DAY_EPOCH)
	}
	redemptionId := recordstypes.UserRedemptionRecordKeyFormatter(hostZone.ChainId, epochTracker.EpochNumber, sender.String())

	_, found = k.RecordsKeeper.GetUserRedemptionRecord(ctx, redemptionId)
	if found {
		return errorsmod.Wrapf(recordstypes.ErrRedemptionAlreadyExists, "user already redeemed this epoch: %s", redemptionId)
	}

	position, found := k.GetPosition(ctx, positionId)
	if !found {
		return errorsmod.Wrapf(types.ErrPositionNotFound, "position not found by position id %s", positionId)
	}

	stDenom := types.StAssetDenomFromHostZoneDenom(hostZone.HostDenom)

	// @TODO how to check module account's balance

	//convert to native Token Amount
	nativeTokenAmount := sdk.NewDecFromInt(position.StTokenAmount).Mul(hostZone.RedemptionRate).RoundInt()
	if !nativeTokenAmount.IsPositive() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "amount must be greater than 0. found: %v", position.StTokenAmount)
	}

	// Check HostZone Balance
	if nativeTokenAmount.GT(hostZone.StakedBal) {
		return errorsmod.Wrapf(types.ErrInsufficientFundsOnHostZone,
			"unstaking amount requested is not allowed to greater than staked balance on the host zone (stToken: %v, swappedNativeToken : %v, redemptionRate : %v",
			position.StTokenAmount, nativeTokenAmount, hostZone.RedemptionRate)
	}

	k.Logger(ctx).Info(fmt.Sprintf("position record stDenom amount: %v%s", position.StTokenAmount, stDenom))

	// userRedemption Record 생성
	userRedemptionRecord := recordstypes.UserRedemptionRecord{
		Id:             redemptionId,
		Sender:         sender.String(),
		Receiver:       receiver,
		Amount:         nativeTokenAmount,
		Denom:          hostZone.HostDenom,
		HostZoneId:     hostZone.ChainId,
		EpochNumber:    epochTracker.EpochNumber,
		ClaimIsPending: false,
		PositionId:     positionId,
	}

	epochUnbondingRecord, found := k.RecordsKeeper.GetEpochUnbondingRecord(ctx, epochTracker.EpochNumber)
	if !found {
		k.Logger(ctx).Error("latest epoch unbonding record not found")
		return errorsmod.Wrapf(recordstypes.ErrEpochUnbondingRecordNotFound, "latest epoch unbonding record not found")
	}

	hostZoneUnbonding, found := k.RecordsKeeper.GetHostZoneUnbondingByChainId(ctx, epochUnbondingRecord.EpochNumber, hostZone.ChainId)
	if !found {
		return errorsmod.Wrapf(recordstypes.ErrEpochUnbondingRecordNotFound, "not found unbondings on host zone by chain id %v and epoch number %v", hostZone.ChainId, epochUnbondingRecord.EpochNumber)
	}

	hostZoneUnbonding.NativeTokenAmount = hostZoneUnbonding.NativeTokenAmount.Add(nativeTokenAmount)
	hostZoneUnbonding.UserRedemptionRecords = append(hostZoneUnbonding.UserRedemptionRecords, userRedemptionRecord.Id)

	//Unbonding에 StTokenAmount 추가
	hostZoneUnbonding.StTokenAmount = hostZoneUnbonding.StTokenAmount.Add(position.StTokenAmount)

	k.RecordsKeeper.SetUserRedemptionRecord(ctx, userRedemptionRecord)

	hostZoneUnbondings := epochUnbondingRecord.GetHostZoneUnbondings()
	if hostZoneUnbondings == nil {
		hostZoneUnbondings = []*recordstypes.HostZoneUnbonding{}
		epochUnbondingRecord.HostZoneUnbondings = hostZoneUnbondings
	}
	updatedEpochUnbondingRecord, success := k.RecordsKeeper.AddHostZoneToEpochUnbondingRecord(ctx, epochUnbondingRecord.EpochNumber, hostZone.ChainId, hostZoneUnbonding)
	if !success {
		return errorsmod.Wrapf(types.ErrFailureUpdateUnbondingRecord, "Failed to set host zone epoch unbonding record: epochNumber %d, chainId %s, hostZoneUnbonding %v", epochUnbondingRecord.EpochNumber, hostZone.ChainId, hostZoneUnbonding)
	}
	k.RecordsKeeper.SetEpochUnbondingRecord(ctx, *updatedEpochUnbondingRecord)

	//update Position with native amount
	position.NativeTokenAmount = nativeTokenAmount
	position.Status = types.PositionStatus_POSITION_UNBONDING_IN_PROGRESS
	k.SetPosition(ctx, position)

	k.Logger(ctx).Info(fmt.Sprintf("position updated with native token amount %v", position.NativeTokenAmount))
	return nil
}

func (k Keeper) ReleaseUnbondedAsset(ctx sdk.Context) error {
	k.Logger(ctx).Info(fmt.Sprintf("Release Unbonded Assets..."))
	epochUnbondingRecords := k.RecordsKeeper.GetAllEpochUnbondingRecord(ctx)
	hostZones := k.GetAllHostZone(ctx)
	for _, epochUnbondingRecord := range epochUnbondingRecords {
		for _, hostZone := range hostZones {
			hostZoneUnbonding, found := k.RecordsKeeper.GetHostZoneUnbondingByChainId(ctx, epochUnbondingRecord.EpochNumber, hostZone.ChainId)
			if !found {
				continue
			}
			inReleaseAssetQueue := hostZoneUnbonding.Status == recordstypes.HostZoneUnbonding_RELEASE_ASSET_QUEUE
			if !inReleaseAssetQueue {
				k.Logger(ctx).Info(fmt.Sprintf("no unbonded tokens ready to release..."))
				continue

			} else {
				k.Logger(ctx).Info(fmt.Sprintf("found unbonded token to release for hostzone : %v", hostZone.ChainId))
				zoneAddress, err := sdk.AccAddressFromBech32(hostZone.Address)
				if err != nil {
					return fmt.Errorf("could not bech32 decode address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
				}
				//only if hostzone unbonding status is release asset queue
				k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Epoch %d - Status: %s, Amount: %v",
					epochUnbondingRecord.EpochNumber, hostZoneUnbonding.Status.String(), hostZoneUnbonding.NativeTokenAmount))
				userRedemptionRecords := hostZoneUnbonding.GetUserRedemptionRecords()
				for _, userRedemptionRecordId := range userRedemptionRecords {
					userRedemptionRecord, found := k.RecordsKeeper.GetUserRedemptionRecord(ctx, userRedemptionRecordId)
					if !found {
						continue
					}
					position, found := k.GetPositionByDenomAndSender(ctx, userRedemptionRecord.Denom, userRedemptionRecord.Sender)
					// leverage case
					if found {
						k.Logger(ctx).Info(fmt.Sprintf("[Release Unbonded Asset] position found for userRedemptionRecord id %v", userRedemptionRecordId))
						if position.Status == types.PositionStatus_POSITION_UNBONDING_IN_PROGRESS {
							k.Logger(ctx).Info(fmt.Sprintf("Transfer dept token to lending pool module with position Id %v", position.Id))
							transferCoinToModule := sdk.Coins{sdk.NewCoin(hostZone.IbcDenom, userRedemptionRecord.Amount)}

							err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, zoneAddress, lendingpooltypes.ModuleName, transferCoinToModule)
							if err != nil {
								k.Logger(ctx).Error(fmt.Sprintf("[Release Unbonded Asset Error] Send Coins to lending pool module from zone Address %v, with amount %v", zoneAddress.String(), transferCoinToModule))
								return errorsmod.Wrap(err, "failed to send tokens from zoneAddress to lendingpool module")
							}
							k.Logger(ctx).Info(fmt.Sprintf("Transfer coin to lendingpool module with amount :%v", transferCoinToModule))
							_, err = k.LendingPoolKeeper.Repay(ctx, position.LoanId, transferCoinToModule)
							if err != nil {
								k.Logger(ctx).Error(fmt.Sprintf("Repay failed for loan id %v", position.LoanId))
							}
							k.RemovePosition(ctx, position.Id)
						}
					} else {
						k.Logger(ctx).Info(fmt.Sprintf("[Release Unbonded Asset] position not found for userRedemptionRecord id %v", userRedemptionRecordId))
						userRedemptionRecordSender, err := sdk.AccAddressFromBech32(userRedemptionRecord.Sender)
						if err != nil {
							return fmt.Errorf("could not bech32 decode address %s of useRedemptionRecord with id: %s", userRedemptionRecord.Sender, userRedemptionRecord.Id)
						}
						zoneAddressBalance := k.bankKeeper.GetBalance(ctx, zoneAddress, hostZone.IbcDenom)
						k.Logger(ctx).Info(fmt.Sprintf("[CUSTOM DEBUG] zoneAddressBalance:%v", zoneAddressBalance.Amount))
						if zoneAddressBalance.Amount.LT(userRedemptionRecord.Amount) {
							continue
						}
						transferCoinToModule := sdk.Coins{sdk.NewCoin(hostZone.IbcDenom, userRedemptionRecord.Amount)}
						err = k.bankKeeper.SendCoins(ctx, zoneAddress, userRedemptionRecordSender, transferCoinToModule)
						if err != nil {
							return fmt.Errorf("[Release Unbonded Asset] Send Coins to user address %v. from zone Address %v, with amount %v err: %v", userRedemptionRecordSender, zoneAddress, transferCoinToModule, err.Error())
						}
						k.Logger(ctx).Info(fmt.Sprintf("Transfer coin to user with userRedemptionRecord :%v", userRedemptionRecord.Id))
					}
					k.DecrementHostZoneUnbondingAmount(ctx, userRedemptionRecord, hostZone.ChainId)
					k.RecordsKeeper.RemoveUserRedemptionRecord(ctx, userRedemptionRecordId)
				}
			}

		}

	}

	return nil
}
