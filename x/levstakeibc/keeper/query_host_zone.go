package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllHostZone(c context.Context, req *types.QueryAllHostZoneRequest) (*types.QueryAllHostZoneResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request..")
	}
	ctx := sdk.UnwrapSDKContext(c)
	k.Logger(ctx).Info("before GetAllHostZone")

	hostZones, pageRes, err := k.GetAllHostZone(ctx, req.Pagination)
	k.Logger(ctx).Info("after GetAllHostZone")

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	k.Logger(ctx).Info("after GetAllHostZone2")

	return &types.QueryAllHostZoneResponse{
		HostZone:   hostZones,
		Pagination: pageRes,
	}, nil
}

func (k Keeper) HostZone(c context.Context, req *types.QueryGetHostZoneRequest) (*types.QueryGetHostZoneResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request..")
	}

	ctx := sdk.UnwrapSDKContext(c)
	hostZone, found := k.GetHostZone(ctx, req.ChainId)

	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetHostZoneResponse{HostZone: hostZone}, nil
}

/**
 * Store Getter / Setter
 */

func (k Keeper) GetHostZone(ctx sdk.Context, chainId string) (hostZone types.HostZone, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.HostZoneKey))
	b := store.Get([]byte(chainId))
	if b == nil {
		return hostZone, false
	}
	k.cdc.MustUnmarshal(b, &hostZone)
	return hostZone, true
}

func (k Keeper) GetAllHostZone(ctx sdk.Context, page *query.PageRequest) (hostZones []types.HostZone, pageRes *query.PageResponse, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.HostZoneKey))

	pageRes, err = query.Paginate(store, page, func(key []byte, value []byte) error {
		var hostZone types.HostZone
		if err := k.cdc.Unmarshal(value, &hostZone); err != nil {
			return err
		}
		hostZones = append(hostZones, hostZone)
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return hostZones, pageRes, nil
}
