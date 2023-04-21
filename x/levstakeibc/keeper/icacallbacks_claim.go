package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	"github.com/gogo/protobuf/proto"
	"github.com/soohoio/stayking/v2/utils"
	icacallbackstypes "github.com/soohoio/stayking/v2/x/icacallbacks/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	recordstypes "github.com/soohoio/stayking/v2/x/records/types"
)

func (k Keeper) MarshalClaimCallbackArgs(ctx sdk.Context, claimCallback types.ClaimCallback) ([]byte, error) {
	out, err := proto.Marshal(&claimCallback)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("MarshalClaimCallbackArgs %v", err.Error()))
		return nil, err
	}
	return out, nil
}

func (k Keeper) UnmarshalClaimCallbackArgs(ctx sdk.Context, claimCallback []byte) (*types.ClaimCallback, error) {
	unmarshalledDelegateCallback := types.ClaimCallback{}
	if err := proto.Unmarshal(claimCallback, &unmarshalledDelegateCallback); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("UnmarshalClaimCallbackArgs %v", err.Error()))
		return nil, err
	}
	return &unmarshalledDelegateCallback, nil
}

func ClaimCallback(k Keeper, ctx sdk.Context, packet channeltypes.Packet, ackResponse *icacallbackstypes.AcknowledgementResponse, args []byte) error {
	claimCallback, err := k.UnmarshalClaimCallbackArgs(ctx, args)
	if err != nil {
		return errorsmod.Wrapf(types.ErrUnmarshalFailure, fmt.Sprintf("Unable to unmarshal claim callback args: %s", err.Error()))
	}
	chainId := claimCallback.ChainId
	k.Logger(ctx).Info(utils.LogICACallbackWithHostZone(chainId, ICACallbackID_Claim,
		"Starting claim callback for Redemption Record: %s", claimCallback.UserRedemptionRecordId))

	userRedemptionRecord, found := k.RecordsKeeper.GetUserRedemptionRecord(ctx, claimCallback.GetUserRedemptionRecordId())
	if !found {
		return errorsmod.Wrapf(types.ErrRecordNotFound, "user redemption record not found %s", claimCallback.GetUserRedemptionRecordId())
	}

	if ackResponse.Status == icacallbackstypes.AckResponseStatus_TIMEOUT {
		k.Logger(ctx).Error(utils.LogICACallbackStatusWithHostZone(chainId, ICACallbackID_Claim,
			icacallbackstypes.AckResponseStatus_TIMEOUT, packet))

		userRedemptionRecord.ClaimIsPending = false
		k.RecordsKeeper.SetUserRedemptionRecord(ctx, userRedemptionRecord)
		return nil
	}

	if ackResponse.Status == icacallbackstypes.AckResponseStatus_FAILURE {
		k.Logger(ctx).Error(utils.LogICACallbackStatusWithHostZone(chainId, ICACallbackID_Claim,
			icacallbackstypes.AckResponseStatus_FAILURE, packet))

		userRedemptionRecord.ClaimIsPending = false
		k.RecordsKeeper.SetUserRedemptionRecord(ctx, userRedemptionRecord)
		return nil
	}

	k.Logger(ctx).Info(utils.LogICACallbackStatusWithHostZone(chainId, ICACallbackID_Claim,
		icacallbackstypes.AckResponseStatus_SUCCESS, packet)) // handle failed tx on host chain

	k.RecordsKeeper.RemoveUserRedemptionRecord(ctx, claimCallback.GetUserRedemptionRecordId())
	err = k.DecrementHostZoneUnbonding(ctx, userRedemptionRecord, *claimCallback)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("ClaimCallback failed (DecrementHostZoneUnbonding), packet %v, err: %s", packet, err.Error()))
		return err
	}

	k.Logger(ctx).Info(fmt.Sprintf("[CLAIM] success on %s", userRedemptionRecord.GetHostZoneId()))
	return nil
}

func (k Keeper) DecrementHostZoneUnbonding(ctx sdk.Context, userRedemptionRecord recordstypes.UserRedemptionRecord, callbackArgs types.ClaimCallback) error {
	// fetch the hzu associated with the user unbonding record
	hostZoneUnbonding, found := k.RecordsKeeper.GetHostZoneUnbondingByChainId(ctx, callbackArgs.EpochNumber, callbackArgs.ChainId)
	if !found {
		return errorsmod.Wrapf(types.ErrRecordNotFound, "host zone unbonding not found %s", callbackArgs.ChainId)
	}
	// decrement the hzu by the amount claimed
	hostZoneUnbonding.NativeTokenAmount = hostZoneUnbonding.NativeTokenAmount.Sub(userRedemptionRecord.Amount)
	// save the updated hzu on the epoch unbonding record
	epochUnbondingRecord, success := k.RecordsKeeper.AddHostZoneToEpochUnbondingRecord(ctx, callbackArgs.EpochNumber, callbackArgs.ChainId, hostZoneUnbonding)
	if !success {
		return errorsmod.Wrapf(types.ErrRecordNotFound, "epoch unbonding record not found %s", callbackArgs.ChainId)
	}
	k.RecordsKeeper.SetEpochUnbondingRecord(ctx, *epochUnbondingRecord)
	return nil
}