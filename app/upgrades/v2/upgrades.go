package v2

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	claimtypes "github.com/soohoio/stayking/v3/x/claim/types"
	icacallbacktypes "github.com/soohoio/stayking/v3/x/icacallbacks/types"
	levstakeibcmodulekeeper "github.com/soohoio/stayking/v3/x/levstakeibc/keeper"
	levstakeibctypes "github.com/soohoio/stayking/v3/x/levstakeibc/types"
	recordtypes "github.com/soohoio/stayking/v3/x/records/types"
	stakeibctypes "github.com/soohoio/stayking/v3/x/stakeibc/types"

	"fmt"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	claimmigration "github.com/soohoio/stayking/v3/x/claim/migrations/v2"
	icacallbacksmigration "github.com/soohoio/stayking/v3/x/icacallbacks/migrations/v2"
	recordsmigration "github.com/soohoio/stayking/v3/x/records/migrations/v2"
	stakeibcmigration "github.com/soohoio/stayking/v3/x/stakeibc/migrations/v2"

	interchainquerykeeper "github.com/soohoio/stayking/v3/x/interchainquery/keeper"
)

const (
	UpgradeName = "v2.0"
)

// Helper function to log the migrated modules consensus versions
func logModuleMigration(ctx sdk.Context, versionMap module.VersionMap, moduleName string) {
	currentVersion := versionMap[moduleName]
	ctx.Logger().Info(fmt.Sprintf("migrating module %s from version %d to version %d", moduleName, currentVersion-1, currentVersion))
}

// CreateUpgradeHandler creates an SDK upgrade handler for v2
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	cdc codec.Codec,
	interchainqueryKeeper interchainquerykeeper.Keeper,
	levstakeibcKeeper levstakeibcmodulekeeper.Keeper,
	claimStoreKey storetypes.StoreKey,
	icacallbackStorekey storetypes.StoreKey,
	recordStoreKey storetypes.StoreKey,
	stakeibcStoreKey storetypes.StoreKey,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		currentVersions := mm.GetVersionMap()
		delete(vm, authz.ModuleName)

		levstakeibcKeeper.SetParams(ctx, levstakeibctypes.DefaultParams())

		// Store Migration

		logModuleMigration(ctx, currentVersions, claimtypes.ModuleName)
		if err := claimmigration.MigrateStore(ctx, claimStoreKey, cdc); err != nil {
			return vm, sdkerrors.Wrapf(err, "unable to migrate claim store")
		}

		logModuleMigration(ctx, currentVersions, icacallbacktypes.ModuleName)
		if err := icacallbacksmigration.MigrateStore(ctx, icacallbackStorekey, cdc); err != nil {
			return vm, sdkerrors.Wrapf(err, "unable to migrate icacallbacks store")
		}

		logModuleMigration(ctx, currentVersions, recordtypes.ModuleName)
		if err := recordsmigration.MigrateStore(ctx, recordStoreKey, cdc); err != nil {
			return vm, sdkerrors.Wrapf(err, "unable to migrate records store")
		}

		logModuleMigration(ctx, currentVersions, stakeibctypes.ModuleName)
		if err := stakeibcmigration.MigrateStore(ctx, stakeibcStoreKey, cdc); err != nil {
			return vm, sdkerrors.Wrapf(err, "unable to migrate stakeibc store")
		}

		vm[claimtypes.ModuleName] = currentVersions[claimtypes.ModuleName]
		vm[icacallbacktypes.ModuleName] = currentVersions[icacallbacktypes.ModuleName]
		vm[recordtypes.ModuleName] = currentVersions[recordtypes.ModuleName]
		vm[stakeibctypes.ModuleName] = currentVersions[stakeibctypes.ModuleName]

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
