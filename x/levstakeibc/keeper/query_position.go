package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
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

func (k Keeper) GetAllPositionByPage(ctx sdk.Context, page *query.PageRequest) (positions []types.Position, pageRes *query.PageResponse, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PositionKey)

	pageRes, err = query.Paginate(store, page, func(key []byte, value []byte) error {
		var position types.Position
		if err := k.cdc.Unmarshal(value, &position); err != nil {
			return err
		}
		positions = append(positions, position)
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return positions, pageRes, nil
}

func (k Keeper) GetPositionListBySender(ctx sdk.Context, sender string) (positions []types.Position) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PositionKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Position
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.GetSender() == sender {
			positions = append(positions, val)
		}
	}

	return positions
}
