package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/lendingpool/types"
)

func (k Keeper) Borrow(ctx sdk.Context, loan types.Loan) error {
	//pool, found := k.GetDenomPool(ctx, loan.Amount[0].Denom)
	//// validate amount
	//if len(loan.Amount) > 1 {
	//	return types.ErrInvalidBorrowCoins
	//} else if !found {
	//	return types.ErrPoolNotFound
	//}
	//
	//pool.

	return nil
}

func (k Keeper) Repay(ctx sdk.Context, amount sdk.Coins, borrower sdk.AccAddress) error {
	return nil
}

func (k Keeper) SetBorrow(ctx sdk.Context, loan types.Loan) {

}

func (k Keeper) GetBorrow(ctx sdk.Context, borrower sdk.AccAddress, denom string) types.Loan {
	return types.Loan{}
}
