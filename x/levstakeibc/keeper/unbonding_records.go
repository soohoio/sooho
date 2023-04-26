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
		errMsg := fmt.Sprintf("Zone %s is missing a delegation address!", hostZone.ChainId)
		k.Logger(ctx).Error(errMsg)
		return nil, sdk.ZeroInt(), nil, nil, sdkerrors.Wrap(types.ErrHostZoneICAAccountNotFound, errMsg)
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
	//msg := ibctransfertypes.NewMsgTransfer(ibctransfertypes.PortID, channelEnd.Counterparty.ChannelId, transferCoin, delegationAccount.Address, hostZone.Address, clienttypes.Height{}, timeoutTimestamp)
	msg := ibctransfertypes.NewMsgTransfer(ibctransfertypes.PortID, hostZone.TransferChannelId, transferCoin, delegationAccount.Address, hostZone.Address, clienttypes.Height{}, timeoutTimestamp)
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

	//redemptionAccount := hostZone.RedemptionAccount
	//if redemptionAccount == nil || redemptionAccount.Address == "" {
	//	k.Logger(ctx).Error(fmt.Sprintf("Zone %s is missing a redemption address!", hostZone.ChainId))
	//	return false, sdk.ZeroInt()
	//}
	//positions := k.GetAllPosition(ctx)
	//if totalAmtTransferToRedemptionAcct.GT(sdk.ZeroInt()) {
	//	for _, position := range positions {
	//		//position Status 확인
	//		if position.Status == types.PositionStatus_POSITION_UNBONDING_IN_PROGRESS {
	//			//unbonding status일 경우 native token amount만큼 repay 함수 호출 with loan id
	//			//err := k.TransferUndelegatedTokensToHostZoneModule(ctx, hostZone, position, delegationAccount.Address)
	//			//if err != nil {
	//			//	k.Logger(ctx).Error(fmt.Sprintf("[Error] Transfer Undelegated Tokens to hoszone Address%v", delegationAccount.Address))
	//			//}
	//			k.Logger(ctx).Info("Transferring NativeTokens from hostzone to stayking")
	//			//zoneAddress, err := sdk.AccAddressFromBech32(hostZone.Address)
	//			//if err != nil {
	//			//	return fmt.Errorf("could not bech32 decode address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	//			//}
	//			transferCoin := sdk.NewCoin(hostZone.HostDenom, position.NativeTokenAmount)
	//			timeoutTimestamp := uint64(ctx.BlockTime().UnixNano()) + k.GetParam(ctx, types.KeyIBCTransferTimeoutNanos)
	//			channelEnd, found := k.IBCKeeper.ChannelKeeper.GetChannel(ctx, ibctransfertypes.PortID, hostZone.TransferChannelId)
	//			if !found {
	//				errMsg := fmt.Sprintf("invalid channel id, %s not found", hostZone.TransferChannelId)
	//				k.Logger(ctx).Error(errMsg)
	//				return false, sdk.ZeroInt()
	//			}
	//			msg := ibctransfertypes.NewMsgTransfer(ibctransfertypes.PortID, channelEnd.Counterparty.ChannelId, transferCoin, delegationAccount.Address, hostZone.Address, clienttypes.Height{}, timeoutTimestamp)
	//			msgs := []sdk.Msg{msg}
	//
	//			transferCallback := types.TransferUndelegatedTokensCallback{
	//				PositionId: position.Id,
	//				HostZoneId: hostZone.ChainId,
	//			}
	//			marshalledCallbackArgs, err := k.MarshalTransferUndelegatedTokensArgs(ctx, transferCallback)
	//			if err != nil {
	//				return false, sdk.ZeroInt()
	//			}
	//
	//			if err != nil {
	//				k.Logger(ctx).Error(err.Error())
	//				return false, sdk.ZeroInt()
	//			}
	//			_, err = k.SubmitTxsEpoch(ctx, hostZone.ConnectionId, msgs, *delegationAccount, epochstypes.DAY_EPOCH, ICACallbackID_TransferUndelegatedTokens, marshalledCallbackArgs)
	//			if err != nil {
	//				k.Logger(ctx).Error(fmt.Sprintf("Failed to SubmitTxs, transfer to redemption account on %s", hostZone.ChainId))
	//				return false, sdk.ZeroInt()
	//			}
	//			k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "ICA MsgSend Successfully Sent"))
	//
	//			position.Status = types.PositionStatus_POSITION_TRANSFER_IN_PROGRESS
	//			k.SetPosition(ctx, position)
	//			//err := k.Transfer(ctx, msg, position, hostZone)
	//
	//			//totalAmountTransferToRedemptionAcct - native token amount
	//			totalAmtTransferToRedemptionAcct = totalAmtTransferToRedemptionAcct.Sub(position.NativeTokenAmount)
	//			if totalAmtTransferToRedemptionAcct.LT(sdk.ZeroInt()) {
	//				k.Logger(ctx).Error(fmt.Sprintf("totalAmtTransferTo Redemption Acct total: %v, position native tokem amt: %v", totalAmtTransferToRedemptionAcct, position.NativeTokenAmount))
	//				return false, totalAmtTransferToRedemptionAcct
	//			}
	//			//position 저장
	//			//k.RemovePosition(ctx, position.Id)
	//		}
	//	}
	//
	//}

	// Build transfer message to transfer from the delegation account to redemption account
	//msgs := []sdk.Msg{
	//	&banktypes.MsgSend{
	//		FromAddress: delegationAccount.Address,
	//		ToAddress:   redemptionAccount.Address,
	//		Amount:      sdk.NewCoins(sdk.NewCoin(hostZone.HostDenom, totalAmtTransferToRedemptionAcct)),
	//	},
	//}
	//k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Preparing MsgSend from Delegation Account to Redemption Account"))
	//
	//// Store the epoch numbers in the callback to identify the epoch unbonding records
	//redemptionCallback := types.RedemptionCallback{
	//	HostZoneId:              hostZone.ChainId,
	//	EpochUnbondingRecordIds: epochUnbondingRecordIds,
	//}
	//marshalledCallbackArgs, err := k.MarshalRedemptionCallbackArgs(ctx, redemptionCallback)
	//if err != nil {
	//	k.Logger(ctx).Error(err.Error())
	//	return false, sdk.ZeroInt()
	//}
	//
	//// Send the transfer ICA
	//_, err = k.SubmitTxsEpoch(ctx, hostZone.ConnectionId, msgs, *delegationAccount, epochstypes.DAY_EPOCH, ICACallbackID_Redemption, marshalledCallbackArgs)
	//if err != nil {
	//	k.Logger(ctx).Error(fmt.Sprintf("Failed to SubmitTxs, transfer to redemption account on %s", hostZone.ChainId))
	//	return false, sdk.ZeroInt()
	//}
	//k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "ICA MsgSend Successfully Sent"))
	//
	// Update the host zone unbonding records to status IN_PROGRESS
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
		return errorsmod.Wrapf(types.ErrInvalidHostZone, "host zone(chainId) is invalid: %s", chainId)
	}

	epochTracker, found := k.GetEpochTracker(ctx, "day")
	if !found {
		return errorsmod.Wrapf(types.ErrEpochNotFound, "epoch tracker found: %s", "day")
	}
	redemptionId := recordstypes.UserRedemptionRecordKeyFormatter(hostZone.ChainId, epochTracker.EpochNumber, sender.String())

	_, found = k.RecordsKeeper.GetUserRedemptionRecord(ctx, redemptionId)
	if found {
		return errorsmod.Wrapf(recordstypes.ErrRedemptionAlreadyExists, "user already redeemed this epoch: %s", redemptionId)
	}

	position, found := k.GetPosition(ctx, positionId)
	if !found {
		return errorsmod.Wrapf(types.ErrPositionNotFound, "position not found: %s", positionId)

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
		return errorsmod.Wrapf(types.ErrInvalidAmount, "cannot unstake an amount g.t. staked balance on host zone: %v", position.StTokenAmount)
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
	}

	epochUnbondingRecord, found := k.RecordsKeeper.GetEpochUnbondingRecord(ctx, epochTracker.EpochNumber)
	if !found {
		k.Logger(ctx).Error("latest epoch unbonding record not found")
		return errorsmod.Wrapf(recordstypes.ErrEpochUnbondingRecordNotFound, "latest epoch unbonding record not found")
	}

	hostZoneUnbonding, found := k.RecordsKeeper.GetHostZoneUnbondingByChainId(ctx, epochUnbondingRecord.EpochNumber, hostZone.ChainId)
	if !found {
		return errorsmod.Wrapf(types.ErrInvalidHostZone, "host zone not found in unbondings: %s", hostZone.ChainId)
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
		k.Logger(ctx).Error(fmt.Sprintf("Failed to set host zone epoch unbonding record: epochNumber %d, chainId %s, hostZoneUnbonding %v", epochUnbondingRecord.EpochNumber, hostZone.ChainId, hostZoneUnbonding))
		return errorsmod.Wrapf(types.ErrEpochNotFound, "couldn't set host zone epoch unbonding record.")
	}
	k.RecordsKeeper.SetEpochUnbondingRecord(ctx, *updatedEpochUnbondingRecord)

	//update Position with native amount
	position.NativeTokenAmount = nativeTokenAmount
	position.Status = types.PositionStatus_POSITION_UNBONDING_IN_PROGRESS
	k.SetPosition(ctx, position)

	k.Logger(ctx).Info(fmt.Sprintf("position updated with native token amount %v", position.NativeTokenAmount))

	return nil
}

// TransferUndelegatedTokensToHostZoneAddress
//func (k Keeper) TransferUndelegatedTokensToHostZoneModule(ctx sdk.Context, hostZone types.HostZone, position types.Position, delegationAddress string) error {
//	k.Logger(ctx).Info("Transferring NativeTokens from hostzone to stayking")
//	//zoneAddress, err := sdk.AccAddressFromBech32(hostZone.Address)
//	//if err != nil {
//	//	return fmt.Errorf("could not bech32 decode address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
//	//}
//	transferCoin := sdk.NewCoin(hostZone.HostDenom, position.NativeTokenAmount)
//	timeoutTimestamp := uint64(ctx.BlockTime().UnixNano()) + k.GetParam(ctx, types.KeyIBCTransferTimeoutNanos)
//	channelEnd, found := k.IBCKeeper.ChannelKeeper.GetChannel(ctx, ibctransfertypes.PortID, hostZone.TransferChannelId)
//	if !found {
//		errMsg := fmt.Sprintf("invalid channel id, %s not found", hostZone.TransferChannelId)
//		k.Logger(ctx).Error(errMsg)
//		return errorsmod.Wrapf(types.ErrChannelNotFound, errMsg)
//	}
//	msg := ibctransfertypes.NewMsgTransfer(ibctransfertypes.PortID, channelEnd.Counterparty.ChannelId, transferCoin, delegationAddress, hostZone.Address, clienttypes.Height{}, timeoutTimestamp)
//	err := k.Transfer(ctx, msg, position, hostZone)
//
//	if err != nil {
//		k.Logger(ctx).Error(fmt.Sprintf("[TransferUndelegatedTokensToHostZoneModule] Failed to initiate IBC transfer to hostzone Module Addres, HostZone: %v, Channel: %v, Amount: %v, ModuleAddress: %v, DelegateAddress: %v, Timeout: %v",
//			hostZone.ChainId, hostZone.TransferChannelId, transferCoin, hostZone.Address, delegationAddress, timeoutTimestamp))
//		k.Logger(ctx).Error(fmt.Sprintf("[TransferUndelegatedTokensToHostZoneModule] err {%s}", err.Error()))
//	}
//
//	return nil
//}
//
//func (k Keeper) Transfer(ctx sdk.Context, msg *ibctransfertypes.MsgTransfer, position types.Position, hostZone types.HostZone) error {
//	goCtx := sdk.WrapSDKContext(ctx)
//	msgTransferResponse, err := k.TransferKeeper.Transfer(goCtx, msg)
//	if err != nil {
//		return err
//	}
//	sequence := msgTransferResponse.Sequence
//
//	transferCallback := types.TransferUndelegatedTokensCallback{
//		EpochUnbondingRecordIds: position.Id,
//		HostZoneId:              hostZone.ChainId,
//	}
//	marshalledCallbackArgs, err := k.MarshalTransferUndelegatedTokensArgs(ctx, transferCallback)
//	if err != nil {
//		return err
//	}
//
//	callback := icacallbackstypes.CallbackData{
//		CallbackKey:  icacallbackstypes.PacketID(msg.SourcePort, msg.SourceChannel, sequence),
//		PortId:       msg.SourcePort,
//		ChannelId:    msg.SourceChannel,
//		Sequence:     sequence,
//		CallbackId:   ICACallbackID_TransferUndelegatedTokens,
//		CallbackArgs: marshalledCallbackArgs,
//	}
//	k.ICACallbacksKeeper.SetCallbackData(ctx, callback)
//	position.Status = types.PositionStatus_POSITION_TRANSFER_IN_PROGRESS
//
//	return nil
//}

//hostZone Found
//transferCoin = NewCoin(Denom, NativeTokenAmount)

//msg := ibctypes.NewMsgTransfer(PortId, hostZon.TransferChannelId

//Transfer
//msgTransferResponse
