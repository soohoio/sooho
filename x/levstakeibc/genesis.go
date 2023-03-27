package levstakeibc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v2/x/levstakeibc/keeper"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func InitGenesis(context sdk.Context, k keeper.Keeper, state types.GenesisState) {
}

func ExportGenesis(context sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesisState := types.DefaultGenesis()

	return genesisState
}
