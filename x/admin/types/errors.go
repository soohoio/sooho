package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrNoAdmins = sdkerrors.Register(ModuleName, 2, "no admins registered")
	ErrNotAdmin = sdkerrors.Register(ModuleName, 3, "not a privileged admin to perform transaction")
)
