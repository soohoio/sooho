package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/lendingpool/types"
)

// SetPool
func (k Keeper) SetPool(ctx sdk.Context, pool types.LendingPool) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&pool)
	store.Set(types.GetLendingPoolKey(pool.Id), bz)
}

// GetPool
func (k Keeper) GetPool(ctx sdk.Context, id uint64) (types.LendingPool, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLendingPoolKey(id))
	if bz == nil {
		return types.LendingPool{}, false
	}
	var pool types.LendingPool
	k.cdc.MustUnmarshal(bz, &pool)
	return pool, true
}

func (k Keeper) GetDenomPool(ctx sdk.Context, denom string) (types.LendingPool, bool) {
	var id uint64
	k.IterateAllPools(ctx, func(pool types.LendingPool) bool {
		if pool.Denom == denom {
			id = pool.Id
			return true
		}
		return false
	})

	// if id == 0 no pool will be found since we enforce it in CreatePool
	return k.GetPool(ctx, id)
}

// CreatePool creates a pool
func (k Keeper) CreatePool(ctx sdk.Context, msg types.MsgCreatePool) (types.MsgCreatePoolResponse, error) {
	if _, found := k.GetDenomPool(ctx, msg.Denom); found {
		return types.MsgCreatePoolResponse{}, types.ErrPoolExists
	}
	id := k.GetNextLendingPoolID(ctx)
	if id == 0 {
		// increment if id == 0 to disallow pool ID with 0
		id++
	}

	pool := types.LendingPool{
		Id:              id,
		Pool:            sdk.NewCoins(),
		RedemptionRate:  sdk.OneDec(),
		UtilizationRate: sdk.ZeroDec(),
		InterestRate:    msg.InterestRate,
		Denom:           msg.Denom,
	}

	id = id + 1
	k.SetNextLendingPoolID(ctx, id)
	k.SetPool(ctx, pool)
	return types.MsgCreatePoolResponse{}, nil
}

// TODO: DeletePool deletes a pool

func (k Keeper) GetNextLendingPoolID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextLendingPoolIDKey)
	return binary.LittleEndian.Uint64(bz)
}

func (k Keeper) SetNextLendingPoolID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, id)
	store.Set(types.GetNextLendingPoolKey(), bz)
}

func (k Keeper) IterateAllPools(ctx sdk.Context, cb func(pool types.LendingPool) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	poolStore := prefix.NewStore(store, types.LendingPoolKey)

	iterator := poolStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var p types.LendingPool
		k.cdc.MustUnmarshal(iterator.Value(), &p)
		if cb(p) {
			break
		}
	}
}
