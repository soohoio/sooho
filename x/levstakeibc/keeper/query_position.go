package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PositionListBySender(_ctx context.Context, req *types.QueryGetPositionListBySenderRequest) (*types.QueryGetPositionListBySenderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(_ctx)

	positions := k.GetPositionListBySender(ctx, req.Sender)

	return &types.QueryGetPositionListBySenderResponse{
		Position: positions,
	}, nil

}

func (k Keeper) AllPosition(_ctx context.Context, req *types.QueryAllPositionRequest) (*types.QueryAllPositionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(_ctx)
	positions, pageRes, err := k.GetAllPositionByPage(ctx, req.Pagination)

	if err != nil {
		return nil, err
	}

	return &types.QueryAllPositionResponse{
		Position:   positions,
		Pagination: pageRes,
	}, nil
}
