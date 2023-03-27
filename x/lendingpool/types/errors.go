package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrEmptyCreator           = sdkerrors.Register(ModuleName, 2, "creator address is empty")
	ErrEmptyDenom             = sdkerrors.Register(ModuleName, 3, "denom field is empty")
	ErrInvalidInterestModel   = sdkerrors.Register(ModuleName, 4, "invalid interest model")
	ErrPoolExists             = sdkerrors.Register(ModuleName, 5, "pool already exists for the selected denomination.")
	ErrPoolNotFound           = sdkerrors.Register(ModuleName, 6, "pool does not exist")
	ErrInvalidDepositor       = sdkerrors.Register(ModuleName, 7, "invalid depositor address")
	ErrInvalidPoolID          = sdkerrors.Register(ModuleName, 8, "invalid pool ID")
	ErrInvalidDepositCoins    = sdkerrors.Register(ModuleName, 9, "invalid deposit coins")
	ErrInvalidProtocolTaxRate = sdkerrors.Register(ModuleName, 10, "invalid protocol tax rate")
	ErrInvalidBorrowCoins     = sdkerrors.Register(ModuleName, 11, "invalid borrow coins")
	ErrInvalidModelParams     = sdkerrors.Register(ModuleName, 12, "invalid interest model params")
)
