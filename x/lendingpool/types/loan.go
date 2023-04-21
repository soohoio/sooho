package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewLoan(id uint64, denom, borrower string, active bool, totalValue, borrowedValue sdk.Dec, clientModule string) Loan {
	return Loan{
		Id:            id,
		Denom:         denom,
		Borrower:      borrower,
		ClientModule:  clientModule, // added
		Active:        active,
		TotalValue:    totalValue,
		BorrowedValue: borrowedValue,
	}
}

// GetDebtRatio assumes total asset value != 0
func (l Loan) GetDebtRatio() sdk.Dec {
	return l.BorrowedValue.Quo(l.TotalValue)
}