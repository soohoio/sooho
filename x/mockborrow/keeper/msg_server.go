package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v3/x/mockborrow/types"
	"strconv"
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

func (m msgServer) Borrow(goCtx context.Context, msg *types.MsgBorrow) (*types.MsgBorrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	res, err := m.Keeper.Borrow(ctx, msg.Denom, fromAddr, msg.BorrowAmount, msg.Collateral)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBorrow,
			sdk.NewAttribute(types.AttributeTypeFrom, msg.From),
			sdk.NewAttribute(types.AttributeTypeDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeTypeLoanId, strconv.FormatUint(res, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.From),
		),
	})

	return &types.MsgBorrowResponse{Id: res}, nil
}

func (m msgServer) Repay(goCtx context.Context, msg *types.MsgRepay) (*types.MsgRepayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	res, err := m.Keeper.Repay(ctx, msg.Id, msg.Amount)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRepay,
			sdk.NewAttribute(types.AttributeTypeFrom, msg.From),
			sdk.NewAttribute(types.AttributeTypeLoanId, strconv.FormatUint(msg.Id, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.From),
		),
	})

	return &types.MsgRepayResponse{Change: res}, nil
}
