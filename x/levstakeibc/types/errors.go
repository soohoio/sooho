package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrEpochNotFound               = sdkerrors.Register(ModuleName, 1, "epoch not found")
	ErrFailedToRegisterHostZone    = sdkerrors.Register(ModuleName, 2, "failed to register host zone")
	ErrInvalidAmount               = sdkerrors.Register(ModuleName, 3, "invalid amount")
	ErrRequiredFieldEmpty          = sdkerrors.Register(ModuleName, 4, "required field is missing")
	ErrLeverageRatio               = sdkerrors.Register(ModuleName, 5, "leverage ratio can not be less than 1.0")
	ErrInvalidHostZone             = sdkerrors.Register(ModuleName, 6, "invalid host zone")
	ErrInvalidToken                = sdkerrors.Register(ModuleName, 7, "invalid token")
	ErrNoValidatorWeights          = sdkerrors.Register(ModuleName, 8, "no non-zero validator weights")
	ErrUnmarshalFailure            = sdkerrors.Register(ModuleName, 9, "unable to unmarshal data structure")
	ErrValidatorDelegationChg      = sdkerrors.Register(ModuleName, 10, "can't change delegation on validator")
	ErrHostZoneNotFound            = sdkerrors.Register(ModuleName, 11, "host zone not found")
	ErrMaxNumValidators            = sdkerrors.Register(ModuleName, 12, "max number of validators reached")
	ErrValidatorAlreadyExists      = sdkerrors.Register(ModuleName, 13, "validator already exists")
	ErrMarkPriceDenomEmpty         = sdkerrors.Register(ModuleName, 14, "base denom for mark price is empty")
	ErrMarkPriceDenomExpired       = sdkerrors.Register(ModuleName, 15, "base denom data for mark price is expired")
	ErrLendingPoolBorrow           = sdkerrors.Register(ModuleName, 16, "received unexpected result after invoking the borrow func")
	ErrInsufficientFunds           = sdkerrors.Register(ModuleName, 17, "balance is insufficient")
	ErrNoValidatorAmts             = sdkerrors.Register(ModuleName, 18, "could not fetch validator amts")
	ErrHostZoneICAAccountNotFound  = sdkerrors.Register(ModuleName, 19, "host zone's ICA account not found")
	ErrUndelegationAmount          = sdkerrors.Register(ModuleName, 20, "Undelegation amount is greater than stakedBal")
	ErrInvalidPacketCompletionTime = sdkerrors.Register(ModuleName, 21, "invalid packet completion time")
	ErrInvalidUserRedemptionRecord = sdkerrors.Register(ModuleName, 22, "user redemption record error")
	ErrRecordNotFound              = sdkerrors.Register(ModuleName, 23, "record not found")
	ErrICAAccountNotFound          = sdkerrors.Register(ModuleName, 24, "ICA acccount not found on host zone")
	ErrICATxFailed                 = sdkerrors.Register(ModuleName, 25, "failed to submit ICA transaction")

	ErrInvalidLeverageRatio = sdkerrors.Register(ModuleName, 100, "invalid leverage ratio")
	ErrInvalidChainId       = sdkerrors.Register(ModuleName, 101, "invalid chainId : there is no hostzone")
)
