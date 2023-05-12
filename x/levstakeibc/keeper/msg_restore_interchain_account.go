package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	admintypes "github.com/soohoio/stayking/v2/x/admin/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	recordtypes "github.com/soohoio/stayking/v2/x/records/types"
)

func (k msgServer) RestoreInterchainAccount(_ctx context.Context, msg *types.MsgRestoreInterchainAccount) (*types.MsgRestoreInterchainAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(_ctx)
	// admin address check
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if !k.AdminKeeper.IsAdmin(ctx, creator) {
		return nil, admintypes.ErrNotAdmin
	}
	hostZone, found := k.GetHostZone(ctx, msg.ChainId)

	if !found {
		k.Logger(ctx).Error(fmt.Sprintf("host zone not found by chain id %s", msg.ChainId))
		return nil, types.ErrHostZoneNotFound
	}

	connectionEnd, found := k.IBCKeeper.ConnectionKeeper.GetConnection(ctx, hostZone.ConnectionId)
	if !found {
		errMsg := fmt.Sprintf("invalid connection id from host %s, %s not found", msg.ChainId, hostZone.ConnectionId)
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, errMsg)
	}

	owner := types.FormatICAAccountOwner(msg.ChainId, msg.AccountType)
	portID, err := icatypes.NewControllerPortID(owner)

	if err != nil {
		errMsg := fmt.Sprintf("could not create portID for ICA controller account address: %s", owner)
		k.Logger(ctx).Error(errMsg)
		return nil, err
	}

	_, exists := k.ICAControllerKeeper.GetInterchainAccountAddress(ctx, hostZone.ConnectionId, portID)
	if !exists {
		errMsg := fmt.Sprintf("ICA controller account address not found: %s", owner)
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(types.ErrInvalidInterchainAccountAddress, errMsg)
	}

	appVersion := string(icatypes.ModuleCdc.MustMarshalJSON(&icatypes.Metadata{
		Version:                icatypes.Version,
		ControllerConnectionId: hostZone.ConnectionId,
		HostConnectionId:       connectionEnd.Counterparty.ConnectionId,
		Encoding:               icatypes.EncodingProtobuf,
		TxType:                 icatypes.TxTypeSDKMultiMsg,
	}))

	if err := k.ICAControllerKeeper.RegisterInterchainAccount(ctx, hostZone.ConnectionId, owner, appVersion); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("unable to register %s account : %s", msg.AccountType.String(), err))
		return nil, err
	}

	if msg.AccountType == types.ICAType_DELEGATION {
		depositRecords := k.RecordsKeeper.GetAllDepositRecord(ctx)
		for _, depositRecord := range depositRecords {
			if depositRecord.HostZoneId == hostZone.ChainId && depositRecord.Status == recordtypes.DepositRecord_DELEGATION_IN_PROGRESS {
				depositRecord.Status = recordtypes.DepositRecord_DELEGATION_QUEUE
				k.Logger(ctx).Info(fmt.Sprintf("Setting DepositRecord %d to status DepositRecord_DELEGATION_IN_PROGRESS", depositRecord.Id))
				k.RecordsKeeper.SetDepositRecord(ctx, depositRecord)
			}
		}

		epochUnbondingRecords := k.RecordsKeeper.GetAllEpochUnbondingRecord(ctx)
		epochNumberForPendingUnbondingRecords := []uint64{}
		epochNumberForPendingTransferRecords := []uint64{}
		for _, epochUnbondingRecord := range epochUnbondingRecords {
			hostZoneUnbonding, found := k.RecordsKeeper.GetHostZoneUnbondingByChainId(ctx, epochUnbondingRecord.EpochNumber, hostZone.ChainId)
			if !found {
				k.Logger(ctx).Info(fmt.Sprintf("No HostZoneUnbonding found for chainId: %s, epoch: %d", hostZone.ChainId, epochUnbondingRecord.EpochNumber))
				continue
			}

			if hostZoneUnbonding.Status == recordtypes.HostZoneUnbonding_UNBONDING_IN_PROGRESS {
				k.Logger(ctx).Info(fmt.Sprintf("HostZoneUnbonding for %s at EpochNumber %d is stuck in status %s",
					hostZone.ChainId, epochUnbondingRecord.EpochNumber, recordtypes.HostZoneUnbonding_UNBONDING_IN_PROGRESS.String(),
				))
				epochNumberForPendingUnbondingRecords = append(epochNumberForPendingUnbondingRecords, epochUnbondingRecord.EpochNumber)

			} else if hostZoneUnbonding.Status == recordtypes.HostZoneUnbonding_EXIT_TRANSFER_IN_PROGRESS {
				k.Logger(ctx).Info(fmt.Sprintf("HostZoneUnbonding for %s at EpochNumber %d to in status %s",
					hostZone.ChainId, epochUnbondingRecord.EpochNumber, recordtypes.HostZoneUnbonding_EXIT_TRANSFER_IN_PROGRESS.String(),
				))
				epochNumberForPendingTransferRecords = append(epochNumberForPendingTransferRecords, epochUnbondingRecord.EpochNumber)
			}
		}
		err := k.RecordsKeeper.SetHostZoneUnbondings(ctx, hostZone.ChainId, epochNumberForPendingUnbondingRecords, recordtypes.HostZoneUnbonding_UNBONDING_QUEUE)
		if err != nil {
			errMsg := fmt.Sprintf("unable to update host zone unbonding record status to %s for chainId: %s and epochUnbondingRecordIds: %v, err: %s",
				recordtypes.HostZoneUnbonding_UNBONDING_QUEUE.String(), hostZone.ChainId, epochNumberForPendingUnbondingRecords, err)
			k.Logger(ctx).Error(errMsg)
			return nil, err
		}

		err = k.RecordsKeeper.SetHostZoneUnbondings(ctx, hostZone.ChainId, epochNumberForPendingTransferRecords, recordtypes.HostZoneUnbonding_EXIT_TRANSFER_QUEUE)
		if err != nil {
			errMsg := fmt.Sprintf("unable to update host zone unbonding record status to %s for chainId: %s and epochUnbondingRecordIds: %v, err: %s",
				recordtypes.HostZoneUnbonding_EXIT_TRANSFER_QUEUE.String(), hostZone.ChainId, epochNumberForPendingTransferRecords, err)
			k.Logger(ctx).Error(errMsg)
			return nil, err
		}
	}

	return &types.MsgRestoreInterchainAccountResponse{}, nil
}
