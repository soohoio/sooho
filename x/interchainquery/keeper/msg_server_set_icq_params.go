package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/interchainquery/types"
)

func (k msgServer) SetIcqParams(goCtx context.Context, msg *types.MsgSetIcqParams) (*types.MsgSetIcqParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	params.PriceQueryChannelId = msg.PriceQueryChannelId
	params.PriceQueryPoolId = msg.PriceQueryPoolId
	params.PriceQueryRoutesPoolId = msg.PriceQueryRoutesPoolId
	params.PriceQueryTokenInDenom = msg.PriceQueryTokenInDenom
	params.PriceQueryTokenOut = msg.PriceQueryTokenOut
	params.PriceQueryPath = msg.PriceQueryPath
	k.SetParams(ctx, params)
	return &types.MsgSetIcqParamsResponse{}, nil
}
