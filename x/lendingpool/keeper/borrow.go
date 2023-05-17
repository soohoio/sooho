package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"encoding/binary"
	"fmt"
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
	if len(borrowAmount) != 1 {
		return 0, types.ErrInvalidBorrowCoins
	} else if !found {
		return 0, types.ErrPoolNotFound
	}
	loanId := k.GetNextLoanID(ctx)
	k.SetNextLoanID(ctx, loanId+1)

	// update denom pool
	borrowAmountDec := sdk.NewDecFromInt(borrowAmount.AmountOf(denom))
	collateralAmountDec := sdk.NewDecFromInt(collateral.AmountOf(pool.Denom))
	if borrowAmountDec.GT(pool.RemainingCoins) {
		return 0, types.ErrNotEnoughReserve
	}
	pool.RemainingCoins = pool.RemainingCoins.Sub(borrowAmountDec)
	k.SetPool(ctx, pool)

	totalAssetDec := collateralAmountDec.Add(borrowAmountDec)
	if borrowAmountDec.Quo(totalAssetDec).GT(pool.MaxDebtRatio) {
		return 0, types.ErrNotEnoughCollateral
	}
	newLoan := types.NewLoan(loanId, denom, borrower.String(), true, totalAssetDec, borrowAmountDec, clientModule, totalAssetDec.Quo(collateralAmountDec))

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
	k.Logger(ctx).Info(fmt.Sprintf("[Repay Called] amount: %v ", amount))
	loan, found := k.GetLoan(ctx, id)
	if !found {
		return nil, types.ErrLoanNotFound
	}

	// update denom pool
	pool, found := k.GetDenomPool(ctx, loan.Denom)
	if !found {
		return nil, types.ErrPoolNotFound
	}

	// Convert vars to Int. chops off decimals for payments
	totalAssetValueInt := loan.TotalValue.TruncateInt()
	borrowedValueInt := loan.BorrowedValue.TruncateInt()

	repayAmountInt := amount.AmountOf(loan.Denom)
	repayInt := sdk.MinInt(repayAmountInt, borrowedValueInt)

	pool.RemainingCoins = pool.RemainingCoins.Add(sdk.NewDecFromInt(repayInt))
	k.SetPool(ctx, pool)

	// if borrowed == repay, delete and return change
	k.Logger(ctx).Info(fmt.Sprintf("Repay Called repayInt: %v repayAmountInt: %v borrowedVaultInt: %v ", repayInt, repayAmountInt, borrowedValueInt))
	if borrowedValueInt.Equal(repayInt) {
		// reduce total and remaining coins for the loss by chopping off decimals
		borrowedRem := getSubInt(loan.BorrowedValue)
		pool.TotalCoins = pool.TotalCoins.Sub(borrowedRem)
		k.SetPool(ctx, pool)

		k.DeleteLoan(ctx, id)
		changeInt := repayAmountInt.Sub(borrowedValueInt)

		change := sdk.NewCoins(sdk.NewCoin(loan.Denom, changeInt))
		borrowerAddr, err := sdk.AccAddressFromBech32(loan.Borrower)
		k.Logger(ctx).Info(fmt.Sprintf("Repay Called change: %v loan.Borrower: %v borrowedRem: %v borrowerAddr:%v ", change, loan.Borrower, borrowedRem, borrowerAddr))
		if err != nil {
			return nil, err
		}
		if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrowerAddr, change); err != nil {
			return nil, err
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.AttributeTypeRepayWithChange,
				sdk.NewAttribute(types.AttributeTypePoolId, string(pool.Id)),
				sdk.NewAttribute(types.AttributeTypeDenom, loan.Denom),
				sdk.NewAttribute(types.AttributeTypeBorrowerAddress, loan.Borrower),
				sdk.NewAttribute(types.AttributeTypeBorrowedValue, loan.BorrowedValue.String()),
				sdk.NewAttribute(types.AttributeTypeRepayValue, repayInt.String()),
				sdk.NewAttribute(types.AttributeTypeChangeValue, changeInt.String()),
			),
		)
		return change, nil
	}
	// else subtract repay amount from borrowed amount and save loan
	loan.BorrowedValue = sdk.NewDecFromInt(borrowedValueInt.Sub(repayInt))
	loan.TotalValue = sdk.NewDecFromInt(totalAssetValueInt.Sub(repayInt))
	k.SetLoan(ctx, loan)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.AttributeTypeRepayWithoutChange,
			sdk.NewAttribute(types.AttributeTypePoolId, string(pool.Id)),
			sdk.NewAttribute(types.AttributeTypeDenom, loan.Denom),
			sdk.NewAttribute(types.AttributeTypeBorrowerAddress, loan.Borrower),
			sdk.NewAttribute(types.AttributeTypeBorrowedValue, loan.BorrowedValue.String()),
			sdk.NewAttribute(types.AttributeTypeRepayValue, repayInt.String()),
		),
	)
	return nil, nil
}

func (k Keeper) AddCollateral(ctx sdk.Context, id uint64, amount sdk.Dec) error {
	l, found := k.GetLoan(ctx, id)
	if !found {
		return types.ErrLoanNotFound
	}
	l.TotalValue = l.TotalValue.Add(amount)

	l.InitTotalValue = l.InitTotalValue.Add(amount)
	l.LeverageRatio = l.InitTotalValue.Quo(l.InitTotalValue.Sub(l.InitBorrowedValue))

	k.SetLoan(ctx, l)
	return nil
}

func (k Keeper) AddDebt(ctx sdk.Context, id uint64, ibcDenom string, amount sdk.Dec) error {

	p, found := k.GetDenomPool(ctx, ibcDenom)
	if !found {
		return types.ErrPoolNotFound
	}

	l, found := k.GetLoan(ctx, id)
	if !found {
		return types.ErrLoanNotFound
	}

	l.TotalValue = l.TotalValue.Add(amount)
	l.BorrowedValue = l.BorrowedValue.Add(amount)

	l.InitTotalValue = l.InitTotalValue.Add(amount)
	l.InitBorrowedValue = l.InitBorrowedValue.Add(amount)
	l.LeverageRatio = l.InitTotalValue.Quo(l.InitTotalValue.Sub(l.InitBorrowedValue))

	if l.BorrowedValue.Quo(l.TotalValue).GTE(p.MaxDebtRatio) {
		return types.ErrOverflowMaxDebtRatio
	}

	k.SetLoan(ctx, l)
	coins := sdk.NewCoins(sdk.NewCoin(l.Denom, amount.TruncateInt()))

	if amount.GT(p.RemainingCoins) {
		return types.ErrNotEnoughReserve
	}
	p.RemainingCoins = p.RemainingCoins.Sub(amount)
	k.SetPool(ctx, p)

	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, l.ClientModule, coins)
}

func (k Keeper) Liquidate(ctx sdk.Context, id uint64) {
	l, _ := k.GetLoan(ctx, id)
	clientModule := *k.clientModules[l.ClientModule]
	err := clientModule.Liquidate(ctx, id)
	if err != nil {
		errorsmod.Wrapf(types.ErrFailedLiquidate, " liquidation failed with loan id %v ", id)
	}

}

func (k Keeper) UpdateTotalValue(ctx sdk.Context, loanId uint64, value sdk.Dec) {
	l, _ := k.GetLoan(ctx, loanId)
	l.TotalValue = value
	k.SetLoan(ctx, l)
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

// GetAllLoans returns all the loans from store
func (k Keeper) GetAllLoans(ctx sdk.Context) (loans []types.Loan) {
	k.IterateAllLoans(ctx, func(loan types.Loan) bool {
		loans = append(loans, loan)
		return false
	})
	return
}

// GetPoolLoans returns all the loans for a denom from store
func (k Keeper) GetPoolLoans(ctx sdk.Context, denom string) (loans []types.Loan) {
	k.IterateAllLoans(ctx, func(loan types.Loan) bool {
		if loan.Denom == denom {
			loans = append(loans, loan)
		}
		return false
	})
	return
}
