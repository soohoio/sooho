package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrEmptyCreator           = errorsmod.Register(ModuleName, 200, "Creator address is empty")
	ErrInvalidDenom           = errorsmod.Register(ModuleName, 201, "Denom field is empty")
	ErrInvalidInterestModel   = errorsmod.Register(ModuleName, 202, "Invalid interest model")
	ErrPoolExists             = errorsmod.Register(ModuleName, 203, "Pool already exists for the selected denomination.")
	ErrPoolNotFound           = errorsmod.Register(ModuleName, 204, "Pool does not exist")
	ErrLoanNotFound           = errorsmod.Register(ModuleName, 205, "Loan does not exist")
	ErrInvalidDepositor       = errorsmod.Register(ModuleName, 206, "Invalid depositor address")
	ErrInvalidPoolID          = errorsmod.Register(ModuleName, 207, "Invalid pool ID")
	ErrInvalidLoanId          = errorsmod.Register(ModuleName, 208, "Invalid loan id")
	ErrInvalidDepositCoins    = errorsmod.Register(ModuleName, 209, "Invalid deposit coins")
	ErrInvalidWithdrawCoins   = errorsmod.Register(ModuleName, 210, "Invalid withdraw coins")
	ErrInvalidProtocolTaxRate = errorsmod.Register(ModuleName, 211, "Invalid protocol tax rate")
	ErrInvalidBorrowCoins     = errorsmod.Register(ModuleName, 212, "Invalid borrow coins")
	ErrInvalidModelParams     = errorsmod.Register(ModuleName, 213, "Invalid interest model params")
	ErrNotEnoughReserve       = errorsmod.Register(ModuleName, 214, "Not enough pool coins")
	ErrNotEnoughCollateral    = errorsmod.Register(ModuleName, 215, "Not enough collateral registered")
	ErrOverflowMaxDebtRatio   = errorsmod.Register(ModuleName, 216, "Overflow the max debt ratio")
	ErrInvalidRedemptionRate  = errorsmod.Register(ModuleName, 217, "Invalid redemption rate")
	ErrFailedLiquidate        = errorsmod.Register(ModuleName, 218, "Failed to liquidate position")
	ErrDivisionByZero         = errorsmod.Register(ModuleName, 219, "Failed to divide by zero")
)
