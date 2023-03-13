package keeper

import (
	"encoding/binary"
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

// CreatePool creates a pool
func (k Keeper) CreatePool(ctx sdk.Context, msg types.MsgCreatePool) (types.MsgCreatePoolResponse, error) {
	id := k.GetNextLendingPoolID(ctx)

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

// Deposit makes a deposit into a pool

// Withdraw withdraws from a pool

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
