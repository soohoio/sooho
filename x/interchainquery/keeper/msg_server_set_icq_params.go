package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/interchainquery/types"
)

func (k msgServer) SetIcqParams(goCtx context.Context, msg *types.MsgSetIcqParams) (*types.MsgSetIcqParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info(fmt.Sprintf("[DEBUG] Set Icq Params- channel-id: %v, poolId: %v, routesPoolId: %v, TokenInDenom: %v, TokenOut: %v ", msg.PriceQueryChannelId, msg.PriceQueryPoolId, msg.PriceQueryRoutesPoolId, msg.PriceQueryTokenInDenom, msg.PriceQueryTokenOut))
	params := k.GetParams(ctx)
	k.Logger(ctx).Info(fmt.Sprintf("1 params channelid:%v Poolid:%v RoutespoolId:%v TokenInDenom:%v TokenOut:%v ", params.PriceQueryChannelId, params.PriceQueryPoolId, params.PriceQueryRoutesPoolId, params.PriceQueryTokenInDenom, params.PriceQueryTokenOut))
	params.PriceQueryChannelId = msg.PriceQueryChannelId
	params.PriceQueryPoolId = msg.PriceQueryPoolId
	params.PriceQueryRoutesPoolId = msg.PriceQueryRoutesPoolId
	params.PriceQueryTokenInDenom = msg.PriceQueryTokenInDenom
	params.PriceQueryTokenOut = msg.PriceQueryTokenOut
	params.PriceQueryPath = msg.PriceQueryPath
	k.Logger(ctx).Info(fmt.Sprintf("2 params channelid:%v Poolid:%v RoutespoolId:%v TokenInDenom:%v TokenOut:%v ", params.PriceQueryChannelId, params.PriceQueryPoolId, params.PriceQueryRoutesPoolId, params.PriceQueryTokenInDenom, params.PriceQueryTokenOut))
	k.SetParams(ctx, params)
	return &types.MsgSetIcqParamsResponse{}, nil
}
