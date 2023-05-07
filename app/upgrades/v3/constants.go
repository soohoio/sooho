package v3

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	upgrades "github.com/soohoio/stayking/v2/app/upgrades"

	icahosttypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/host/types"
)

const (
	UpgradeName = "tmp" // TODO: come up with an upgrade name
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{icahosttypes.StoreKey},
	},
}
