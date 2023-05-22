package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	epochtypes "github.com/soohoio/stayking/v2/x/epochs/types"
	lendingpooltypes "github.com/soohoio/stayking/v2/x/lendingpool/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func (k msgServer) LeverageStake(goCtx context.Context, msg *types.MsgLeverageStake) (*types.MsgLeverageStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Define Arguments
	equity := msg.Equity
	hostDenom := msg.HostDenom
	leverageRatio := msg.LeverageRatio
	levType := msg.GetStakeType(leverageRatio)

	if levType == types.StakingType_NOT_LEVERAGE_TYPE {
		msg, err := k.stakeWithoutLeverage(ctx, equity, hostDenom, msg.Creator)
		if err != nil {
			k.Logger(ctx).Error("[CUSTOM DEBUG] LeverageStake.NotLeverageType error reason : " + err.Error())
			return nil, err
		}
		return msg, nil
	} else if levType == types.StakingType_LEVERAGE_TYPE {
		msg, err := k.stakeWithLeverage(ctx, equity, hostDenom, msg.Creator, leverageRatio, levType, msg.LendingPoolDenom)
		if err != nil {
			k.Logger(ctx).Error("[CUSTOM DEBUG] LeverageStake.LeverageType error reason : " + err.Error())
			return nil, err
		}
		return msg, nil
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)
	return nil, errorsmod.Wrapf(types.ErrLeverageRatio, "invalid leverage type value (lev ratio %v) ", leverageRatio)

}

func (k msgServer) stakeWithoutLeverage(ctx sdk.Context, equity sdk.Int, hostDenom string, creator string) (*types.MsgLeverageStakeResponse, error) {

	hostZone, err := k.GetHostZoneFromHostDenom(ctx, hostDenom)

	if err != nil {
		printErr := fmt.Sprintf("no host zone found by host denom (%s)", hostDenom)
		k.Logger(ctx).Error(printErr)
		return nil, errorsmod.Wrap(types.ErrHostZoneNotFound, printErr)
	}

	sender, _ := sdk.AccAddressFromBech32(creator)
	ibcDenom := hostZone.GetIbcDenom()
	coinString := equity.String() + ibcDenom
	inCoin, err := sdk.ParseCoinNormalized(coinString)

	if err != nil {
		printErr := fmt.Sprintf("failed to parse coin (%s)", coinString)
		k.Logger(ctx).Error(printErr)
		return nil, errorsmod.Wrap(err, printErr)
	}

	balance := k.bankKeeper.GetBalance(ctx, sender, ibcDenom)

	if balance.IsLT(inCoin) {
		printErr := fmt.Sprintf("balance is lower than staking amount. staking amount: %v, balance: %v", equity, balance.Amount)
		k.Logger(ctx).Error(printErr)
		return nil, errorsmod.Wrap(sdkerrors.ErrInsufficientFunds, printErr)
	}

	isIbcToken := types.IsIBCToken(ibcDenom)
	if !isIbcToken {
		printErr := fmt.Sprintf("denom is not an IBC token (%s)", ibcDenom)
		k.Logger(ctx).Error(printErr)
		return nil, errorsmod.Wrap(types.ErrInvalidToken, printErr)
	}

	zoneAddress, err := sdk.AccAddressFromBech32(hostZone.Address)
	k.Logger(ctx).Info(fmt.Sprintf("[DEBUG] hostZone Address : %v", zoneAddress))

	if err != nil {
		return nil, fmt.Errorf("could not bech32 decode address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}

	err = k.bankKeeper.SendCoins(ctx, sender, zoneAddress, sdk.NewCoins(inCoin))

	if err != nil {
		printErr := fmt.Sprintf("failed to send tokens from Account to ZoneAddress")
		k.Logger(ctx).Error(printErr)
		return nil, errorsmod.Wrap(err, printErr)
	}

	// mint `amount` of the corresponding stAsset from module
	stCoins, err := k.MintStAsset(ctx, equity, hostDenom)

	if err != nil {
		k.Logger(ctx).Error("Failed to mint stToken")
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Failed to mint stToken", stCoins.GetDenomByIndex(0))
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, stCoins)

	if err != nil {
		k.Logger(ctx).Error("failed to send st tokens from Account to Module")
		return nil, errorsmod.Wrapf(err, "failed to mint %s stAssets to user", hostDenom)
	}

	// create a deposit record of these tokens (pending transfer)
	staykingEpochTracker, found := k.GetEpochTracker(ctx, epochtypes.STAYKING_EPOCH)

	if !found {
		k.Logger(ctx).Error("failed to find stayking epoch")
		return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, "no epoch number for epoch (%s)", epochtypes.STAYKING_EPOCH)
	}

	depositRecord, found := k.RecordsKeeper.GetDepositRecordByEpochAndChain(ctx, staykingEpochTracker.EpochNumber, hostZone.ChainId)
	if !found {
		k.Logger(ctx).Error("failed to find deposit record")
		return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, fmt.Sprintf("no deposit record for epoch (%d)", staykingEpochTracker.EpochNumber))
	}
	depositRecord.Amount = depositRecord.Amount.Add(equity)

	k.RecordsKeeper.SetDepositRecord(ctx, *depositRecord)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeStakeWithoutLeverage,
			sdk.NewAttribute(types.AttributeKeyRecipientChain, hostZone.ChainId),
			sdk.NewAttribute(types.AttributeKeyFromAddress, creator),
			sdk.NewAttribute(types.AttributeKeyHostDenom, hostZone.HostDenom),
			sdk.NewAttribute(types.AttributeKeyIBCDenom, hostZone.IbcDenom),
			sdk.NewAttribute(types.AttributeKeyEpochNumber, string(staykingEpochTracker.EpochNumber)),
			sdk.NewAttribute(types.AttributeKeyNativeTokenAmount, equity.String()),
			sdk.NewAttribute(types.AttributeKeyDepositRecordId, string(depositRecord.Id)),
		),
	)
	return &types.MsgLeverageStakeResponse{}, nil
}

// TODO: Not Stake 와 중복되는 로직임... 리팩토링 필요
func (k msgServer) stakeWithLeverage(ctx sdk.Context, equity sdk.Int, denom string, creator string, ratio sdk.Dec, levType types.StakingType, lendingPoolDenom string) (*types.MsgLeverageStakeResponse, error) {
	k.Logger(ctx).Info("leverageType Mode ... ")
	k.Logger(ctx).Info(fmt.Sprintf("stakeWithLeverage => equity: %v, denom: %v, creator: %v, ratio: %v, reverageType: %v, lendingPoolDenom: %v", equity, denom, creator, ratio, levType, lendingPoolDenom))

	moduleAddress := k.accountKeeper.GetModuleAddress(types.ModuleName)
	k.Logger(ctx).Info(fmt.Sprintf("[DEBUG] module Address : %v", moduleAddress.String()))
	k.Logger(ctx).Info(fmt.Sprintf("[DEBUG] lendingpool module address : %v", k.accountKeeper.GetModuleAddress(lendingpooltypes.ModuleName).String()))

	sender, _ := sdk.AccAddressFromBech32(creator)

	hostZone, found := k.GetHostZoneByHostDenom(ctx, denom)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrHostZoneNotFound, "not found host zone by host denom %v", denom)
	}

	existsPostion, found := k.GetPositionByDenomAndSender(ctx, denom, creator)

	if found {
		return nil, errorsmod.Wrapf(types.ErrAlreadyExistsPosition, fmt.Sprintf("existed position by position id %v and denom %v", existsPostion.Id, denom))
	}

	borrowingAmount := sdk.NewDecFromInt(equity).Mul(ratio.Sub(sdk.NewDec(1))).TruncateInt()
	totalAsset := equity.Add(borrowingAmount)

	var debtRatio sdk.Dec

	if totalAsset.Equal(sdk.NewInt(0)) {
		debtRatio = sdk.ZeroDec()
	} else {
		debtRatio = sdk.NewDecFromInt(borrowingAmount).Quo(sdk.NewDecFromInt(totalAsset))
	}

	loanId, err := k.LendingPoolKeeper.Borrow(
		ctx,
		hostZone.GetIbcDenom(),
		types.ModuleName,
		sender,
		sdk.NewCoins(sdk.NewCoin(hostZone.GetIbcDenom(), borrowingAmount)),
		sdk.NewCoins(sdk.NewCoin(hostZone.GetIbcDenom(), equity)),
	)

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrLendingPoolBorrow, "Can't borrow debt amount (%v) from the pool (ibc denom > %v)", borrowingAmount, hostZone.GetIbcDenom())
	}

	k.Logger(ctx).Info(fmt.Sprintf("Successfully done for borrowing amount, LoanId : %v", loanId))

	ibcDenom := hostZone.GetIbcDenom()
	coinString := equity.String() + ibcDenom
	inCoin, err := sdk.ParseCoinNormalized(coinString)
	if err != nil {
		return nil, fmt.Errorf("failed to parsecoin normalized for %s", coinString)
	}
	balance := k.bankKeeper.GetBalance(ctx, sender, ibcDenom)

	if balance.IsLT(inCoin) {
		k.Logger(ctx).Error(fmt.Sprintf("balance is lower than staking amount. staking amount: %v, balance: %v", equity, balance.Amount))
		return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "balance is lower than staking amount. staking amount: %v, balance: %v", equity, balance.Amount)
	}

	zoneAddress, err := sdk.AccAddressFromBech32(hostZone.Address)
	if err != nil {
		return nil, fmt.Errorf("could not bech32 decode address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}
	//equity를 zone Address로 보냄
	err = k.bankKeeper.SendCoins(ctx, sender, zoneAddress, sdk.NewCoins(inCoin))
	if err != nil {
		return nil, fmt.Errorf("could not sendcoin to zone address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}
	//borrow한 asset을 zone address로 보냄.
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, zoneAddress, sdk.NewCoins(sdk.NewCoin(hostZone.GetIbcDenom(), borrowingAmount)))
	if err != nil {
		return nil, fmt.Errorf("could not sendcoin to zone address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}
	// borrowingAmount 에 대한 Transfer 를 받은 evmos 를 stToken 으로 민팅하여 모듈 어카운트에 저장하고
	stCoins, err := k.MintStAsset(ctx, totalAsset, denom)
	if err != nil {
		return nil, fmt.Errorf("could not miont stasset for denom %s and totalAsset: %s", denom, totalAsset)
	}
	k.Logger(ctx).Info(fmt.Sprintf("totalAsset : %v, collateral : %v, borrowingAmount : %v, debtRatio : %v", totalAsset, equity, borrowingAmount, debtRatio))
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, zoneAddress, stCoins)

	if err != nil {
		k.Logger(ctx).Error("failed to send st tokens from module to host address")
		return nil, errorsmod.Wrapf(err, "failed to mint %s stAssets to host address", hostZone.HostDenom)
	}

	staykingEpochTracker, found := k.GetEpochTracker(ctx, epochtypes.STAYKING_EPOCH)

	if !found {
		k.Logger(ctx).Error("failed to find stayking epoch")
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "no epoch number for epoch (%s)", epochtypes.STAYKING_EPOCH)
	}

	depositRecord, found := k.RecordsKeeper.GetDepositRecordByEpochAndChain(ctx, staykingEpochTracker.EpochNumber, hostZone.ChainId)
	if !found {
		k.Logger(ctx).Error("failed to find deposit record")
		return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, fmt.Sprintf("no deposit record for epoch (%d)", staykingEpochTracker.EpochNumber))
	}
	depositRecord.Amount = depositRecord.Amount.Add(totalAsset)

	k.RecordsKeeper.SetDepositRecord(ctx, *depositRecord)

	// x/levstakeibc store 객체에 Position 객체를 생성하여 저장한다.
	positionId := k.GetNextPositionID(ctx)

	position := types.Position{
		Id:                positionId,
		LoanId:            loanId,
		Sender:            creator,
		Denom:             denom,
		StTokenAmount:     stCoins.AmountOf(types.StAssetDenomFromHostZoneDenom(denom)),
		NativeTokenAmount: totalAsset,
		Status:            types.PositionStatus_POSITION_PENDING,
		Liquidated:        false,
		DepositRecordId:   depositRecord.Id,
	}

	k.SetPosition(ctx, position)

	k.SetNextPositionID(ctx, positionId+1) // 다음 포지션 ID ++

	k.Logger(ctx).Info(fmt.Sprintf("Successfully done for saving position data, PositionId : %v", positionId))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeStakeWithLeverage,
			sdk.NewAttribute(types.AttributeKeyRecipientChain, hostZone.ChainId),
			sdk.NewAttribute(types.AttributeKeyFromAddress, creator),
			sdk.NewAttribute(types.AttributeKeyHostDenom, hostZone.HostDenom),
			sdk.NewAttribute(types.AttributeKeyIBCDenom, hostZone.IbcDenom),
			sdk.NewAttribute(types.AttributeKeyEpochNumber, string(staykingEpochTracker.EpochNumber)),
			sdk.NewAttribute(types.AttributeKeyNativeTokenAmount, equity.String()),
			sdk.NewAttribute(types.AttributeKeyDepositRecordId, string(depositRecord.Id)),
			sdk.NewAttribute(types.AttributeKeyPositionId, string(positionId)),
		),
	)
	// 최종 결과 리턴 하기
	return &types.MsgLeverageStakeResponse{}, nil
}

func (k msgServer) MintStAsset(ctx sdk.Context, amount sdk.Int, denom string) (sdk.Coins, error) {
	stAssetDenom := types.StAssetDenomFromHostZoneDenom(denom)

	hz, _ := k.GetHostZoneFromHostDenom(ctx, denom)
	// redemption rate 를 최소 1.00000 셋팅함
	amountToMint := (sdk.NewDecFromInt(amount).Quo(hz.RedemptionRate)).TruncateInt()
	coinString := amountToMint.String() + stAssetDenom
	stCoins, err := sdk.ParseCoinsNormalized(coinString)

	if err != nil {
		k.Logger(ctx).Error("Failed to parse coins")
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Failed to parse coins %s", coinString)
	}

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, stCoins)
	if err != nil {
		k.Logger(ctx).Error("Failed to mint coins")
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Failed to mint coins")
	}

	k.Logger(ctx).Info(fmt.Sprintf("[MINT ST ASSET] success on %s.", hz.GetChainId()))

	return stCoins, nil
}
