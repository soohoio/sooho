package interchainquery

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v3/x/interchainquery/keeper"
	"github.com/soohoio/stayking/v3/x/interchainquery/types"
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
	k.Logger(ctx).Info(fmt.Sprintf("[Genesis Debug] params channelid:%v Poolid:%v RoutespoolId:%v TokenInDenom:%v TokenOut:%v ", genState.Params.PriceQueryChannelId, genState.Params.PriceQueryPoolId, genState.Params.PriceQueryRoutesPoolId, genState.Params.PriceQueryTokenInDenom, genState.Params.PriceQueryTokenOut))
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Queries: k.AllQueries(ctx),
	}
}
