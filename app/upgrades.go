package app

import (
	"fmt"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	v2 "github.com/soohoio/stayking/app/upgrades/v2"
	claimtypes "github.com/soohoio/stayking/x/claim/types"
	icacallbacktypes "github.com/soohoio/stayking/x/icacallbacks/types"
	recordtypes "github.com/soohoio/stayking/x/records/types"
	stakeibctypes "github.com/soohoio/stayking/x/stakeibc/types"
)

func (app *StayKingApp) setupUpgradeHandlers() {

	// v2 upgrade
	app.UpgradeKeeper.SetUpgradeHandler(
		v2.UpgradeName,
		v2.CreateUpgradeHandler(
			app.mm,
			app.configurator,
			app.appCodec,
			app.InterchainqueryKeeper,
			app.StakeibcKeeper,
			app.keys[claimtypes.StoreKey],
			app.keys[icacallbacktypes.StoreKey],
			app.keys[recordtypes.StoreKey],
			app.keys[stakeibctypes.StoreKey],
		),
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
	case "v2":
		storeUpgrades = &storetypes.StoreUpgrades{
			Deleted: []string{authz.ModuleName},
		}
	}

	if storeUpgrades != nil {
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, storeUpgrades))
	}
}
