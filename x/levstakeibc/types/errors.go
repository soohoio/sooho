package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrEpochNotFound            = sdkerrors.Register(ModuleName, 1501, "epoch not found")
	ErrFailedToRegisterHostZone = sdkerrors.Register(ModuleName, 1502, "failed to register host zone")
	ErrInvalidAmount            = sdkerrors.Register(ModuleName, 1503, "invalid amount")
	ErrRequiredFieldEmpty       = sdkerrors.Register(ModuleName, 1504, "required field is missing")
	ErrLeverageRatio            = sdkerrors.Register(ModuleName, 1505, "leverage ratio can not be less than 1.0")
)
