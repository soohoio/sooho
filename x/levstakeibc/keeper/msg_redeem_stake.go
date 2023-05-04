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

func (k msgServer) RedeemStake(_ctx context.Context, msg *types.MsgRedeemStake) (*types.MsgRedeemStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(_ctx)
	k.Logger(ctx).Info(fmt.Sprintf("redeem stake: %s", msg.String()))

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

	stDenom := types.StAssetDenomFromHostZoneDenom(hostZone.HostDenom)

	balance := k.bankKeeper.GetBalance(ctx, sender, stDenom)
	if balance.Amount.LT(msg.StTokenAmount) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "balance is lower than redemption amount. redemption amount: %v, balance %v: ", msg.StTokenAmount, balance.Amount)
	}

	nativeTokenAmount := sdk.NewDecFromInt(msg.StTokenAmount).Mul(hostZone.RedemptionRate).RoundInt()
	if !nativeTokenAmount.IsPositive() {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "amount must be greater than 0. found: %v", msg.StTokenAmount)
	}

	if nativeTokenAmount.GT(hostZone.StakedBal) {
		return nil, errorsmod.Wrapf(types.ErrInvalidAmount, "cannot unstake an amount g.t. staked balance on host zone: %v", msg.StTokenAmount)
	}

	coinString := nativeTokenAmount.String() + stDenom
	inCoin, err := sdk.ParseCoinNormalized(coinString)

	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "could not parse inCoin: %s. err: %s", coinString, err.Error())
	}

	k.Logger(ctx).Info(fmt.Sprintf("Redemption requested redemption amount: %v%s", inCoin.Amount, inCoin.Denom))
	k.Logger(ctx).Info(fmt.Sprintf("Redemption issuer IBCDenom balance: %v%s", balance.Amount, balance.Denom))

	userRedemptionRecord := recordstypes.UserRedemptionRecord{
		Id:             redemptionId,
		Sender:         sender.String(),
		Receiver:       msg.GetReceiver(),
		Amount:         nativeTokenAmount,
		Denom:          hostZone.HostDenom,
		HostZoneId:     hostZone.ChainId,
		EpochNumber:    epochTracker.EpochNumber,
		ClaimIsPending: false,
		PositionId:     0,
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

	redeemCoin := sdk.NewCoins(sdk.NewCoin(stDenom, msg.StTokenAmount))
	bech32ZoneAddress, err := sdk.AccAddressFromBech32(hostZone.Address)

	if err != nil {
		return nil, fmt.Errorf("could not bech32 decode address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}
	err = k.bankKeeper.SendCoins(ctx, sender, bech32ZoneAddress, redeemCoin)
	if err != nil {
		k.Logger(ctx).Error("Failed to send sdk.NewCoins(inCoins) from account to module")
		return nil, sdkerrors.Wrapf(types.ErrInsufficientFunds, "couldn't send %v derivative %s tokens to module account. err: %s", msg.StTokenAmount, hostZone.HostDenom, err.Error())
	}
	hostZoneUnbonding.StTokenAmount = hostZoneUnbonding.StTokenAmount.Add(msg.StTokenAmount)

	k.RecordsKeeper.SetUserRedemptionRecord(ctx, userRedemptionRecord)

	hostZoneUnbondings := epochUnbondingRecord.GetHostZoneUnbondings()
	if hostZoneUnbondings == nil {
		hostZoneUnbondings = []*recordstypes.HostZoneUnbonding{}
		epochUnbondingRecord.HostZoneUnbondings = hostZoneUnbondings
	}
	updatedEpochUnbondingRecord, success := k.RecordsKeeper.AddHostZoneToEpochUnbondingRecord(ctx, epochUnbondingRecord.EpochNumber, hostZone.ChainId, hostZoneUnbonding)
	if !success {
		k.Logger(ctx).Error(fmt.Sprintf("Failed to set host zone epoch unbonding record: epochNumber %d, chainId %s, hostZoneUnbonding %v", epochUnbondingRecord.EpochNumber, hostZone.ChainId, hostZoneUnbonding))
		return nil, sdkerrors.Wrapf(types.ErrEpochNotFound, "couldn't set host zone epoch unbonding record. err: %s", err.Error())
	}
	k.RecordsKeeper.SetEpochUnbondingRecord(ctx, *updatedEpochUnbondingRecord)

	k.Logger(ctx).Info(fmt.Sprintf("executed redeem stake: %s", msg.String()))
	return &types.MsgRedeemStakeResponse{}, nil
}
