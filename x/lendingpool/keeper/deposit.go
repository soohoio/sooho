package keeper

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/lendingpool/types"
	lendingpooltypes "github.com/soohoio/stayking/v2/x/lendingpool/types"
)

// Deposit makes a deposit into a pool
func (k Keeper) Deposit(ctx sdk.Context, msg types.MsgDeposit) (types.MsgDepositResponse, error) {
	pool, found := k.GetPool(ctx, msg.PoolId)
	if !found {
		return types.MsgDepositResponse{}, types.ErrPoolNotFound
	}

	exchangeRate := pool.RedemptionRate

	// assume inputting base tokens for now, e.g. msg.Amount = XXXuevmos
	ibAmount := sdk.NewDecFromInt(msg.Amount.AmountOf(pool.Denom)).Quo(exchangeRate).TruncateInt()

	// get the coins first
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return types.MsgDepositResponse{}, err
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, msg.Amount); err != nil {
		return types.MsgDepositResponse{}, err
	}

	amountDec := sdk.NewDecFromInt(msg.Amount.AmountOf(pool.Denom))
	pool.RemainingCoins = pool.RemainingCoins.Add(amountDec)
	pool.TotalCoins = pool.TotalCoins.Add(amountDec)
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

	ibDenom := getIBDenom(pool.Denom)

	// verify ib tokens amount
	if len(msg.Amount) != 1 || msg.Amount.AmountOf(ibDenom).Equal(sdk.ZeroInt()) {
		return types.MsgWithdrawResponse{}, types.ErrInvalidWithdrawCoins
	}

	// assume inputting ib tokens for now, e.g. msg.Amount = XXXibuevmos
	taxRate := pool.RedemptionRateWithoutTax.Sub(pool.RedemptionRate)
	if taxRate.LT(sdk.ZeroDec()) {
		return types.MsgWithdrawResponse{}, types.ErrInvalidRedemptionRate
	}
	baseAmount := sdk.NewDecFromInt(msg.Amount.AmountOf(ibDenom)).Mul(pool.RedemptionRate).TruncateInt()
	taxAmount := sdk.NewDecFromInt(msg.Amount.AmountOf(ibDenom)).Mul(taxRate).TruncateInt()
	// get the coins first
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return types.MsgWithdrawResponse{}, err
	}
	lendingpoolFeeAccount, err := sdk.AccAddressFromBech32(lendingpooltypes.LendingPoolFeeAccount)
	if err != nil {
		return types.MsgWithdrawResponse{}, err
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, msg.Amount); err != nil {
		return types.MsgWithdrawResponse{}, err
	}
	// burn the ib tokens from the coins
	if err := k.burnIBToken(ctx, pool.Denom, msg.Amount.AmountOf(ibDenom)); err != nil {
		return types.MsgWithdrawResponse{}, err
	}

	baseCoins := sdk.NewCoins(sdk.NewCoin(pool.Denom, baseAmount))
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, from, baseCoins); err != nil {
		return types.MsgWithdrawResponse{}, err
	}
	taxCoins := sdk.NewCoins(sdk.NewCoin(pool.Denom, taxAmount))
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, lendingpoolFeeAccount, taxCoins); err != nil {
		return types.MsgWithdrawResponse{}, err
	}
	baseCoinsWithTax := sdk.NewCoins(sdk.NewCoin(pool.Denom, baseAmount.Add(taxAmount)))

	// update pool
	amountDec := sdk.NewDecFromInt(baseCoinsWithTax.AmountOf(pool.Denom))
	pool.RemainingCoins = pool.RemainingCoins.Sub(amountDec)
	pool.TotalCoins = pool.TotalCoins.Sub(amountDec)

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
