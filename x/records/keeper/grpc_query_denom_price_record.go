package keeper

import (
	"context"
	"github.com/soohoio/stayking/v2/x/records/types"
)

func (k Keeper) DenomPriceRecordAll(ctx context.Context, req *types.QueryAllDenomPriceRecordRequest) (*types.QueryAllDenomPriceRecordResponse, error) {

	return &types.QueryAllDenomPriceRecordResponse{
		DenomPriceRecord: []types.DenomPriceRecord{},
		Pagination:       nil,
	}, nil
}

func (k Keeper) DenomPriceRecord(ctx context.Context, req *types.QueryGetDenomPriceRecordRequest) (*types.QueryGetDenomPriceRecordResponse, error) {
	return &types.QueryGetDenomPriceRecordResponse{
		DenomPriceRecord: types.DenomPriceRecord{},
	}, nil
}
