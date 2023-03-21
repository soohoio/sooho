package keeper

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v2/x/lendingpool/types"
)

// Deposit makes a deposit into a pool
func (k Keeper) Deposit(ctx sdk.Context, msg types.MsgDeposit) (types.MsgDepositResponse, error) {
	pool, found := k.GetPool(ctx, msg.PoolId)
	if !found {
		return types.MsgDepositResponse{}, types.ErrPoolNotFound
	}

	exchangeRate := pool.RedemptionRate

	// assume inputting base tokens for now, e.g. msg.Amount = XXXuevmos
	ibAmount := sdk.NewDecFromInt(msg.Amount.AmountOf(pool.Denom)).Mul(exchangeRate).TruncateInt()

	// get the coins first
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return types.MsgDepositResponse{}, err
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, msg.Amount); err != nil {
		return types.MsgDepositResponse{}, err
	}

	pool.Coins = pool.Coins.Add(msg.Amount...)

	// new utilization rate calculation
	// new totalAsset = pool.Coins / pool.UtilizationRate + msg.Amout
	// new utilization rate = totalBorrowed / totalAsset = (totalAsset - pool.Coins) / totalAsset
	// TODO: this calculation will be just a little bit off since it calculates total asset from
	// current coins and utilization rate (?)

	if pool.UtilizationRate.GT(sdk.ZeroDec()) {
		poolCoins := sdk.NewDecFromInt(pool.Coins.AmountOf(pool.Denom))
		totalAsset := poolCoins.
			Quo(pool.UtilizationRate).
			Add(sdk.NewDecFromInt(msg.Amount.AmountOf(pool.Denom)))
		pool.UtilizationRate = totalAsset.Sub(poolCoins).Quo(totalAsset)
	}
	k.SetPool(ctx, pool)

	// mint the ib tokens from the coins
	if err := k.mintIBToken(ctx, pool.Denom, ibAmount); err != nil {
		return types.MsgDepositResponse{}, err
	}

	ibCoins := sdk.NewCoins(sdk.NewCoin(getIBDenom(pool.Denom), ibAmount))
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, from, ibCoins); err != nil {
		return types.MsgDepositResponse{}, err
	}
	return types.MsgDepositResponse{Amount: ibCoins}, nil
}

// Withdraw withdraws from a pool
func (k Keeper) Withdraw(ctx sdk.Context, msg types.MsgWithdraw) (types.MsgWithdrawResponse, error) {
	pool, found := k.GetPool(ctx, msg.PoolId)
	if !found {
		return types.MsgWithdrawResponse{}, types.ErrPoolNotFound
	}

	exchangeRate := pool.RedemptionRate

	// assume inputting ib tokens for now, e.g. msg.Amount = XXXibuevmos
	baseAmount := sdk.NewDecFromInt(msg.Amount.AmountOf(pool.Denom)).Quo(exchangeRate).TruncateInt()

	// get the coins first
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return types.MsgWithdrawResponse{}, err
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, msg.Amount); err != nil {
		return types.MsgWithdrawResponse{}, err
	}
	// burn the ib tokens from the coins
	if err := k.burnIBToken(ctx, pool.Denom, msg.Amount.AmountOf(getIBDenom(pool.Denom))); err != nil {
		return types.MsgWithdrawResponse{}, err
	}

	baseCoins := sdk.NewCoins(sdk.NewCoin(pool.Denom, baseAmount))
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, from, baseCoins); err != nil {
		return types.MsgWithdrawResponse{}, err
	}

	// update pool
	pool.Coins = pool.Coins.Sub(msg.Amount...)

	// new utilization rate calculation
	// new totalAsset = pool.Coins / pool.UtilizationRate - baseCoins
	// new utilization rate = totalBorrowed / totalAsset = (totalAsset - pool.Coins) / totalAsset
	// TODO: this calculation will be just a little bit off since it calculates total asset from
	// current coins and utilization rate (?)

	if pool.UtilizationRate.GT(sdk.ZeroDec()) {
		poolCoins := sdk.NewDecFromInt(pool.Coins.AmountOf(pool.Denom))
		totalAsset := poolCoins.
			Quo(pool.UtilizationRate).
			Sub(sdk.NewDecFromInt(msg.Amount.AmountOf(pool.Denom)))
		pool.UtilizationRate = totalAsset.Sub(poolCoins).Quo(totalAsset)
	}
	k.SetPool(ctx, pool)

	return types.MsgWithdrawResponse{Amount: baseCoins}, nil
}

// mintIBToken mints corresponding interest-bearing token for deposits
func (k Keeper) mintIBToken(ctx sdk.Context, base_denom string, amount math.Int) error {
	ibDenom := getIBDenom(base_denom)
	newCoins := sdk.NewCoins(sdk.NewCoin(ibDenom, amount))
	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

// burnIBToken burns corresponding interest-bearing token for withdrawals
func (k Keeper) burnIBToken(ctx sdk.Context, base_denom string, amount math.Int) error {
	ibDenom := getIBDenom(base_denom)
	newCoins := sdk.NewCoins(sdk.NewCoin(ibDenom, amount))
	return k.bankKeeper.BurnCoins(ctx, types.ModuleName, newCoins)
}

func getIBDenom(base_denom string) string {
	return types.IBPrefix + base_denom
}