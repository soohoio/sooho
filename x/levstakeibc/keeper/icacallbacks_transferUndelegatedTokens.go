package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	"github.com/golang/protobuf/proto"
	"github.com/soohoio/stayking/v2/utils"
	icacallbackstypes "github.com/soohoio/stayking/v2/x/icacallbacks/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	recordstypes "github.com/soohoio/stayking/v2/x/records/types"
	//recordstypes "github.com/soohoio/stayking/v2/x/records/types"
)

// Marshalls TransferUndelegatedTokens callback arguments
func (k Keeper) MarshalTransferUndelegatedTokensArgs(ctx sdk.Context, TransferUndelegatedTokensCallback types.TransferUndelegatedTokensCallback) ([]byte, error) {
	out, err := proto.Marshal(&TransferUndelegatedTokensCallback)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("MarshalTransferUndelegatedTokensCallbackArgs %v", err.Error()))
		return nil, err
	}
	return out, nil
}

// Unmarshalls TransferUndelegatedTokens callback arguments into a TransferUndelegatedTokens struct
func (k Keeper) UnmarshalTransferUndelegatedTokensCallbackArgs(ctx sdk.Context, transferUndelegatedTokensCallback []byte) (*types.TransferUndelegatedTokensCallback, error) {
	unmarshalledTransferUndelegatedTokensCallback := types.TransferUndelegatedTokensCallback{}
	if err := proto.Unmarshal(transferUndelegatedTokensCallback, &unmarshalledTransferUndelegatedTokensCallback); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("UnmarshalTransferUndelegatedTokensCallbackArgs %v", err.Error()))
		return nil, err
	}
	return &unmarshalledTransferUndelegatedTokensCallback, nil
}

func TransferUndelegatedTokensCallback(k Keeper, ctx sdk.Context, packet channeltypes.Packet, ackResponse *icacallbackstypes.AcknowledgementResponse, args []byte) error {
	// Deserialize the callback args
	transferUndelegatedTokensCallback, err := k.UnmarshalTransferUndelegatedTokensCallbackArgs(ctx, args)
	if err != nil {
		return errorsmod.Wrapf(types.ErrUnmarshalFailure, fmt.Sprintf("Unable to unmarshal transferUndelegatedTokens callback args: %s", err.Error()))
	}
	chainId := transferUndelegatedTokensCallback.HostZoneId
	k.Logger(ctx).Info(utils.LogICACallbackWithHostZone(chainId, ICACallbackID_TransferUndelegatedTokens,
		"Starting TransferUndelegatedTokens callback for HostZone: %v", transferUndelegatedTokensCallback.HostZoneId))

	// Confirm chainId and deposit record Id exist
	hostZone, found := k.GetHostZone(ctx, chainId)
	if !found {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "host zone not found %s", chainId)
	}
	k.Logger(ctx).Info(fmt.Sprintf("[DEBUG]hostZone Address : %v", hostZone.Address))

	zoneAddress, err := sdk.AccAddressFromBech32(hostZone.Address)
	k.Logger(ctx).Info(fmt.Sprintf("[DEBUG]hostZone Address : %v", zoneAddress))
	if err != nil {
		return fmt.Errorf("could not bech32 decode address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}
	// Check for timeout (ack nil)
	if ackResponse.Status == icacallbackstypes.AckResponseStatus_TIMEOUT {
		k.Logger(ctx).Error(utils.LogICACallbackStatusWithHostZone(chainId, ICACallbackID_TransferUndelegatedTokens,
			icacallbackstypes.AckResponseStatus_TIMEOUT, packet))

		// Reset hostzone unbonding record status
		err = k.RecordsKeeper.SetHostZoneUnbondings(ctx, hostZone.ChainId, transferUndelegatedTokensCallback.EpochUnbondingRecordIds, recordstypes.HostZoneUnbonding_EXIT_TRANSFER_QUEUE)
		if err != nil {
			k.Logger(ctx).Error(err.Error())
			fmt.Errorf("Error SetHostZone Unbondings hostZone:%s", hostZone.ChainId)
		}
		return nil
	}

	// Check for a failed transaction (ack error)
	// Reset the deposit record status upon failure
	if ackResponse.Status == icacallbackstypes.AckResponseStatus_FAILURE {
		k.Logger(ctx).Error(utils.LogICACallbackStatusWithHostZone(chainId, ICACallbackID_Delegate,
			icacallbackstypes.AckResponseStatus_FAILURE, packet))

		err = k.RecordsKeeper.SetHostZoneUnbondings(ctx, hostZone.ChainId, transferUndelegatedTokensCallback.EpochUnbondingRecordIds, recordstypes.HostZoneUnbonding_EXIT_TRANSFER_QUEUE)
		if err != nil {
			k.Logger(ctx).Error(err.Error())
			fmt.Errorf("Error SetHostZone Unbondings hostZone:%s", hostZone.ChainId)
		}
		return nil
	}

	err = k.RecordsKeeper.SetHostZoneUnbondings(ctx, chainId, transferUndelegatedTokensCallback.EpochUnbondingRecordIds, recordstypes.HostZoneUnbonding_RELEASE_ASSET_QUEUE)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) DecrementHostZoneUnbondingAmount(ctx sdk.Context, userRedemptionRecord recordstypes.UserRedemptionRecord, chainId string) error {
	// fetch the hzu associated with the user unbonding record
	hostZoneUnbonding, found := k.RecordsKeeper.GetHostZoneUnbondingByChainId(ctx, userRedemptionRecord.EpochNumber, chainId)
	if !found {
		return errorsmod.Wrapf(types.ErrRecordNotFound, "host zone unbonding not found %s", chainId)
	}
	// decrement the hzu by the amount claimed
	hostZoneUnbonding.NativeTokenAmount = hostZoneUnbonding.NativeTokenAmount.Sub(userRedemptionRecord.Amount)
	// save the updated hzu on the epoch unbonding record
	epochUnbondingRecord, success := k.RecordsKeeper.AddHostZoneToEpochUnbondingRecord(ctx, userRedemptionRecord.EpochNumber, chainId, hostZoneUnbonding)
	if !success {
		return errorsmod.Wrapf(types.ErrRecordNotFound, "epoch unbonding record not found %s", chainId)
	}
	k.RecordsKeeper.SetEpochUnbondingRecord(ctx, *epochUnbondingRecord)
	return nil
}
