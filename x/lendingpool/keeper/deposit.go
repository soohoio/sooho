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

	amount := k.GetPoolExchangeRate(msg.PoolId)

	k.mintIBToken(ctx, pool.Denom, amount)

	// send the minted ib tokens to depositor
	DepositorAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return types.MsgDepositResponse{}, err
	}
}

// Withdraw withdraws from a pool
func (k Keeper) Withdraw(ctx sdk.Context, msg types.MsgWithdraw) (types.MsgWithdrawResponse, error) {

}

// mintIBToken mints corresponding interest-bearing token for deposits
func (k Keeper) mintIBToken(ctx sdk.Context, base_denom string, amount math.Int) error {
	ibDenom := types.IBPrefix + base_denom
	newCoins := sdk.NewCoins(sdk.NewCoin(ibDenom, amount))
	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}