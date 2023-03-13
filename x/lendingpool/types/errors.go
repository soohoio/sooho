package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrEmptyCreator        = sdkerrors.Register(ModuleName, 2, "creator address is empty")
	ErrEmptyDenom          = sdkerrors.Register(ModuleName, 3, "denom field is empty")
	ErrInvalidInterestRate = sdkerrors.Register(ModuleName, 4, "invalid interest rate")
)
