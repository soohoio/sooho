package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	epochtypes "github.com/soohoio/stayking/v2/x/epochs/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"strconv"
)

func (k msgServer) AdjustPosition(_ctx context.Context, req *types.MsgAdjustPosition) (*types.MsgAdjustPositionResponse, error) {

	ctx := sdk.UnwrapSDKContext(_ctx)
	k.Logger(ctx).Info(" Adjust Position ....")
	k.Logger(ctx).Info("Params : " + req.String())
	k.Logger(ctx).Info("Params : " + req.Creator + " , " + strconv.FormatUint(req.PositionId, 10) + " , " + req.HostDenom + " , " + req.Collateral.String() + " , " + req.Debt.String())

	// 입력 받은 denom 으로 host zone 존재 여부 확인
	hostZone, found := k.GetHostZoneByHostDenom(ctx, req.HostDenom)

	if !found {
		return nil, errorsmod.Wrapf(types.ErrHostZoneNotFound, "host zone not found by host denom %v", req.HostDenom)
	}

	// 입력 받은 position 으로 host zone 존재 여부 확인
	position, found := k.GetPosition(ctx, req.PositionId)

	if !found {
		printErr := fmt.Sprintf("position not found by position id (%v)", req.PositionId)
		k.Logger(ctx).Error("[CUSTOM DEBUG] " + printErr)
		return nil, errorsmod.Wrap(types.ErrPositionNotFound, printErr)
	}

	// position 은 존재하나 현재 POSITION_UNBONDING_IN_PROGRESS 상태일 때는 adjust 가 안된다.
	if position.Status == types.PositionStatus_POSITION_UNBONDING_IN_PROGRESS {
		printErr := fmt.Sprintf("position status (=POSITION_UNBONDING_IN_PROGRESS) %v is not available to adjust", req.PositionId)
		k.Logger(ctx).Error("[CUSTOM DEBUG] " + printErr)
		return nil, errorsmod.Wrap(types.ErrPositionIsNotActive, printErr)
	}

	creator, _ := sdk.AccAddressFromBech32(req.Creator)

	// collateral, debt 둘 중 하나 혹은 둘 다 포용하기 위한 sdk.Int 토큰양 변수
	addedStakeAmount := sdk.ZeroInt()

	if req.Debt.GT(sdk.ZeroInt()) {
		// 빚을 더 지는 경우 Pool 에서 더 빌려오고 ZoneAddress 에 전송 후, Loan, Position 데이터에 반영
		debtAmount, err := k.addDebt(ctx, position, hostZone, req.Debt)
		if err != nil {
			k.Logger(ctx).Error("[CUSTOM DEBUG] " + err.Error())
			return nil, errorsmod.Wrap(types.ErrFailureOperatePosition, err.Error())
		}
		// State 반영 후 추가할 stakeAmount 에 Sum
		addedStakeAmount = addedStakeAmount.Add(debtAmount)
		k.Logger(ctx).Info(fmt.Sprintf("[CUSTOM DEBUG] debtAmount : %v , ", debtAmount))
	}

	k.Logger(ctx).Info(fmt.Sprintf("[CUSTOM DEBUG] Successfully done for adding debt to the existed loan data, LoanId : %v", position.LoanId))

	if req.Collateral.GT(sdk.ZeroInt()) {
		// 담보를 추가하는 경우 유저 지갑의 잔고와 입력 받은 값을 확인 후 ZoneAddress 에 전송 후, Loan, Position 데이터에 반영
		collateralAmount, err := k.addCollateral(ctx, position, hostZone, creator, req.Collateral)
		if err != nil {
			k.Logger(ctx).Error("[CUSTOM DEBUG] " + err.Error())
			return nil, errorsmod.Wrap(types.ErrFailureOperatePosition, err.Error())
		}
		addedStakeAmount = addedStakeAmount.Add(collateralAmount)
		k.Logger(ctx).Info(fmt.Sprintf("[CUSTOM DEBUG] collateralAmount : %v , ", collateralAmount))
	}

	k.Logger(ctx).Info(fmt.Sprintf("[CUSTOM DEBUG] Successfully done for adding collateral to the existed position data, PositionId : %v", position.Id))

	// 전체 추가된 담보 + 빚 토큰 양을 stToken 으로 mint 함 > 모듈 어카운트에 존재
	stCoin, err := k.MintStAsset(ctx, addedStakeAmount, req.HostDenom)
	k.Logger(ctx).Info(fmt.Sprintf("[CUSTOM DEBUG] stCoin : %v , ", stCoin.AmountOf(types.StAssetDenomFromHostZoneDenom(req.HostDenom))))

	if err != nil {
		printErr := fmt.Sprintf("failed to mint st tokens")
		k.Logger(ctx).Error("[CUSTOM DEBUG] " + printErr)
		return nil, errorsmod.Wrap(types.ErrFailureMintStAsset, printErr)
	}
	zoneAddress, err := sdk.AccAddressFromBech32(hostZone.Address)
	if err != nil {
		printErr := fmt.Sprintf("could not bech32 decode address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
		k.Logger(ctx).Error("[CUSTOM DEBUG] " + printErr)
		return nil, errorsmod.Wrap(types.ErrInvalidAccount, printErr)
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, zoneAddress, stCoin)
	if err != nil {
		printErr := fmt.Sprintf("failed to mint %s stAssets to host address", hostZone.HostDenom)
		k.Logger(ctx).Error("[CUSTOM DEBUG] " + printErr)
		return nil, errorsmod.Wrap(err, printErr)
	}

	// save deposit record
	staykingEpochTracker, found := k.GetEpochTracker(ctx, epochtypes.STAYKING_EPOCH)

	if !found {
		printErr := fmt.Sprintf("no epoch number for epoch (%s)", epochtypes.STAYKING_EPOCH)
		k.Logger(ctx).Error("[CUSTOM DEBUG] " + printErr)
		return nil, errorsmod.Wrap(sdkerrors.ErrNotFound, printErr)
	}

	depositRecord, found := k.RecordsKeeper.GetDepositRecordByEpochAndChain(ctx, staykingEpochTracker.EpochNumber, hostZone.ChainId)
	if !found {
		printErr := fmt.Sprintf("no deposit record for epoch (%d)", staykingEpochTracker.EpochNumber)
		k.Logger(ctx).Error("[CUSTOM DEBUG] " + printErr)
		return nil, errorsmod.Wrap(sdkerrors.ErrNotFound, printErr)
	}
	depositRecord.Amount = depositRecord.Amount.Add(addedStakeAmount)

	k.RecordsKeeper.SetDepositRecord(ctx, *depositRecord)

	// 추가된 stToken, NativeToken 을 Position 에 기록하고 상태를 다시 Pending 상태로 돌린다.
	position.StTokenAmount = position.StTokenAmount.Add(stCoin.AmountOf(types.StAssetDenomFromHostZoneDenom(req.HostDenom)))
	position.NativeTokenAmount = position.NativeTokenAmount.Add(addedStakeAmount)
	position.DepositRecordId = depositRecord.Id
	position.Status = types.PositionStatus_POSITION_PENDING

	k.SetPosition(ctx, position)

	k.Logger(ctx).Info(fmt.Sprintf("[CUSTOM DEBUG] Successfully adjusted position (positionId: %v) ", position.Id))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAdjustPosition,
			sdk.NewAttribute(types.AttributeKeyRecipientChain, hostZone.ChainId),
			sdk.NewAttribute(types.AttributeKeyAddress, req.Creator),
			sdk.NewAttribute(types.AttributeKeyColleteralTokenAmount, req.Collateral.String()),
			sdk.NewAttribute(types.AttributeKeyDebtTokenAmount, req.Debt.String()),
		),
	)
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
		return sdk.ZeroInt(), errorsmod.Wrapf(types.ErrInvalidToken, "failed to parsecoin normalized for %s", coinString)
	}

	balance := k.bankKeeper.GetBalance(ctx, creator, ibcDenom)

	if balance.IsLT(inCoin) {
		return sdk.ZeroInt(), errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "balance is lower than collateral amount. collateral amount: %v, balance: %v", collateral, balance.Amount)
	}

	zoneAddress, err := sdk.AccAddressFromBech32(hostZone.Address)

	if err != nil {
		return sdk.ZeroInt(), errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "could not bech32 decode address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}

	err = k.bankKeeper.SendCoins(ctx, creator, zoneAddress, sdk.NewCoins(inCoin))

	if err != nil {
		return sdk.ZeroInt(), errorsmod.Wrapf(types.ErrFailureSendToken, "could not sendcoin to zone address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}

	err = k.LendingPoolKeeper.AddCollateral(ctx, position.LoanId, sdk.NewDecFromInt(collateral))

	if err != nil {
		return sdk.ZeroInt(), err
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
		return sdk.ZeroInt(), errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "could not bech32 decode address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}

	err = k.LendingPoolKeeper.AddDebt(ctx, position.LoanId, ibcDenom, debtAmount)

	if err != nil {
		return sdk.ZeroInt(), err
	}

	coinString := debt.String() + ibcDenom
	receivedDebt, err := sdk.ParseCoinNormalized(coinString)

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, zoneAddress, sdk.NewCoins(receivedDebt))

	if err != nil {
		return sdk.ZeroInt(), errorsmod.Wrapf(types.ErrFailureSendToken, "could not sendcoin to zone address %s of zone with id: %s", hostZone.Address, hostZone.ChainId)
	}

	return debt, nil
}
