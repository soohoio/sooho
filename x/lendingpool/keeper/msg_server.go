package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/lendingpool/types"
)

type msgServer struct {
	Keeper
}

func (m msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	res, err := m.Keeper.CreatePool(ctx, *msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMsgCreatePool,
			sdk.NewAttribute(types.AttributeTypeDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeTypeId, strconv.FormatUint(res.PoolId, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgCreatePoolResponse{}, nil
}

func (m msgServer) Deposit(ctx context.Context, deposit *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) Withdraw(ctx context.Context, withdraw *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewMsgServerImpl returns an implementation of the shield MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
