package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	ibctypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	"github.com/soohoio/stayking/v3/utils"
	epochstypes "github.com/soohoio/stayking/v3/x/epochs/types"
	"github.com/soohoio/stayking/v3/x/levstakeibc/types"
	recordstypes "github.com/soohoio/stayking/v3/x/records/types"
	"github.com/spf13/cast"
)

// Create a new deposit record for each host zone for the given epoch
func (k Keeper) CreateDepositRecordsForEpoch(ctx sdk.Context, epochNumber uint64) {
	k.Logger(ctx).Info(fmt.Sprintf("Creating Deposit Records for Epoch %d", epochNumber))

	for _, hostZone := range k.GetAllHostZone(ctx) {
		k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Creating Deposit Record"))

		depositRecord := recordstypes.DepositRecord{
			Amount:             sdk.ZeroInt(),
			Denom:              hostZone.HostDenom,
			HostZoneId:         hostZone.ChainId,
			Status:             recordstypes.DepositRecord_TRANSFER_QUEUE,
			DepositEpochNumber: epochNumber,
		}
		k.RecordsKeeper.AppendDepositRecord(ctx, depositRecord)
	}
}

func (k Keeper) filterDepositRecords(epochNumber uint64, status recordstypes.DepositRecord_Status, depositRecords []recordstypes.DepositRecord) []recordstypes.DepositRecord {
	return utils.FilterDepositRecords(depositRecords, func(record recordstypes.DepositRecord) (condition bool) {
		isTransferRecord := record.Status == status
		isBeforeCurrentEpoch := record.DepositEpochNumber < epochNumber
		return isTransferRecord && isBeforeCurrentEpoch
	})
}

func (k Keeper) TransferExistingDepositsToHostZones(ctx sdk.Context, epochNumber uint64, depositRecords []recordstypes.DepositRecord) {
	k.Logger(ctx).Info("Transferring ibc token from Stayking to hostzone")

	transferDepositRecords := k.filterDepositRecords(epochNumber, recordstypes.DepositRecord_TRANSFER_QUEUE, depositRecords)

	for _, depositRecord := range transferDepositRecords {
		k.Logger(ctx).Info(utils.LogWithHostZone(depositRecord.HostZoneId,
			"Processing deposit record %d: %v%s", depositRecord.Id, depositRecord.Amount, depositRecord.Denom))

		// if a TRANSFER_QUEUE record has 0 balance and was created in the previous epoch, it's safe to remove since it will never be updated or used
		if depositRecord.Amount.LTE(sdk.ZeroInt()) && depositRecord.DepositEpochNumber < epochNumber {
			k.Logger(ctx).Info(utils.LogWithHostZone(depositRecord.HostZoneId, "Empty deposit record - Removing."))
			k.RecordsKeeper.RemoveDepositRecord(ctx, depositRecord.Id)
			continue
		}

		hostZone, hostZoneFound := k.GetHostZone(ctx, depositRecord.HostZoneId)
		if !hostZoneFound {
			k.Logger(ctx).Error(fmt.Sprintf("[TransferExistingDepositsToHostZones] Host zone not found for deposit record id %d", depositRecord.Id))
			continue
		}

		hostZoneModuleAddress := hostZone.Address
		delegateAccount := hostZone.DelegationAccount
		if delegateAccount == nil || delegateAccount.Address == "" {
			k.Logger(ctx).Error(fmt.Sprintf("[TransferExistingDepositsToHostZones] Zone %s is missing a delegation address!", hostZone.ChainId))
			continue
		}
		delegateAddress := delegateAccount.Address

		k.Logger(ctx).Info(utils.LogWithHostZone(depositRecord.HostZoneId, "Transferring %v%s", depositRecord.Amount, hostZone.HostDenom))
		transferCoin := sdk.NewCoin(hostZone.IbcDenom, depositRecord.Amount)

		// timeout 30 min in the future
		// NOTE: this assumes no clock drift between chains, which tendermint guarantees
		// if we onboard non-tendermint chains, we need to use the time on the host chain to
		// calculate the timeout
		// https://github.com/tendermint/tendermint/blob/v0.34.x/spec/consensus/bft-time.md
		timeoutTimestamp := uint64(ctx.BlockTime().UnixNano()) + k.GetParam(ctx, types.KeyIBCTransferTimeoutNanos)
		msg := ibctypes.NewMsgTransfer(ibctransfertypes.PortID, hostZone.TransferChannelId, transferCoin, hostZoneModuleAddress, delegateAddress, clienttypes.Height{}, timeoutTimestamp)
		k.Logger(ctx).Info(utils.LogWithHostZone(depositRecord.HostZoneId, "Transfer Msg: %+v", msg))

		// transfer the deposit record and update its status to TRANSFER_IN_PROGRESS
		err := k.RecordsKeeper.Transfer(ctx, msg, depositRecord)

		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("[TransferExistingDepositsToHostZones] Failed to initiate IBC transfer to host zone, HostZone: %v, Channel: %v, Amount: %v, ModuleAddress: %v, DelegateAddress: %v, Timeout: %v",
				hostZone.ChainId, hostZone.TransferChannelId, transferCoin, hostZoneModuleAddress, delegateAddress, timeoutTimestamp))
			k.Logger(ctx).Error(fmt.Sprintf("[TransferExistingDepositsToHostZones] err {%s}", err.Error()))
			continue
		}

		k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Successfully submitted transfer"))
	}
}

func (k Keeper) StakeExistingDepositsOnHostZones(ctx sdk.Context, epochNumber uint64, depositRecords []recordstypes.DepositRecord) {
	k.Logger(ctx).Info("Staking deposit records...")

	stakingDepositRecords := k.filterDepositRecords(epochNumber, recordstypes.DepositRecord_DELEGATION_QUEUE, depositRecords)

	if len(stakingDepositRecords) == 0 {
		k.Logger(ctx).Info("No deposit records in state DELEGATION_QUEUE")
		return
	}

	maxDepositRecordsToStake := utils.Min(len(stakingDepositRecords), cast.ToInt(k.GetParam(ctx, types.KeyMaxStakeICACallsPerEpoch)))
	if maxDepositRecordsToStake < len(stakingDepositRecords) {
		k.Logger(ctx).Info(fmt.Sprintf("  MaxStakeICACallsPerEpoch limit reached - Only staking %d out of %d deposit records", maxDepositRecordsToStake, len(stakingDepositRecords)))
	}

	for _, depositRecord := range stakingDepositRecords[:maxDepositRecordsToStake] {
		k.Logger(ctx).Info(utils.LogWithHostZone(depositRecord.HostZoneId,
			"Processing deposit record %d: %v%s", depositRecord.Id, depositRecord.Amount, depositRecord.Denom))

		hostZone, hostZoneFound := k.GetHostZone(ctx, depositRecord.HostZoneId)
		if !hostZoneFound {
			k.Logger(ctx).Error(fmt.Sprintf("[StakeExistingDepositsOnHostZones] Host zone not found for deposit record {%d}", depositRecord.Id))
			continue
		}

		delegateAccount := hostZone.DelegationAccount
		if delegateAccount == nil || delegateAccount.Address == "" {
			k.Logger(ctx).Error(fmt.Sprintf("[StakeExistingDepositsOnHostZones] Zone %s is missing a delegation address!", hostZone.ChainId))
			continue
		}

		k.Logger(ctx).Info(utils.LogWithHostZone(depositRecord.HostZoneId, "Staking %v%s", depositRecord.Amount, hostZone.HostDenom))
		stakeAmount := sdk.NewCoin(hostZone.HostDenom, depositRecord.Amount)

		err := k.DelegateOnHost(ctx, hostZone, stakeAmount, depositRecord)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Did not stake %s on %s | err: %s", stakeAmount.String(), hostZone.ChainId, err.Error()))
			continue
		}
		k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Successfully submitted stake"))

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute("hostZone", hostZone.ChainId),
				sdk.NewAttribute("newAmountStaked", depositRecord.Amount.String()),
			),
		)
	}
}

func (k Keeper) DelegateOnHost(ctx sdk.Context, hostZone types.HostZone, amt sdk.Coin, depositRecord recordstypes.DepositRecord) error {
	// the relevant ICA is the delegate account
	owner := types.FormatICAAccountOwner(hostZone.ChainId, types.ICAType_DELEGATION)
	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "%s has no associated portId", owner)
	}
	connectionId, err := k.GetConnectionId(ctx, portID)

	if err != nil {
		return errorsmod.Wrapf(types.ErrInvalidICAChannel, "%s has no associated connection", portID)
	}

	// Fetch the relevant ICA
	delegationIca := hostZone.DelegationAccount
	if delegationIca == nil || delegationIca.Address == "" {
		k.Logger(ctx).Error(fmt.Sprintf("Zone %s is missing a delegation address!", hostZone.ChainId))
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid delegation account")
	}

	// Construct the transaction
	targetDelegatedAmts, err := k.GetTargetValAmtsForHostZone(ctx, hostZone, amt.Amount)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error getting target delegation amounts for host zone %s", hostZone.ChainId))
		return err
	}

	var splitDelegations []*types.SplitDelegation
	var msgs []sdk.Msg
	for _, validator := range hostZone.Validators {
		relativeAmount := sdk.NewCoin(amt.Denom, targetDelegatedAmts[validator.Address])
		if relativeAmount.Amount.IsPositive() {
			msgs = append(msgs, &stakingTypes.MsgDelegate{
				DelegatorAddress: delegationIca.Address,
				ValidatorAddress: validator.Address,
				Amount:           relativeAmount,
			})
		}
		splitDelegations = append(splitDelegations, &types.SplitDelegation{
			Validator: validator.Address,
			Amount:    relativeAmount.Amount,
		})
	}
	k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Preparing MsgDelegates from the delegation account to each validator"))

	// add callback data
	delegateCallback := types.DelegateCallback{
		HostZoneId:       hostZone.ChainId,
		DepositRecordId:  depositRecord.Id,
		SplitDelegations: splitDelegations,
	}
	k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Marshalling DelegateCallback args: %+v", delegateCallback))
	marshalledCallbackArgs, err := k.MarshalDelegateCallbackArgs(ctx, delegateCallback)
	if err != nil {
		return err
	}

	// Send the transaction through SubmitTx
	_, err = k.SubmitTxsEpoch(ctx, connectionId, msgs, *delegationIca, epochstypes.STAYKING_EPOCH, ICACallbackID_Delegate, marshalledCallbackArgs)
	if err != nil {
		return errorsmod.Wrapf(err, "Failed to SubmitTxs for connectionId %s on %s. Messages: %s", connectionId, hostZone.ChainId, msgs)
	}
	k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "ICA MsgDelegates Successfully Sent"))

	// update the record state to DELEGATION_IN_PROGRESS
	depositRecord.Status = recordstypes.DepositRecord_DELEGATION_IN_PROGRESS
	k.RecordsKeeper.SetDepositRecord(ctx, depositRecord)

	return nil
}

func (k Keeper) ReinvestRewards(ctx sdk.Context) {
	k.Logger(ctx).Info("Reinvesting tokens...")

	for _, hostZone := range k.GetAllHostZone(ctx) {
		// only process host zones once withdrawal accounts are registered
		withdrawalIca := hostZone.WithdrawalAccount
		if withdrawalIca == nil {
			k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Withdrawal account not registered for host zone"))
			continue
		}

		// read clock time on host zone
		blockTime, err := k.GetLightClientTimeSafely(ctx, hostZone.ConnectionId)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Could not find blockTime for host zone %s, err: %s", hostZone.ConnectionId, err.Error()))
			continue
		}
		k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "BlockTime for host zone: %d", blockTime))

		err = k.SubmitBalanceICQ(ctx, hostZone)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Error updating withdrawal balance for host zone %s: %s", hostZone.ConnectionId, err.Error()))
			continue
		}
	}
}
