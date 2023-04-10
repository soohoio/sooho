package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	lendingpooltypes "github.com/soohoio/stayking/v2/x/lendingpool/types"
	"github.com/soohoio/stayking/v2/x/mockborrow/types"
)

// Keeper of the distribution store
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec
	ak       types.AccountKeeper
	bk       types.BankKeeper
	lk       types.LendingPoolKeeper
}

// NewKeeper creates a new distribution Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey,
	ak types.AccountKeeper, bk types.BankKeeper, lk types.LendingPoolKeeper) Keeper {

	// ensure mockborrow module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	return Keeper{
		storeKey: key,
		cdc:      cdc,
		ak:       ak,
		bk:       bk,
		lk:       lk,
	}
}

func (k Keeper) Borrow(ctx sdk.Context, denom string, borrower sdk.AccAddress,
	borrowAmount, collateral sdk.Coins) (uint64, error) {
	err := k.bk.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, collateral)
	if err != nil {
		return 0, err
	}
	return k.lk.Borrow(ctx, denom, types.ModuleName, borrower, borrowAmount, collateral)
}

// Repay just pays back the borrowed coins directly from mockborrow module.
// Since mockborrow is already holding the coins, we don't have to transfer from the account to module
// Anybody can trigger anybody's repay
func (k Keeper) Repay(ctx sdk.Context, id uint64, amount sdk.Coins) (sdk.Coins, error) {
	err := k.bk.SendCoinsFromModuleToModule(ctx, types.ModuleName, lendingpooltypes.ModuleName, amount)
	if err != nil {
		return nil, err
	}
	return k.lk.Repay(ctx, id, amount)
}

// GetTotalAssetValue return the total asset value given a loan id.
// assumes the loan exists.
func (k Keeper) GetTotalAssetValue(ctx sdk.Context, id uint64) sdk.Dec {
	loan, _ := k.lk.GetLoan(ctx, id)
	return loan.TotalAssetValue
}

// Fully liquidates a position on request from lending pool
func (k Keeper) Liquidate(ctx sdk.Context, id uint64) (sdk.Coins, error) {
	loan, _ := k.lk.GetLoan(ctx, id)
	repayAmountInt := k.GetTotalAssetValue(ctx, id).TruncateInt()
	repayAmount := sdk.NewCoins(sdk.NewCoin(loan.Denom, repayAmountInt))
	return k.Repay(ctx, id, repayAmount)
}
