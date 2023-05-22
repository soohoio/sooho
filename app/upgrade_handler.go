package app

import (
	"fmt"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	v3 "github.com/soohoio/stayking/v3/app/upgrades/v3"
	admintypes "github.com/soohoio/stayking/v3/x/admin/types"
	lendingpooltypes "github.com/soohoio/stayking/v3/x/lendingpool/types"
	"github.com/soohoio/stayking/v3/x/levstakeibc"
	levstakeibctypes "github.com/soohoio/stayking/v3/x/levstakeibc/types"
	mockborrowtypes "github.com/soohoio/stayking/v3/x/mockborrow/types"
	stakeibckeeper "github.com/soohoio/stayking/v3/x/stakeibc/keeper"
	stakeibctypes "github.com/soohoio/stayking/v3/x/stakeibc/types"
)

func (app *StayKingApp) setupUpgradeHandlers() {
	// v3 upgrade
	app.UpgradeKeeper.SetUpgradeHandler(
		v3.UpgradeName,
		func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
			delete(vm, stakeibctypes.StoreKey)
			ctx.Logger().Info("applying v3.0.0 upgrade...")
			// TODO: implement store preprocessing
			scopedStakeibcKeeper := app.CapabilityKeeper.ScopeToModule(stakeibctypes.ModuleName)
			stakeibcKeeper := stakeibckeeper.NewKeeper(
				app.AppCodec(),
				app.GetKey(stakeibctypes.ModuleName),
				app.GetMemKey(stakeibctypes.ModuleName),
				app.GetSubspace(stakeibctypes.ModuleName),
				app.AccountKeeper,
				app.BankKeeper,
				app.ICAControllerKeeper,
				*app.IBCKeeper,
				scopedStakeibcKeeper,
				app.InterchainqueryKeeper,
				app.RecordsKeeper,
				app.StakingKeeper,
				app.IcacallbacksKeeper)
			// skip levstakeibc initgenesis
			vm[levstakeibctypes.ModuleName] = app.mm.Modules[levstakeibctypes.ModuleName].ConsensusVersion()
			hostZones := stakeibcKeeper.GetAllHostZone(ctx)
			params := stakeibcKeeper.GetParams(ctx)
			epochTrackers := stakeibcKeeper.GetAllEpochTracker(ctx)

			// migrate stakeibc capability to levstakeibc
			maxIndex := app.CapabilityKeeper.GetLatestIndex(ctx)
			var i uint64
			for ; i < maxIndex; i++ {
				capOwners, _ := app.CapabilityKeeper.GetOwners(ctx, i)
				for j, o := range capOwners.Owners {
					if o.Module == stakeibctypes.ModuleName {
						capOwners.Owners[j].Module = levstakeibctypes.ModuleName
					}
				}
				app.CapabilityKeeper.SetOwners(ctx, i, capOwners)
			}

			stakeIBCAddr := app.AccountKeeper.GetModuleAddress(stakeibctypes.ModuleName)
			stakeIBCBalance := app.BankKeeper.GetAllBalances(ctx, stakeIBCAddr)
			err := app.BankKeeper.SendCoinsFromAccountToModule(ctx, stakeIBCAddr, levstakeibctypes.ModuleName, stakeIBCBalance)
			if err != nil {
				panic(err)
			}

			for _, hz := range hostZones {
				stakeibcKeeper.RemoveHostZone(ctx, hz.ChainId)
			}

			for _, eT := range epochTrackers {
				stakeibcKeeper.RemoveEpochTracker(ctx, eT.EpochIdentifier)
			}

			levstakeHostZones := v3.NewHostZones(hostZones)
			levstakeParams := v3.NewParams(params)
			levstakeEpochTrackers := v3.NewEpochTrackers(epochTrackers)

			defaultGenesis := levstakeibctypes.DefaultGenesis()

			genesis := levstakeibctypes.GenesisState{
				Params:           levstakeParams,
				PortId:           defaultGenesis.PortId,
				HostZoneList:     levstakeHostZones,
				EpochTrackerList: levstakeEpochTrackers,
				NextPositionId:   defaultGenesis.NextPositionId,
				PositionList:     defaultGenesis.PositionList,
			}

			levstakeibc.InitGenesis(ctx, app.LevstakeibcKeeper, genesis)

			ctx.Logger().Info("start to run module migrations...")

			return app.mm.RunMigrations(ctx, app.configurator, vm)
		},
	)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Errorf("Failed to read upgrade info from disk: %w", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	var storeUpgrades *storetypes.StoreUpgrades

	switch upgradeInfo.Name {
	case v3.UpgradeName:
		storeUpgrades = &storetypes.StoreUpgrades{
			Added: []string{levstakeibctypes.StoreKey, admintypes.StoreKey, lendingpooltypes.StoreKey, mockborrowtypes.ModuleName},
			//Renamed: []storetypes.StoreRename{{OldKey: stakeibctypes.StoreKey, NewKey: levstakeibctypes.StoreKey}},
			Deleted: []string{},
		}
	}

	if storeUpgrades != nil {
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, storeUpgrades))
	}
}
