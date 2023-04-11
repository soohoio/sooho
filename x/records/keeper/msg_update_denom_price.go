package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/records/types"
)

func (m msgServer) UpdateDenomPrice(_ctx context.Context, req *types.MsgUpdateDenomPrice) (*types.MsgUpdateDenomPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(_ctx)

	m.SetDenomPriceRecord(ctx, types.DenomPriceRecord{
		BaseDenom:   req.BaseDenom,
		TargetDenom: req.TargetDenom,
		DenomPrice:  req.DenomPrice,
		Timestamp:   uint64(ctx.BlockTime().UnixNano()),
	})

	return &types.MsgUpdateDenomPriceResponse{}, nil
}
