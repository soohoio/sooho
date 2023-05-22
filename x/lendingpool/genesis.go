package lendingpool

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v3/x/lendingpool/keeper"
	"github.com/soohoio/stayking/v3/x/lendingpool/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	// set pools and borrowers
	for _, p := range data.Pools {
		k.SetPool(ctx, p)
	}

	// set trackers
	k.SetNextPoolID(ctx, data.NextPoolId)
	k.SetNextLoanID(ctx, data.NextLoanId)

	// set params
	k.SetParams(ctx, data.Params)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	params := k.GetParams(ctx)
	pools := k.GetAllPools(ctx)
	nextPoolID := k.GetNextPoolID(ctx)
	return types.GenesisState{
		Params:     params,
		Pools:      pools,
		NextPoolId: nextPoolID,
	}
}
