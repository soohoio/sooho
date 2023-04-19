package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	recordstypes "github.com/soohoio/stayking/v2/x/records/types"
)

func (k msgServer) ExitLeverageStake(goCtx context.Context, msg *types.MsgExitLeverageStake) (*types.MsgExitLeverageStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger(ctx).Info(fmt.Sprintf("Exit Leverage stake: %s", msg.String()))
	sender, _ := sdk.AccAddressFromBech32(msg.GetCreator())
	hostZone, found := k.GetHostZone(ctx, msg.ChainId)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrInvalidHostZone, "host zone(chainId) is invalid: %s", msg.ChainId)
	}
	epochTracker, found := k.GetEpochTracker(ctx, "day")
	if !found {
		return nil, errorsmod.Wrapf(types.ErrEpochNotFound, "epoch tracker found: %s", "day")
	}
	redemptionId := recordstypes.UserRedemptionRecordKeyFormatter(hostZone.ChainId, epochTracker.EpochNumber, sender.String())

	_, found = k.RecordsKeeper.GetUserRedemptionRecord(ctx, redemptionId)
	if found {
		return nil, errorsmod.Wrapf(recordstypes.ErrRedemptionAlreadyExists, "user already redeemed this epoch: %s", redemptionId)
	}

	position, found := k.GetPosition(ctx, msg.PositionId)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrPositionNotFound, "position not found: %s", msg.PositionId)

	}

	stDenom := types.StAssetDenomFromHostZoneDenom(hostZone.HostDenom)

	// @TODO how to check module account's balance

	//convert to native Token Amount
	nativeTokenAmount := sdk.NewDecFromInt(position.StTokenAmount).Mul(hostZone.RedemptionRate).RoundInt()
	if !nativeTokenAmount.IsPositive() {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "amount must be greater than 0. found: %v", position.StTokenAmount)
	}

	// Check HostZone Balance
	if nativeTokenAmount.GT(hostZone.StakedBal) {
		return nil, errorsmod.Wrapf(types.ErrInvalidAmount, "cannot unstake an amount g.t. staked balance on host zone: %v", position.StTokenAmount)
	}

	k.Logger(ctx).Info(fmt.Sprintf("position record stDenom amount: %v%s", position.StTokenAmount, stDenom))

	// userRedemption Record 생성
	userRedemptionRecord := recordstypes.UserRedemptionRecord{
		Id:             redemptionId,
		Sender:         sender.String(),
		Receiver:       msg.GetReceiver(),
		Amount:         nativeTokenAmount,
		Denom:          hostZone.HostDenom,
		HostZoneId:     hostZone.ChainId,
		EpochNumber:    epochTracker.EpochNumber,
		ClaimIsPending: false,
	}

	epochUnbondingRecord, found := k.RecordsKeeper.GetEpochUnbondingRecord(ctx, epochTracker.EpochNumber)
	if !found {
		k.Logger(ctx).Error("latest epoch unbonding record not found")
		return nil, errorsmod.Wrapf(recordstypes.ErrEpochUnbondingRecordNotFound, "latest epoch unbonding record not found")
	}

	hostZoneUnbonding, found := k.RecordsKeeper.GetHostZoneUnbondingByChainId(ctx, epochUnbondingRecord.EpochNumber, hostZone.ChainId)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrInvalidHostZone, "host zone not found in unbondings: %s", hostZone.ChainId)
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
		return nil, sdkerrors.Wrapf(types.ErrEpochNotFound, "couldn't set host zone epoch unbonding record.")
	}
	k.RecordsKeeper.SetEpochUnbondingRecord(ctx, *updatedEpochUnbondingRecord)

	//update Position with native amount
	position.NativeTokenAmount = nativeTokenAmount
	position.Status = types.PositionStatus_POSITION_UNBONDING_IN_PROGRESS
	k.SetPosition(ctx, position)
	k.Logger(ctx).Info(fmt.Sprintf("position updated with native token amount %v", position.NativeTokenAmount))

	k.Logger(ctx).Info(fmt.Sprintf("executed Exit Leverage stake: %s", msg.String()))

	return &types.MsgExitLeverageStakeResponse{}, nil

}
