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
			return nil, err
		}
		return msg, nil
	} else if levType == types.StakingType_LEVERAGE_TYPE {
		msg, err := k.stakeWithLeverage(ctx, equity, hostDenom, msg.Creator, leverageRatio, levType, msg.LendingPoolDenom)
		if err != nil {
			return nil, err
		}
		return msg, nil
	}

	return nil, errorsmod.Wrapf(types.ErrInvalidLeverageRatio, "invalid leverage type value (lev ratio %v) ", leverageRatio)
}

func (k msgServer) stakeWithoutLeverage(ctx sdk.Context, equity sdk.Int, hostDenom string, creator string) (*types.MsgLeverageStakeResponse, error) {

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
		return nil, errorsmod.Wrapf(err, "failed to parse coin (%s)", coinString)
	}

	balance := k.bankKeeper.GetBalance(ctx, sender, ibcDenom)

	if balance.IsLT(inCoin) {
		k.Logger(ctx).Error(fmt.Sprintf("balance is lower than staking amount. staking amount: %v, balance: %v", equity, balance.Amount))
		return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "balance is lower than staking amount. staking amount: %v, balance: %v", equity, balance.Amount)
	}

	// check that the token is an IBC token
	isIbcToken := types.IsIBCToken(ibcDenom)
	if !isIbcToken {
		k.Logger(ctx).Error("invalid token denom - denom is not an IBC token (%s)", ibcDenom)
		return nil, errorsmod.Wrapf(types.ErrInvalidToken, "denom is not an IBC token (%s)", ibcDenom)
	}

	zoneAddress, err := sdk.AccAddressFromBech32(hostZone.Address)
	k.Logger(ctx).Info(fmt.Sprintf("[DEBUG]hostZone Address : %v", zoneAddress))
	moduleAddress := k.accountKeeper.GetModuleAddress(types.ModuleName)
	k.Logger(ctx).Info(fmt.Sprintf("[DEBUG] modulee Address : %v", moduleAddress.String()))
	if err != nil {
		return nil, fmt.Errorf("could not bech32 decode address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}

	err = k.bankKeeper.SendCoins(ctx, sender, zoneAddress, sdk.NewCoins(inCoin))

	if err != nil {
		k.Logger(ctx).Error("failed to send tokens from Account to ZoneAddress")
		return nil, errorsmod.Wrap(err, "failed to send tokens from Account to ZoneAddress")
	}

	// mint user `amount` of the corresponding stAsset
	// NOTE: We should ensure that denoms are unique - we don't want anyone spoofing denoms
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

// TODO: Not Stake 와 중복되는 로직임... 리팩토링 필요
func (k msgServer) stakeWithLeverage(ctx sdk.Context, equity sdk.Int, denom string, creator string, ratio sdk.Dec, levType types.StakingType, receiver string) (*types.MsgLeverageStakeResponse, error) {
	k.Logger(ctx).Info("leverageType Mode ... ")
	k.Logger(ctx).Info(fmt.Sprintf("stakeWithLeverage => equity: %v, denom: %v, creator: %v, ratio: %v, reverageType: %v, markPriceBaseDenom: %v", equity, denom, creator, ratio, levType, receiver))

	moduleAddress := k.accountKeeper.GetModuleAddress(types.ModuleName)
	k.Logger(ctx).Info(fmt.Sprintf("[DEBUG] module Address : %v", moduleAddress.String()))
	k.Logger(ctx).Info(fmt.Sprintf("[DEBUG] lendingpool module address : %v", k.accountKeeper.GetModuleAddress(lendingpooltypes.ModuleName).String()))

	sender, _ := sdk.AccAddressFromBech32(creator)

	hostZone, found := k.GetHostZoneByHostDenom(ctx, denom)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrHostZoneNotFound, "not found : hostzone")
	}

	existsPostion, found := k.GetPositionByDenomAndSender(ctx, denom, creator)

	if found {
		return nil, errorsmod.Wrapf(types.ErrAlreadyExistsPosition, fmt.Sprintf("Exists the position Id : %v and denom : %v", existsPostion.Id, denom))
	}

	borrowingAmount := sdk.NewDecFromInt(equity).Mul(ratio.Sub(sdk.NewDec(1))).TruncateInt()
	totalAsset := equity.Add(borrowingAmount)
	debtRatio := sdk.NewDecFromInt(borrowingAmount).Quo(sdk.NewDecFromInt(totalAsset))

	loanId, err := k.LendingPoolKeeper.Borrow(
		ctx,
		hostZone.GetIbcDenom(),
		types.ModuleName,
		sender,
		sdk.NewCoins(sdk.NewCoin(hostZone.GetIbcDenom(), borrowingAmount)),
		sdk.NewCoins(sdk.NewCoin(hostZone.GetIbcDenom(), equity)),
	)

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrLendingPoolBorrow, "unexpected error : lendingpool keeper borrow func")
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

	// x/levstakeibc store 객체에 Position 객체를 생성하여 저장한다.
	positionId := k.GetNextPositionID(ctx)

	position := types.Position{
		Id:                positionId,
		LoanId:            loanId,
		Sender:            creator,
		Denom:             denom,
		StTokenAmount:     stCoins.AmountOf(types.StAssetDenomFromHostZoneDenom(denom)),
		NativeTokenAmount: totalAsset,
		Status:            types.PositionStatus_POSITION_ACTIVE,
		Liquidated:        false,
	}

	k.SetPosition(ctx, position)

	k.SetNextPositionID(ctx, positionId+1) // 다음 포지션 ID ++

	k.Logger(ctx).Info(fmt.Sprintf("Successfully done for saving position data, PositionId : %v", positionId))

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

	// 최종 결과 리턴 하기
	return &types.MsgLeverageStakeResponse{}, nil
}

func (k msgServer) MintStAsset(ctx sdk.Context, amount sdk.Int, denom string) (sdk.Coins, error) {
	stAssetDenom := types.StAssetDenomFromHostZoneDenom(denom)

	hz, _ := k.GetHostZoneFromHostDenom(ctx, denom)
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
