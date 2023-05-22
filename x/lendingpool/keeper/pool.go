package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v3/x/lendingpool/types"
)

// SetPool
func (k Keeper) SetPool(ctx sdk.Context, pool types.Pool) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&pool)
	store.Set(types.GetLendingPoolKey(pool.Id), bz)
}

// GetPool
func (k Keeper) GetPool(ctx sdk.Context, id uint64) (types.Pool, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLendingPoolKey(id))
	if bz == nil {
		return types.Pool{}, false
	}
	var pool types.Pool
	k.cdc.MustUnmarshal(bz, &pool)
	return pool, true
}

func (k Keeper) GetDenomPool(ctx sdk.Context, denom string) (types.Pool, bool) {
	var id uint64
	k.IterateAllPools(ctx, func(pool types.Pool) bool {
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
	id := k.GetNextPoolID(ctx)
	if id == 0 {
		// increment if id == 0 to disallow pool ID with 0
		id++
	}

	pool := types.Pool{
		Id:                       id,
		RemainingCoins:           sdk.ZeroDec(),
		RedemptionRate:           sdk.OneDec(),
		RedemptionRateWithoutTax: sdk.OneDec(),
		TotalCoins:               sdk.ZeroDec(),
		MaxDebtRatio:             msg.MaxDebtRatio,
		InterestModel:            msg.InterestModel,
		Denom:                    msg.Denom,
	}

	id = id + 1
	k.SetNextPoolID(ctx, id)
	k.SetPool(ctx, pool)
	return types.MsgCreatePoolResponse{}, nil
}

// TODO: DeletePool deletes a pool

func (k Keeper) GetNextPoolID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextLendingPoolIDKey)
	return binary.LittleEndian.Uint64(bz)
}

func (k Keeper) SetNextPoolID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, id)
	store.Set(types.GetNextLendingPoolKey(), bz)
}

func (k Keeper) IterateAllPools(ctx sdk.Context, cb func(pool types.Pool) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	poolStore := prefix.NewStore(store, types.LendingPoolKey)

	iterator := poolStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var p types.Pool
		k.cdc.MustUnmarshal(iterator.Value(), &p)
		if cb(p) {
			break
		}
	}
}

// GetPools returns all the proposals from store
func (k Keeper) GetAllPools(ctx sdk.Context) (pools types.Pools) {
	k.IterateAllPools(ctx, func(pool types.Pool) bool {
		pools = append(pools, pool)
		return false
	})
	return
}
