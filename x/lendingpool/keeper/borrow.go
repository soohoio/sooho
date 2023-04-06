package keeper

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/lendingpool/types"
)

func (k Keeper) GetLoan(ctx sdk.Context, id uint64) (types.Loan, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLoanKey(id))
	if bz == nil {
		return types.Loan{}, false
	}
	var loan types.Loan
	k.cdc.MustUnmarshal(bz, &loan)
	return loan, true
}

func (k Keeper) SetLoan(ctx sdk.Context, loan types.Loan) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&loan)
	store.Set(types.GetLoanKey(loan.Id), bz)
}

func (k Keeper) DeleteLoan(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetLoanKey(id))
}

func (k Keeper) GetNextLoanID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetNextLoanKey())
	return binary.LittleEndian.Uint64(bz)
}

func (k Keeper) SetNextLoanID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, id)
	store.Set(types.GetNextLoanKey(), bz)
}

func (k Keeper) Borrow(ctx sdk.Context, denom string, borrowAmount, collateral sdk.Coins, borrower sdk.AccAddress) error {
	pool, found := k.GetDenomPool(ctx, denom)

	// validate borrow amount and denom
	if len(borrowAmount) > 1 {
		return types.ErrInvalidBorrowCoins
	} else if !found {
		return types.ErrPoolNotFound
	}
	loanId := k.GetNextLoanID(ctx)

	borrowedDec := sdk.NewDecFromInt(borrowAmount.AmountOf(pool.Denom))
	totalAssetDec := sdk.NewDecFromInt(collateral.AmountOf(pool.Denom)).Add(borrowedDec)
	newLoan := types.Loan{
		Id:       loanId,
		Denom:    denom,
		Borrower: borrower.String(),
		// dummy price for now
		InitMarkPrice:   sdk.NewDecCoinsFromCoins(sdk.NewCoins(sdk.NewCoin("dummy", sdk.OneInt()))...),
		TotalAssetValue: totalAssetDec,
		BorrowedValue:   borrowedDec,
	}

	k.SetLoan(ctx, newLoan)

	return nil
}

// Repay processes incoming repay.
// Assumes tokens are transferred before calling this function.
func (k Keeper) Repay(ctx sdk.Context, id uint64, amount sdk.Coins) (sdk.Coins, error) {
	loan, found := k.GetLoan(ctx, id)
	if !found {
		return nil, types.ErrLoanNotFound
	}

	// update denom pool
	pool, found := k.GetDenomPool(ctx, loan.Denom)
	pool.Coins = pool.Coins.Add(amount...)
	k.SetPool(ctx, pool)

	// truncate borrowed amount to Int
	borrowedValueInt := loan.BorrowedValue.TruncateInt()
	repayInt := amount.AmountOf(loan.Denom)

	// if borrowed <= repay, delete and return change
	if borrowedValueInt.LTE(repayInt) {
		k.DeleteLoan(ctx, id)
		change := repayInt.Sub(borrowedValueInt)
		return sdk.NewCoins(sdk.NewCoin(loan.Denom, change)), nil
	}

	// else subtract repay amount from borrowed amount and save loan

	return nil, nil
}
