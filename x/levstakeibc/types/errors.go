package types

import (
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrEpochNotFound                = errorsmod.Register(ModuleName, 100, "Epoch not found")                     // epoch(Blockchain Time) 값을 찾을 수 없을 때
	ErrFailureUpdateUnbondingRecord = errorsmod.Register(ModuleName, 101, "Failed to update a unbonding record") // HostZone 별로 Unbonding Record 의 상태 변경이 실패했을 때
	ErrFailedToRegisterHostZone     = errorsmod.Register(ModuleName, 102, "failed to register host zone")        // HostZone 등록 실패할 때
	ErrInvalidAmount                = errorsmod.Register(ModuleName, 103, "Invalid amount")                      // 입력된 Collateral(Equity), Debt 값이 기대되는 값하고 다른 값이 들어왔을 때
	ErrInsufficientFundsOnHostZone  = errorsmod.Register(ModuleName, 104, "Insufficient funds on HostZone")      // 언스테이킹시 HostZone 에 존재하는 Staked Balance 가 부족할 시
	ErrRequiredFieldEmpty           = errorsmod.Register(ModuleName, 105, "Required field is missing")           // 필수로 입력받아야 하는 값들이 빈 값일 때
	ErrLeverageRatio                = errorsmod.Register(ModuleName, 106, "Leverage ratio is not valid")         // 레버리지 배율이 유효한 값이 아닐 때
	ErrHostZoneNotFound             = errorsmod.Register(ModuleName, 107, "host zone not found")                 // 모듈 state 에 저장된 host zone 정보를 가져오지 못할 때

	ErrInvalidToken                    = sdkerrors.Register(ModuleName, 7, "invalid token")
	ErrNoValidatorWeights              = sdkerrors.Register(ModuleName, 8, "no non-zero validator weights")
	ErrUnmarshalFailure                = sdkerrors.Register(ModuleName, 9, "unable to unmarshal data structure")
	ErrValidatorDelegationChg          = sdkerrors.Register(ModuleName, 10, "can't change delegation on validator")
	ErrMaxNumValidators                = sdkerrors.Register(ModuleName, 12, "max number of validators reached")
	ErrValidatorAlreadyExists          = sdkerrors.Register(ModuleName, 13, "validator already exists")
	ErrMarkPriceDenomEmpty             = sdkerrors.Register(ModuleName, 14, "base denom for mark price is empty")
	ErrMarkPriceDenomExpired           = sdkerrors.Register(ModuleName, 15, "base denom data for mark price is expired")
	ErrLendingPoolBorrow               = sdkerrors.Register(ModuleName, 16, "received unexpected result after invoking the borrow func")
	ErrInsufficientFunds               = sdkerrors.Register(ModuleName, 17, "balance is insufficient")
	ErrNoValidatorAmts                 = sdkerrors.Register(ModuleName, 18, "could not fetch validator amts")
	ErrHostZoneICAAccountNotFound      = sdkerrors.Register(ModuleName, 19, "host zone's ICA account not found")
	ErrUndelegationAmount              = sdkerrors.Register(ModuleName, 20, "Undelegation amount is greater than stakedBal")
	ErrInvalidPacketCompletionTime     = sdkerrors.Register(ModuleName, 21, "invalid packet completion time")
	ErrInvalidUserRedemptionRecord     = sdkerrors.Register(ModuleName, 22, "user redemption record error")
	ErrRecordNotFound                  = sdkerrors.Register(ModuleName, 23, "record not found")
	ErrICAAccountNotFound              = sdkerrors.Register(ModuleName, 24, "ICA acccount not found on host zone")
	ErrICATxFailed                     = sdkerrors.Register(ModuleName, 25, "failed to submit ICA transaction")
	ErrLoanNotFound                    = sdkerrors.Register(ModuleName, 26, "loan not found")
	ErrPositionNotFound                = sdkerrors.Register(ModuleName, 27, "position not found")
	ErrReceiverNotFound                = sdkerrors.Register(ModuleName, 28, "receiver not found")
	ErrInvalidUnStakeWithLeverage      = sdkerrors.Register(ModuleName, 29, "invalid unstaked position with leverage")
	ErrInvalidAccount                  = sdkerrors.Register(ModuleName, 30, "invalid account")
	ErrAlreadyExistsPosition           = sdkerrors.Register(ModuleName, 31, "Already Exists the position")
	ErrInvalidInterchainAccountAddress = sdkerrors.Register(ModuleName, 32, "invalid interchain account address")
	ErrFailureAddCollateral            = sdkerrors.Register(ModuleName, 33, "failure add collateral to the position")
	ErrFailureAddDebt                  = sdkerrors.Register(ModuleName, 34, "failure add debt to the position")
	ErrPositionIsNotActive             = sdkerrors.Register(ModuleName, 35, "position is not active status")
	ErrMintAddedStAsset                = sdkerrors.Register(ModuleName, 36, "error added stake asset to mint st asset")
	ErrValidatorNotFound               = sdkerrors.Register(ModuleName, 37, "error validator not found")
	ErrDeleteValidatorFailed           = sdkerrors.Register(ModuleName, 38, "delete validator failed")
	ErrInvalidNumValidator             = sdkerrors.Register(ModuleName, 39, "invalid number of validator error")
	ErrWeightsNotDifferent             = sdkerrors.Register(ModuleName, 40, "validator weights haven't changed")
	ErrInvalidLeverageRatio            = sdkerrors.Register(ModuleName, 100, "invalid leverage ratio")
	ErrInvalidChainId                  = sdkerrors.Register(ModuleName, 101, "invalid chainId : there is no hostzone")
	ErrFeeAccountNotRegistered         = sdkerrors.Register(ModuleName, 102, "fee account is not registered")
	ErrChannelNotFound                 = sdkerrors.Register(ModuleName, 103, "error channel not found")
)
