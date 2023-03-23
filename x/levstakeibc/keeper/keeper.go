package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
	memKey   storetypes.StoreKey
	//paramstore            paramtypes.Subspace
	accountKeeper types.AccountKeeper
	bankKeeper    bankkeeper.Keeper
	scopedKeeper  capabilitykeeper.ScopedKeeper
	StakingKeeper stakingkeeper.Keeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	memKey storetypes.StoreKey,
	//ps paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	StakingKeeper stakingkeeper.Keeper,
) Keeper {
	//if !ps.HasKeyTable() {
	//	ps = ps.WithKeyTable(types.ParamKeyTable())
	//}

	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		//paramstore: ps,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		scopedKeeper:  scopedKeeper,
		StakingKeeper: StakingKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
