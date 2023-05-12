package types

import errorsmod "cosmossdk.io/errors"

// DONTCOVER

// x/icacallbacks module sentinel errors
var (
	ErrCallbackHandlerNotFound = errorsmod.Register(ModuleName, 400, "icacallback handler not found")
	ErrCallbackFailed          = errorsmod.Register(ModuleName, 401, "icacallback failed")
	ErrInvalidAcknowledgement  = errorsmod.Register(ModuleName, 402, "invalid acknowledgement")
)
