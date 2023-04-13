package types

import (
	"errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrAlreadyFulfilled     = errors.New("query already fulfilled")
	ErrSucceededNoDelete    = errors.New("query succeeded; do not not execute default behavior")
	ErrInvalidICQProof      = errors.New("icq query response failed")
	ErrICQCallbackNotFound  = errors.New("icq callback id not found")
	ErrFailedICQResponse    = errors.New("icq query response with failed response")
	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")
)
