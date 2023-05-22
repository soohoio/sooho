package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v3/x/admin/types"
)

// Keeper of the distribution store
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec
}

// NewKeeper creates a new distribution Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
) Keeper {

	return Keeper{
		storeKey: key,
		cdc:      cdc,
	}
}

func (k Keeper) IsAdmin(ctx sdk.Context, addr sdk.AccAddress) bool {
	admins := k.GetAdmins(ctx)
	addrString := addr.String()
	for _, a := range admins.Admins {
		if a == addrString {
			return true
		}
	}
	return false
}

func (k Keeper) GetAdmins(ctx sdk.Context) types.Admins {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetAdminKey())
	var admins types.Admins
	k.cdc.MustUnmarshal(bz, &admins)
	return admins
}

func (k Keeper) SetAdmins(ctx sdk.Context, admins types.Admins) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&admins)
	store.Set(types.GetAdminKey(), bz)
}
