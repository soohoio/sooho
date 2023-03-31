package levstakeibc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v2/x/levstakeibc/keeper"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, hostZone := range genState.HostZoneList {
		k.SetHostZone(ctx, hostZone)
	}
	for _, epochTracker := range genState.EpochTrackerList {
		k.SetEpochTracker(ctx, epochTracker)
	}

	k.SetParams(ctx, genState.Params)
}

func ExportGenesis(context sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesisState := types.DefaultGenesis()

	return genesisState
}
