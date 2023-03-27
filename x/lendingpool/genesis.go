package lendingpool

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/lendingpool/keeper"
	"github.com/soohoio/stayking/v2/x/lendingpool/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	// set pools and borrowers
	for _, p := range data.Pools {
		k.SetPool(ctx, p)
	}

	// set trackers
	k.SetNextPoolID(ctx, data.NextPoolId)

	// set params
	k.SetTaxRate(ctx, data.Params.ProtocolTaxRate)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	params := types.Params{
		ProtocolTaxRate: k.GetTaxRate(ctx),
	}
	pools := k.GetAllPools(ctx)
	nextPoolID := k.GetNextPoolID(ctx)
	return types.GenesisState{
		Params:     params,
		Pools:      pools,
		NextPoolId: nextPoolID,
	}
}
