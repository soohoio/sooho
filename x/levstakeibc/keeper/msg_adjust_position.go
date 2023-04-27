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

func (k msgServer) AdjustPosition(_ctx context.Context, req *types.MsgAdjustPosition) (*types.MsgAdjustPositionResponse, error) {

	ctx := sdk.UnwrapSDKContext(_ctx)

	// 입력 받은 denom 으로 host zone 존재 여부 확인
	hostZone, found := k.GetHostZone(ctx, req.HostDenom)

	if !found {
		return nil, errorsmod.Wrap(types.ErrHostZoneNotFound, "err : hostzone not found")
	}

	// 입력 받은 position 으로 host zone 존재 여부 확인
	position, found := k.GetPosition(ctx, req.PositionId)

	if !found {
		return nil, errorsmod.Wrap(types.ErrPositionNotFound, "err : position not found")
	}

	// position 은 존재하나 현재 ACTIVE 상태인지 체크 필요
	if position.Status != types.PositionStatus_POSITION_ACTIVE {
		return nil, errorsmod.Wrap(types.ErrPositionIsNotActive, "err : position is not ACTIVE status")
	}

	creator, _ := sdk.AccAddressFromBech32(req.Creator)

	// collateral, debt 둘 중 하나 혹은 둘 다 포용하기 위한 sdk.Int 토큰양 변수
	addedStakeAmount := sdk.ZeroInt()

	if req.Debt.GT(sdk.ZeroInt()) {
		// 빚을 더 지는 경우 Pool 에서 더 빌려오고 ZoneAddress 에 전송 후, Loan, Position 데이터에 반영
		debtAmount, err := k.addDebt(ctx, position, hostZone, req.Debt)
		if err != nil {
			return nil, errorsmod.Wrap(types.ErrPositionNotFound, "add debt func error")
		}
		// State 반영 후 추가할 stakeAmount 에 Sum
		addedStakeAmount.Add(debtAmount)
	}

	k.Logger(ctx).Info(fmt.Sprintf("Successfully done for adding debt to the existed loan data, LoanId : %v", position.LoanId))

	if req.Collateral.GT(sdk.ZeroInt()) {
		// 담보를 추가하는 경우 유저 지갑의 잔고와 입력 받은 값을 확인 후 ZoneAddress 에 전송 후, Loan, Position 데이터에 반영
		collateralAmount, err := k.addCollateral(ctx, position, hostZone, creator, req.Collateral)
		if err != nil {
			return nil, errorsmod.Wrap(types.ErrFailureAddCollateral, "add collateral func error")
		}
		addedStakeAmount.Add(collateralAmount)
	}

	k.Logger(ctx).Info(fmt.Sprintf("Successfully done for adding collateral to the existed position data, PositionId : %v", position.Id))

	// 전체 추가된 담보 + 빚 토큰 양을 stToken 으로 mint 함 > 모듈 어카운트에 존재
	stCoin, err := k.MintStAsset(ctx, addedStakeAmount, req.HostDenom)

	if err != nil {
		return nil, types.ErrMintAddedStAsset
	}

	// 추가된 stToken, NativeToken 을 Position 에 기록함
	err = k.updatePosition(ctx, position, addedStakeAmount, stCoin.AmountOf(types.StAssetDenomFromHostZoneDenom(req.HostDenom)))
	if err != nil {
		return nil, err
	}

	// save deposit record
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
	depositRecord.Amount = depositRecord.Amount.Add(addedStakeAmount)

	k.RecordsKeeper.SetDepositRecord(ctx, *depositRecord)

	return &types.MsgAdjustPositionResponse{}, nil
}

func (k msgServer) addCollateral(
	ctx sdk.Context,
	position types.Position,
	hostZone types.HostZone,
	creator sdk.AccAddress,
	collateral sdk.Int,
) (sdk.Int, error) {

	ibcDenom := hostZone.GetIbcDenom()
	coinString := collateral.String() + ibcDenom
	inCoin, err := sdk.ParseCoinNormalized(coinString)

	if err != nil {
		return sdk.ZeroInt(), fmt.Errorf("failed to parsecoin normalized for %s", coinString)
	}

	balance := k.bankKeeper.GetBalance(ctx, creator, ibcDenom)

	if balance.IsLT(inCoin) {
		k.Logger(ctx).Error(fmt.Sprintf("balance is lower than collateral amount. collateral amount: %v, balance: %v", collateral, balance.Amount))
		return sdk.ZeroInt(), errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "balance is lower than collateral amount. collateral amount: %v, balance: %v", collateral, balance.Amount)
	}

	zoneAddress, err := sdk.AccAddressFromBech32(hostZone.Address)

	if err != nil {
		return sdk.ZeroInt(), fmt.Errorf("could not bech32 decode address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}

	err = k.bankKeeper.SendCoins(ctx, creator, zoneAddress, sdk.NewCoins(inCoin))
	if err != nil {
		return sdk.ZeroInt(), fmt.Errorf("could not sendcoin to zone address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}

	err = k.LendingPoolKeeper.AddCollateral(ctx, position.LoanId, sdk.NewDecFromInt(collateral))

	if err != nil {
		return sdk.ZeroInt(), types.ErrFailureAddCollateral
	}

	return inCoin.Amount, nil
}

func (k msgServer) addDebt(
	ctx sdk.Context,
	position types.Position,
	hostZone types.HostZone,
	debt sdk.Int,
) (sdk.Int, error) {

	debtAmount := sdk.NewDecFromInt(debt)
	ibcDenom := hostZone.GetIbcDenom()
	zoneAddress, err := sdk.AccAddressFromBech32(hostZone.Address)

	if err != nil {
		return sdk.ZeroInt(), fmt.Errorf("could not bech32 decode address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}

	err = k.LendingPoolKeeper.AddDebt(ctx, position.LoanId, ibcDenom, debtAmount)

	if err != nil {
		return sdk.ZeroInt(), types.ErrFailureAddDebt
	}

	coinString := debt.String() + ibcDenom
	receivedDebt, err := sdk.ParseCoinNormalized(coinString)

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, zoneAddress, sdk.NewCoins(receivedDebt))

	if err != nil {
		return sdk.ZeroInt(), fmt.Errorf("could not sendcoin to zone address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}

	return debt, nil
}

func (k msgServer) updatePosition(ctx sdk.Context, position types.Position, addedStakeAmount sdk.Int, addedStStakeAmount sdk.Int) error {

	position.StTokenAmount = position.StTokenAmount.Add(addedStStakeAmount)
	position.NativeTokenAmount = position.NativeTokenAmount.Add(addedStakeAmount)

	k.SetPosition(ctx, position)

	return nil
}