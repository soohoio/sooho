package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrEpochNotFound            = sdkerrors.Register(ModuleName, 1501, "epoch not found")
	ErrFailedToRegisterHostZone = sdkerrors.Register(ModuleName, 1502, "failed to register host zone")
	ErrInvalidAmount            = sdkerrors.Register(ModuleName, 1503, "invalid amount")
	ErrRequiredFieldEmpty       = sdkerrors.Register(ModuleName, 1504, "required field is missing")
	ErrLeverageRatio            = sdkerrors.Register(ModuleName, 1505, "leverage ratio can not be less than 1.0")
	ErrInvalidHostZone          = sdkerrors.Register(ModuleName, 1506, "invalid host zone")
	ErrInvalidToken             = sdkerrors.Register(ModuleName, 1507, "invalid token")
	ErrNoValidatorWeights       = sdkerrors.Register(ModuleName, 1508, "no non-zero validator weights")
	ErrUnmarshalFailure         = sdkerrors.Register(ModuleName, 1509, "unable to unmarshal data structure")
	ErrValidatorDelegationChg   = sdkerrors.Register(ModuleName, 1510, "can't change delegation on validator")
	ErrHostZoneNotFound         = sdkerrors.Register(ModuleName, 1511, "host zone not found")
	ErrMaxNumValidators         = sdkerrors.Register(ModuleName, 1512, "max number of validators reached")
	ErrValidatorAlreadyExists   = sdkerrors.Register(ModuleName, 1513, "validator already exists")
	ErrMarkPriceDenomEmpty      = sdkerrors.Register(ModuleName, 1514, "base denom for mark price is empty")
	ErrMarkPriceDenomExpired    = sdkerrors.Register(ModuleName, 1515, "base denom data for mark price is expired")
	ErrLendingPoolBorrow        = sdkerrors.Register(ModuleName, 1516, "received unexpected result after invoking the borrow func")

	ErrInvalidLeverageRatio = sdkerrors.Register(ModuleName, 1600, "invalid leverage ratio")
)
