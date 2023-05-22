package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v3/x/admin/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (q Querier) Admins(c context.Context, request *types.QueryAdminsRequest) (*types.QueryAdminsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	admins := q.GetAdmins(ctx)
	return &types.QueryAdminsResponse{
		Admins: admins,
	}, nil
}
