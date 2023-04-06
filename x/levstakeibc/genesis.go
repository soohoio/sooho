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

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Params = k.GetParams(ctx)
	genesis.HostZoneList = k.GetAllHostZone(ctx)
	genesis.EpochTrackerList = k.GetAllEpochTracker(ctx)

	return genesis
}
