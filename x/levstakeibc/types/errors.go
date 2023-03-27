package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrEpochNotFound            = sdkerrors.Register(ModuleName, 1505, "epoch not found")
	ErrFailedToRegisterHostZone = sdkerrors.Register(ModuleName, 1529, "failed to register host zone")
)
