package keeper

import (
	"context"
	admintypes "github.com/soohoio/stayking/v3/x/admin/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v3/x/lendingpool/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the shield MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if !m.adminKeeper.IsAdmin(ctx, creator) {
		return nil, admintypes.ErrNotAdmin
	}
	res, err := m.Keeper.CreatePool(ctx, *msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMsgCreatePool,
			sdk.NewAttribute(types.AttributeTypeDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeTypePoolId, strconv.FormatUint(res.PoolId, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgCreatePoolResponse{PoolId: res.PoolId}, nil
}

func (m msgServer) DeletePool(goCtx context.Context, msg *types.MsgDeletePool) (*types.MsgDeletePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if !m.adminKeeper.IsAdmin(ctx, creator) {
		return nil, admintypes.ErrNotAdmin
	}
	res, err := m.Keeper.DeletePool(ctx, *msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMsgDeletePool,
			sdk.NewAttribute(types.AttributeTypePoolId, strconv.FormatUint(res.PoolId, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgDeletePoolResponse{PoolId: res.PoolId}, nil
}

func (m msgServer) UpdatePool(goCtx context.Context, msg *types.MsgUpdatePool) (*types.MsgUpdatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if !m.adminKeeper.IsAdmin(ctx, creator) {
		return nil, admintypes.ErrNotAdmin
	}
	res, err := m.Keeper.UpdatePool(ctx, *msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMsgUpdatePool,
			sdk.NewAttribute(types.AttributeTypePoolId, strconv.FormatUint(res.PoolId, 10)),
			sdk.NewAttribute(types.AttributeTypeDenom, msg.Denom),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgUpdatePoolResponse{PoolId: res.PoolId}, nil
}

func (m msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	res, err := m.Keeper.Deposit(ctx, *msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMsgDeposit,
			sdk.NewAttribute(types.AttributeTypeAmountIn, msg.Amount.String()),
			sdk.NewAttribute(types.AttributeTypeAmountOut, res.Amount.String()),
			sdk.NewAttribute(types.AttributeTypePoolId, strconv.FormatUint(msg.PoolId, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.From),
		),
	})
	return &res, err
}

func (m msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	res, err := m.Keeper.Withdraw(ctx, *msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMsgWithdraw,
			sdk.NewAttribute(types.AttributeTypeAmountIn, msg.Amount.String()),
			sdk.NewAttribute(types.AttributeTypeAmountOut, res.Amount.String()),
			sdk.NewAttribute(types.AttributeTypePoolId, strconv.FormatUint(msg.PoolId, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.From),
		),
	})
	return &res, err
}

func (m msgServer) Liquidate(goCtx context.Context, msg *types.MsgLiquidate) (*types.MsgLiquidateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	creator, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}
	if !m.adminKeeper.IsAdmin(ctx, creator) {
		return nil, admintypes.ErrNotAdmin
	}
	m.Keeper.Liquidate(ctx, msg.LoanId)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMsgLiquidate,
			sdk.NewAttribute(types.AttributeTypeLoanId, strconv.FormatUint(msg.LoanId, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.From),
		),
	})
	return &types.MsgLiquidateResponse{}, nil
}
