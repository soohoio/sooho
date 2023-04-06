package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/records/types"
)

func (m msgServer) UpdateDenomPrice(_ctx context.Context, req *types.MsgUpdateDenomPrice) (*types.MsgUpdateDenomPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(_ctx)

	m.Logger(ctx).Info("hereee!!!!!!!!!!!!!!!!!!!!")

	m.SetDenomPriceRecord(ctx, types.DenomPriceRecord{
		BaseDenom:   req.BaseDenom,
		TargetDenom: req.TargetDenom,
		DenomPrice:  req.DenomPrice,
		Timestamp:   ctx.BlockTime().UnixNano(),
	})
	m.Logger(ctx).Info("hereee!!!!!!!!!!!!!!!!!!!!")

	return &types.MsgUpdateDenomPriceResponse{}, nil
}
