package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewLoan(id uint64, denom, borrower string, active bool, totalValue, borrowedValue sdk.Dec, clientModule string, leverageRatio sdk.Dec) Loan {
	return Loan{
		Id:                id,
		Denom:             denom,
		Borrower:          borrower,
		ClientModule:      clientModule, // added
		Active:            active,
		TotalValue:        totalValue,
		BorrowedValue:     borrowedValue,
		LeverageRatio:     leverageRatio,
		InitTotalValue:    totalValue,
		InitBorrowedValue: borrowedValue,
	}
}

// GetDebtRatio assumes total asset value != 0
func (l Loan) GetDebtRatio() sdk.Dec {
	if l.TotalValue.Equal(sdk.ZeroDec()) {
		return sdk.ZeroDec()
	}
	return l.BorrowedValue.Quo(l.TotalValue)
}
