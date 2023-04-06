package interchainquery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v2/x/interchainquery/keeper"
	"github.com/soohoio/stayking/v2/x/interchainquery/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// set registered zones info from genesis
	k.SetPort(ctx, genState.PortId)
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	for _, query := range genState.Queries {
		// Initialize empty epoch values via Cosmos SDK
		k.SetQuery(ctx, query)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Queries: k.AllQueries(ctx),
	}
}
