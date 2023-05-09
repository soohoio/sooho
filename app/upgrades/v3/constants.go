package v3

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	upgrades "github.com/soohoio/stayking/v2/app/upgrades"
	admintypes "github.com/soohoio/stayking/v2/x/admin/types"
	lendingpooltypes "github.com/soohoio/stayking/v2/x/lendingpool/types"
	levstakeibctypes "github.com/soohoio/stayking/v2/x/levstakeibc/types"
	stakeibctypes "github.com/soohoio/stayking/v2/x/stakeibc/types"
)

const (
	UpgradeName = "tmp" // TODO: come up with an upgrade name
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{levstakeibctypes.StoreKey, admintypes.StoreKey, lendingpooltypes.StoreKey},
		Deleted: []string{stakeibctypes.StoreKey},
	},
}
