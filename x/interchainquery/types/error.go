package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrInvalidICQProof      = errorsmod.Register(ModuleName, 500, "icq query response failed")
	ErrICQCallbackNotFound  = errorsmod.Register(ModuleName, 501, "icq callback id not found")
	ErrFailedICQResponse    = errorsmod.Register(ModuleName, 502, "icq query response with failed response")
	ErrFailedICQPacketQuery = errorsmod.Register(ModuleName, 503, "failed icq packet query")
	ErrInvalidVersion       = errorsmod.Register(ModuleName, 504, "invalid version")
)
