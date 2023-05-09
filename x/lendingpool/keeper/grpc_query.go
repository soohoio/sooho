package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v2/x/lendingpool/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (k Keeper) Pool(c context.Context, request *types.QueryPoolRequest) (*types.QueryPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	pool, found := k.GetPool(ctx, request.PoolId)
	if !found {
		return nil, types.ErrPoolNotFound
	}
	return &types.QueryPoolResponse{
		Pool: pool,
	}, nil
}

func (k Keeper) Pools(c context.Context, request *types.QueryPoolsRequest) (*types.QueryPoolsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	pools := k.GetAllPools(ctx)
	return &types.QueryPoolsResponse{
		Pools: pools,
	}, nil
}

func (k Keeper) Loan(c context.Context, request *types.QueryLoanRequest) (*types.QueryLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	loan, found := k.GetLoan(ctx, request.LoanId)
	if !found {
		return nil, types.ErrPoolNotFound
	}
	return &types.QueryLoanResponse{
		Loan: loan,
	}, nil
}

func (k Keeper) Loans(c context.Context, request *types.QueryLoansRequest) (*types.QueryLoansResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	loans := k.GetAllLoans(ctx)
	return &types.QueryLoansResponse{
		Loans: loans,
	}, nil
}

func (k Keeper) APR(c context.Context, request *types.QueryAPRRequest) (*types.QueryAPRResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	pool, found := k.GetPool(ctx, request.PoolId)
	if !found {
		return nil, types.ErrPoolNotFound
	}
	return &types.QueryAPRResponse{
		Apr: pool.GetInterestModel().GetAPR(pool.GetUtilizationRate()),
	}, nil
}

func (k Keeper) Params(c context.Context, request *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)
	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}
