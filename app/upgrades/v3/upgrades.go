package v3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/soohoio/stayking/v2/app"
	"github.com/soohoio/stayking/v2/app/upgrades"
	levstakeibctypes "github.com/soohoio/stayking/v2/x/levstakeibc/types"
	stakeibckeeper "github.com/soohoio/stayking/v2/x/stakeibc/keeper"
	stakeibctypes "github.com/soohoio/stayking/v2/x/stakeibc/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	bpm upgrades.BaseAppParamManager,
	app *app.StayKingApp,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		// TODO: implement store preprocessing
		scopedLevstakeibcKeeper := app.CapabilityKeeper.ScopeToModule(stakeibctypes.ModuleName)
		stakeibcKeeper := stakeibckeeper.NewKeeper(
			app.AppCodec(),
			app.GetKey(stakeibctypes.ModuleName),
			app.GetMemKey(stakeibctypes.ModuleName),
			app.GetSubspace(stakeibctypes.ModuleName),
			app.AccountKeeper,
			app.BankKeeper,
			app.ICAControllerKeeper,
			*app.IBCKeeper,
			scopedLevstakeibcKeeper,
			app.InterchainqueryKeeper,
			app.RecordsKeeper,
			app.StakingKeeper,
			app.IcacallbacksKeeper)
		// skip levstakeibc initgenesis
		vm[levstakeibctypes.ModuleName] = mm.Modules[levstakeibctypes.ModuleName].ConsensusVersion()
		hostZones := stakeibcKeeper.GetAllHostZone(ctx)
		params := stakeibcKeeper.GetParams(ctx)
		dg := stakeibctypes.DefaultGenesis()

		defaultGenesis := levstakeibctypes.DefaultGenesis()

		genesis := levstakeibctypes.GenesisState{
			Params:           levstakeibctypes.Params{},
			PortId:           "",
			HostZoneList:     nil,
			EpochTrackerList: nil,
			NextPositionId:   0,
			PositionList:     nil,
		}

		ctx.Logger().Info("start to run module migrations...")

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
