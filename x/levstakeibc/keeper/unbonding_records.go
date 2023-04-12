package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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
