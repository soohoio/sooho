package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
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

func (k Keeper) Borrow(ctx sdk.Context, denom, clientModule string, borrower sdk.AccAddress, borrowAmount, collateral sdk.Coins) (uint64, error) {
	pool, found := k.GetDenomPool(ctx, denom)

	// validate borrow amount and denom
	if len(borrowAmount) > 1 || len(borrowAmount) == 0 {
		return 0, types.ErrInvalidBorrowCoins
	} else if !found {
		return 0, types.ErrPoolNotFound
	}
	loanId := k.GetNextLoanID(ctx)
	k.SetNextLoanID(ctx, loanId+1)

	// update denom pool
	if borrowAmount.IsAnyGT(pool.Coins) {
		return 0, types.ErrNotEnoughReserve
	}
	pool.Coins = pool.Coins.Sub(borrowAmount...)
	k.SetPool(ctx, pool)

	borrowedDec := sdk.NewDecFromInt(borrowAmount.AmountOf(pool.Denom))
	totalAssetDec := sdk.NewDecFromInt(collateral.AmountOf(pool.Denom)).Add(borrowedDec)
	newLoan := types.Loan{
		Id:       loanId,
		Denom:    denom,
		Borrower: borrower.String(),
		// dummy price for now
		// TODO: fix this
		InitMarkPrice:   sdk.NewDecCoinsFromCoins(sdk.NewCoin("dummy", sdk.OneInt())),
		TotalAssetValue: totalAssetDec,
		BorrowedValue:   borrowedDec,
	}

	k.SetLoan(ctx, newLoan)

	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, clientModule, borrowAmount)
	if err != nil {
		return 0, err
	}

	return loanId, nil
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
	if !found {
		return nil, types.ErrPoolNotFound
	}

	// Convert vars to Int
	totalAssetValueInt := loan.TotalAssetValue.TruncateInt()
	borrowedValueInt := loan.BorrowedValue.TruncateInt()
	repayAmountInt := amount.AmountOf(loan.Denom)
	repayInt := sdk.MinInt(repayAmountInt, borrowedValueInt)

	pool.Coins = pool.Coins.Add(sdk.NewCoin(pool.Denom, repayInt))
	k.SetPool(ctx, pool)

	// if borrowed == repay, delete and return change
	if borrowedValueInt.Equal(repayInt) {
		k.DeleteLoan(ctx, id)
		change := repayInt.Sub(borrowedValueInt)
		return sdk.NewCoins(sdk.NewCoin(loan.Denom, change)), nil
	}
	// else subtract repay amount from borrowed amount and save loan
	loan.BorrowedValue = sdk.NewDecFromInt(borrowedValueInt.Sub(repayInt))
	loan.TotalAssetValue = sdk.NewDecFromInt(totalAssetValueInt.Sub(repayInt))

	k.SetLoan(ctx, loan)

	return nil, nil
}

func (k Keeper) IterateAllLoans(ctx sdk.Context, cb func(loan types.Loan) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	loanStore := prefix.NewStore(store, types.LoanKey)

	iterator := loanStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var l types.Loan
		k.cdc.MustUnmarshal(iterator.Value(), &l)
		if cb(l) {
			break
		}
	}
}

// GetPools returns all the proposals from store
func (keeper Keeper) GetAllLoans(ctx sdk.Context) (loans []types.Loan) {
	keeper.IterateAllLoans(ctx, func(loan types.Loan) bool {
		loans = append(loans, loan)
		return false
	})
	return
}
