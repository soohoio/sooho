package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/records module sentinel errors
var (
	ErrInvalidVersion               = errorsmod.Register(ModuleName, 300, "invalid version")
	ErrRedemptionAlreadyExists      = errorsmod.Register(ModuleName, 301, "redemption record already exists")
	ErrEpochUnbondingRecordNotFound = errorsmod.Register(ModuleName, 302, "epoch unbonding record not found") // unbonding record 를 epoch 로 찾을 수 없을 시
	ErrUnknownDepositRecord         = errorsmod.Register(ModuleName, 303, "unknown deposit record")
	ErrUnmarshalFailure             = errorsmod.Register(ModuleName, 304, "cannot unmarshal")
	ErrAddingHostZone               = errorsmod.Register(ModuleName, 305, "could not add hzu to epoch unbonding record")
)
