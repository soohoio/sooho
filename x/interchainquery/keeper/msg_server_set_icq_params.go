package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	admintypes "github.com/soohoio/stayking/v3/x/admin/types"
	"github.com/soohoio/stayking/v3/x/interchainquery/types"
)

func (k msgServer) SetIcqParams(goCtx context.Context, msg *types.MsgSetIcqParams) (*types.MsgSetIcqParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	// admin address check
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if !k.AdminKeeper.IsAdmin(ctx, creator) {
		return nil, admintypes.ErrNotAdmin
	}

	params.PriceQueryChannelId = msg.PriceQueryChannelId
	params.PriceQueryPoolId = msg.PriceQueryPoolId
	params.PriceQueryRoutesPoolId = msg.PriceQueryRoutesPoolId
	params.PriceQueryTokenInDenom = msg.PriceQueryTokenInDenom
	params.PriceQueryTokenOut = msg.PriceQueryTokenOut
	params.PriceQueryPath = msg.PriceQueryPath
	k.SetParams(ctx, params)
	return &types.MsgSetIcqParamsResponse{}, nil
}
