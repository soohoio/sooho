package keeper

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func (k Keeper) SetPosition(ctx sdk.Context, posistion types.Position) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&posistion)
	store.Set(types.GetPositionKey(posistion.Id), bz)
}

func (k Keeper) GetPosition(ctx sdk.Context, id uint64) (types.Position, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPositionKey(id))
	if bz == nil {
		return types.Position{}, false
	}
	var position types.Position
	k.cdc.MustUnmarshal(bz, &position)
	return position, true
}

func (k Keeper) GetNextPositionID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextPositionIDKey)
	return binary.LittleEndian.Uint64(bz)
}

func (k Keeper) SetNextPositionID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, id)
	store.Set(types.NextPositionIDKey, bz)
}
