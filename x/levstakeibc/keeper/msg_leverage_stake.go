package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	epochtypes "github.com/soohoio/stayking/v2/x/epochs/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func (k msgServer) LeverageStake(goCtx context.Context, msg *types.MsgLeverageStake) (*types.MsgLeverageStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Define Arguments
	equity := msg.Equity
	hostDenom := msg.HostDenom
	leverageRatio := msg.LeverageRatio
	levType := msg.GetStakeType(leverageRatio)

	if levType == types.StakingType_NotLeverageType {
		msg, err := k.stakeWithoutLeverage(ctx, equity, hostDenom, msg.Creator, levType)
		if err != nil {
			return nil, err
		}
		return msg, nil
	} else if levType == types.StakingType_LeverageType {
		k.stakeWithLeverage(ctx, equity, hostDenom, msg.Creator, leverageRatio, levType)
	}

	return nil, errorsmod.Wrapf(types.ErrInvalidLeverageRatio, "invalid leverage type value (lev ratio %v) ", leverageRatio)
}

func (k msgServer) stakeWithoutLeverage(ctx sdk.Context, equity sdk.Int, hostDenom string, creator string, levType types.StakingType) (*types.MsgLeverageStakeResponse, error) {

	hostZone, err := k.GetHostZoneFromHostDenom(ctx, hostDenom)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Host Zone not found for denom (%s)", hostDenom))
		return nil, errorsmod.Wrapf(types.ErrInvalidHostZone, "no host zone found for denom (%s)", hostDenom)
	}

	sender, _ := sdk.AccAddressFromBech32(creator)
	ibcDenom := hostZone.GetIbcDenom()
	coinString := equity.String() + ibcDenom
	inCoin, err := sdk.ParseCoinNormalized(coinString)

	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("failed to parse coin (%s)", coinString))
		return nil, sdkerrors.Wrapf(err, "failed to parse coin (%s)", coinString)
	}

	balance := k.bankKeeper.GetBalance(ctx, sender, ibcDenom)

	if balance.IsLT(inCoin) {
		k.Logger(ctx).Error(fmt.Sprintf("balance is lower than staking amount. staking amount: %v, balance: %v", equity, balance.Amount))
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "balance is lower than staking amount. staking amount: %v, balance: %v", equity, balance.Amount)
	}

	// check that the token is an IBC token
	isIbcToken := types.IsIBCToken(ibcDenom)
	if !isIbcToken {
		k.Logger(ctx).Error("invalid token denom - denom is not an IBC token (%s)", ibcDenom)
		return nil, sdkerrors.Wrapf(types.ErrInvalidToken, "denom is not an IBC token (%s)", ibcDenom)
	}

	zoneAddress, err := sdk.AccAddressFromBech32(hostZone.Address)
	if err != nil {
		return nil, fmt.Errorf("could not bech32 decode address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}

	err = k.bankKeeper.SendCoins(ctx, sender, zoneAddress, sdk.NewCoins(inCoin))

	if err != nil {
		k.Logger(ctx).Error("failed to send tokens from Account to Module")
		return nil, sdkerrors.Wrap(err, "failed to send tokens from Account to Module")
	}

	// mint user `amount` of the corresponding stAsset
	// NOTE: We should ensure that denoms are unique - we don't want anyone spoofing denoms
	err = k.MintStAssetAndTransfer(ctx, sender, equity, hostDenom, levType)
	if err != nil {
		k.Logger(ctx).Error("failed to send tokens from Account to Module")
		return nil, sdkerrors.Wrapf(err, "failed to mint %s stAssets to user", hostDenom)
	}

	// create a deposit record of these tokens (pending transfer)
	staykingEpochTracker, found := k.GetEpochTracker(ctx, epochtypes.STAYKING_EPOCH)

	if !found {
		k.Logger(ctx).Error("failed to find stayking epoch")
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "no epoch number for epoch (%s)", epochtypes.STAYKING_EPOCH)
	}

	depositRecord, found := k.RecordsKeeper.GetDepositRecordByEpochAndChain(ctx, staykingEpochTracker.EpochNumber, hostZone.ChainId)
	if !found {
		k.Logger(ctx).Error("failed to find deposit record")
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, fmt.Sprintf("no deposit record for epoch (%d)", staykingEpochTracker.EpochNumber))
	}
	depositRecord.Amount = depositRecord.Amount.Add(equity)

	k.RecordsKeeper.SetDepositRecord(ctx, *depositRecord)

	return &types.MsgLeverageStakeResponse{}, nil
}

func (k msgServer) stakeWithLeverage(ctx sdk.Context, equity sdk.Int, denom string, creator string, ratio sdk.Dec, levType types.StakingType) {

	k.Logger(ctx).Info("leverageType Mode ... ")
	k.Logger(ctx).Info(fmt.Sprintf("stakeWithLeverage => equity: %v, denom: %v, creator: %v, ratio: %v, reverageType: %v", equity, denom, creator, ratio, levType))

	// TODO: 1) Denom 에 따른 가격비 x/record 에서 가져오기

	// TODO: 2) x/lendingpool Borrow 호출 전 collateral, leverageRatio 요소를 통해 borrowingAmount, markPrice 를 계산해서 Borrow 함수 호출 하기
}

func (k msgServer) MintStAssetAndTransfer(ctx sdk.Context, receiver sdk.AccAddress, amount sdk.Int, denom string, leverageType types.StakingType) error {
	stAssetDenom := types.StAssetDenomFromHostZoneDenom(denom)

	hz, _ := k.GetHostZoneFromHostDenom(ctx, denom)
	amountToMint := (sdk.NewDecFromInt(amount).Quo(hz.RedemptionRate)).TruncateInt()
	coinString := amountToMint.String() + stAssetDenom
	stCoins, err := sdk.ParseCoinsNormalized(coinString)

	if err != nil {
		k.Logger(ctx).Error("Failed to parse coins")
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Failed to parse coins %s", coinString)
	}

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, stCoins)
	if err != nil {
		k.Logger(ctx).Error("Failed to mint coins")
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Failed to mint coins")
	}

	// TODO: Mint 와 Transfer 분리
	if leverageType == types.StakingType_NotLeverageType {
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, stCoins)
		if err != nil {
			k.Logger(ctx).Error("Failed to send coins from module to account")
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Failed to send %s from module to account", stCoins.GetDenomByIndex(0))
		}
	} else {
		// TODO: stToken Module Account 에 그대로 두고 기록하기
	}

	k.Logger(ctx).Info(fmt.Sprintf("[MINT ST ASSET] success on %s.", hz.GetChainId()))
	return nil
}
