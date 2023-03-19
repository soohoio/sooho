package keeper

import (
	"context"

	"github.com/soohoio/stayking/v2/x/lendingpool/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Pool(ctx context.Context, request *types.QueryPoolRequest) (*types.QueryPoolResponse, error) {
	//TODO implement me
	panic("implement me")
}
