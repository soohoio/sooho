package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrNoAdmins = errorsmod.Register(ModuleName, 2000, "no admins registered")
	ErrNotAdmin = errorsmod.Register(ModuleName, 2001, "not a privileged admin to perform transaction")
)
