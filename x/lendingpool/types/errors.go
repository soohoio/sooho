package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrEmptyCreator           = sdkerrors.Register(ModuleName, 2, "creator address is empty")
	ErrInvalidDenom           = sdkerrors.Register(ModuleName, 3, "denom field is empty")
	ErrInvalidInterestModel   = sdkerrors.Register(ModuleName, 4, "invalid interest model")
	ErrPoolExists             = sdkerrors.Register(ModuleName, 5, "pool already exists for the selected denomination.")
	ErrPoolNotFound           = sdkerrors.Register(ModuleName, 6, "pool does not exist")
	ErrLoanNotFound           = sdkerrors.Register(ModuleName, 7, "loan does not exist")
	ErrInvalidDepositor       = sdkerrors.Register(ModuleName, 8, "invalid depositor address")
	ErrInvalidPoolID          = sdkerrors.Register(ModuleName, 9, "invalid pool ID")
	ErrInvalidLoanId          = sdkerrors.Register(ModuleName, 10, "invalid loan id")
	ErrInvalidDepositCoins    = sdkerrors.Register(ModuleName, 11, "invalid deposit coins")
	ErrInvalidWithdrawCoins   = sdkerrors.Register(ModuleName, 12, "invalid withdraw coins")
	ErrInvalidProtocolTaxRate = sdkerrors.Register(ModuleName, 13, "invalid protocol tax rate")
	ErrInvalidBorrowCoins     = sdkerrors.Register(ModuleName, 14, "invalid borrow coins")
	ErrInvalidModelParams     = sdkerrors.Register(ModuleName, 15, "invalid interest model params")
	ErrNotEnoughReserve       = sdkerrors.Register(ModuleName, 16, "not enough pool coins")
	ErrNotEnoughCollateral    = sdkerrors.Register(ModuleName, 17, "not enough collateral registered")
	ErrOverflowMaxDebtRatio   = sdkerrors.Register(ModuleName, 18, "overflow the max debt ratio")
	ErrInvalidRedemptionRate  = sdkerrors.Register(ModuleName, 19, "invalid redemption rate")
)
