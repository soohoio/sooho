package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidFromAddr   = sdkerrors.Register(ModuleName, 2, "invalid from address")
	ErrEmptyDenom        = sdkerrors.Register(ModuleName, 3, "empty denom field")
	ErrEmptyCollateral   = sdkerrors.Register(ModuleName, 4, "empty collateral field")
	ErrEmptyBorrowAmount = sdkerrors.Register(ModuleName, 5, "empty borrow field")
	ErrInvalidLoanId     = sdkerrors.Register(ModuleName, 6, "invalid loan id")
	ErrInvalidAmount     = sdkerrors.Register(ModuleName, 7, "invalid repay amount field")
)
