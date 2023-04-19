package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v2/x/records/types"
)

// SetPositionRecord set a specific positionRecord in the store
func (k Keeper) SetPositionRecord(ctx sdk.Context, positionRecord types.PositionRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionRecordKey))
	b := k.Cdc.MustMarshal(&positionRecord)
	store.Set([]byte(positionRecord.Id), b)
}

// GetUserPositionRecord returns a positionRecord from its id
func (k Keeper) GetPositionRecord(ctx sdk.Context, id string) (val types.PositionRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionRecordKey))
	b := store.Get([]byte(id))
	if b == nil {
		return val, false
	}
	k.Cdc.MustUnmarshal(b, &val)
	return val, true
}

// PositionRecord removes a positionRecord from the store
func (k Keeper) RemovePositionRecord(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionRecordKey))
	store.Delete([]byte(id))
}

// GetAllPositionRecord returns all userRedemptionRecord
func (k Keeper) GetAllPositionRecord(ctx sdk.Context) (list []types.PositionRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionRecordKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PositionRecord
		k.Cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// IteratePosition Records iterates zones
func (k Keeper) IteratePositionRecords(ctx sdk.Context,
	fn func(index int64, positionRecord types.PositionRecord) (stop bool),
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionRecordKey))

	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		positionRecord := types.PositionRecord{}
		k.Cdc.MustUnmarshal(iterator.Value(), &positionRecord)

		stop := fn(i, positionRecord)

		if stop {
			break
		}
		i++
	}
}
