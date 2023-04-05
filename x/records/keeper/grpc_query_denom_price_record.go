package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/soohoio/stayking/v2/x/records/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DenomPriceRecordAll(_ctx context.Context, req *types.QueryAllDenomPriceRecordRequest) (*types.QueryAllDenomPriceRecordResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(_ctx)

	var denomPriceRecords []types.DenomPriceRecord
	denomPriceRecordStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomPriceRecordKey))

	pageReq, err := query.Paginate(denomPriceRecordStore, req.Pagination, func(key []byte, value []byte) error {
		var denomPriceRecord types.DenomPriceRecord
		if err := k.Cdc.Unmarshal(value, &denomPriceRecord); err != nil {
			return err
		}
		denomPriceRecords = append(denomPriceRecords, denomPriceRecord)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDenomPriceRecordResponse{
		DenomPriceRecord: denomPriceRecords,
		Pagination:       pageReq,
	}, nil
}

func (k Keeper) DenomPriceRecord(_ctx context.Context, req *types.QueryGetDenomPriceRecordRequest) (*types.QueryGetDenomPriceRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(_ctx)
	denomId := req.BaseDenom + "-" + req.GetTargetDenom()
	denomPriceRecord, found := k.GetDenomPriceRecord(ctx, denomId)

	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetDenomPriceRecordResponse{
		DenomPriceRecord: denomPriceRecord,
	}, nil
}
