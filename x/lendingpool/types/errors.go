package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrEmptyCreator           = errorsmod.Register(ModuleName, 200, "creator address is empty")
	ErrInvalidDenom           = errorsmod.Register(ModuleName, 201, "denom field is empty")
	ErrInvalidInterestModel   = errorsmod.Register(ModuleName, 202, "invalid interest model")
	ErrPoolExists             = errorsmod.Register(ModuleName, 203, "pool already exists for the selected denomination.")
	ErrPoolNotFound           = errorsmod.Register(ModuleName, 204, "pool does not exist")
	ErrLoanNotFound           = errorsmod.Register(ModuleName, 205, "loan does not exist")
	ErrInvalidDepositor       = errorsmod.Register(ModuleName, 206, "invalid depositor address")
	ErrInvalidPoolID          = errorsmod.Register(ModuleName, 207, "invalid pool ID")
	ErrInvalidLoanId          = errorsmod.Register(ModuleName, 208, "invalid loan id")
	ErrInvalidDepositCoins    = errorsmod.Register(ModuleName, 209, "invalid deposit coins")
	ErrInvalidWithdrawCoins   = errorsmod.Register(ModuleName, 210, "invalid withdraw coins")
	ErrInvalidProtocolTaxRate = errorsmod.Register(ModuleName, 211, "invalid protocol tax rate")
	ErrInvalidBorrowCoins     = errorsmod.Register(ModuleName, 212, "invalid borrow coins")
	ErrInvalidModelParams     = errorsmod.Register(ModuleName, 213, "invalid interest model params")
	ErrNotEnoughReserve       = errorsmod.Register(ModuleName, 214, "not enough pool coins")
	ErrNotEnoughCollateral    = errorsmod.Register(ModuleName, 215, "not enough collateral registered")
	ErrOverflowMaxDebtRatio   = errorsmod.Register(ModuleName, 216, "overflow the max debt ratio")
	ErrInvalidRedemptionRate  = errorsmod.Register(ModuleName, 217, "invalid redemption rate")
)
