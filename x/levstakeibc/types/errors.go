package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrEpochNotFound                = errorsmod.Register(ModuleName, 100, "Epoch not found")                                  // epoch(Blockchain Time) 값을 찾을 수 없을 때
	ErrFailureUpdateUnbondingRecord = errorsmod.Register(ModuleName, 101, "Failed to update a unbonding record")              // HostZone 별로 Unbonding Record 의 상태 변경이 실패했을 때
	ErrFailedToRegisterHostZone     = errorsmod.Register(ModuleName, 102, "failed to register host zone")                     // HostZone 등록 실패할 때
	ErrInvalidAmount                = errorsmod.Register(ModuleName, 103, "Invalid amount")                                   // 입력된 Collateral(Equity), Debt 값이 기대되는 값하고 다른 값이 들어왔을 때
	ErrInsufficientFundsOnHostZone  = errorsmod.Register(ModuleName, 104, "Insufficient funds on HostZone")                   // 언스테이킹시 HostZone 에 존재하는 Staked Balance 가 부족할 시
	ErrRequiredFieldEmpty           = errorsmod.Register(ModuleName, 105, "Required field is missing")                        // 필수로 입력받아야 하는 값들이 빈 값일 때
	ErrLeverageRatio                = errorsmod.Register(ModuleName, 106, "Leverage ratio is not valid")                      // 레버리지 배율이 유효한 값이 아닐 때
	ErrHostZoneNotFound             = errorsmod.Register(ModuleName, 107, "Host zone not found")                              // 모듈 state 에 저장된 host zone 정보를 가져오지 못할 때
	ErrInvalidToken                 = errorsmod.Register(ModuleName, 108, "Invalid token")                                    // Token 이름이 기대되는 값과 맞지 않을 시 (eg. ibc Denom )
	ErrNoValidatorWeights           = errorsmod.Register(ModuleName, 109, "No non-zero validator weights")                    // HostZone 의 Validators 들의 Weight 합이 0 이 될 때
	ErrNoValidatorAmts              = errorsmod.Register(ModuleName, 110, "Could not fetch validator amts")                   // Validator 의 Delegation/Undelegation Amount 양의 값에 대한 에러
	ErrUnmarshalFailure             = errorsmod.Register(ModuleName, 111, "Unable to unmarshal data structure")               // Callback Serialized Data 를 Unmarshal 할 수 없을 때
	ErrValidatorDelegationChg       = errorsmod.Register(ModuleName, 112, "Can't change delegation on validator")             // Validator Rebalance 실패 시
	ErrMaxNumValidators             = errorsmod.Register(ModuleName, 113, "Max number of validators reached")                 // HostZone 별 Weight > 0 인 Validator 최대 등록 수 가 넘을 시
	ErrValidatorAlreadyExists       = errorsmod.Register(ModuleName, 114, "Validator already exists")                         // HostZone 에 등록할 Validator 가 이미 존재할 경우
	ErrLendingPoolBorrow            = errorsmod.Register(ModuleName, 115, "Failure borrow debt amount when leverage staking") // Lev staking 시 lending pool 에서 debt 를 borrow 할 때 나는 에러
	ErrInsufficientFunds            = errorsmod.Register(ModuleName, 116, "Balance is insufficient")                          // Send, Burn 시 Balance 가 부족할 때
	ErrHostZoneICAAccountNotFound   = errorsmod.Register(ModuleName, 117, "Host zone's ICA account not found")                // HostZone 에서 ICA Account 를 찾지 못할 때
	ErrInvalidPacketCompletionTime  = errorsmod.Register(ModuleName, 118, "Invalid packet completion time")                   // Packet Completion Time 이 0인 경우
	ErrRecordNotFound               = errorsmod.Register(ModuleName, 119, "Record not found")                                 // x/records 모듈에서 deposit, unbonding, user redemption record 를 찾지 못할 때
	ErrICATxFailed                  = errorsmod.Register(ModuleName, 120, "Failed to submit ICA transaction")                 // Send ICA Tx 가 실패할 때
	ErrLoanNotFound                 = errorsmod.Register(ModuleName, 121, "Loan not found")                                   // x/lendingpool store 에서 Loan 을 못찾을 때
	ErrPositionNotFound             = errorsmod.Register(ModuleName, 122, "Position not found")                               // position 을 못찾는 경우
	ErrPoolNotFound                 = errorsmod.Register(ModuleName, 123, "Pool not found")                                   // x/lendingpool store 에서 Pool 을 못찾을 때
	ErrFailureOperatePosition       = errorsmod.Register(ModuleName, 124, "Failure operate position")                         // Position 생성, 조정, 종료, 청산 시 문제가 있을 때
	ErrInvalidAccount               = errorsmod.Register(ModuleName, 125, "Invalid account")                                  // Account 계정 파싱 시 문제가 있을 때
	ErrAlreadyExistsPosition        = errorsmod.Register(ModuleName, 126, "Already Exists the position")                      // 이미 특정 Denom 으로 레버리지한 Position 이 존재 할 때
	ErrFailureOperateLoan           = errorsmod.Register(ModuleName, 127, "Failure operate a loan in a pool")                 // Loan 에 Collateral, Debt 정보 업데이트 실패 시
	ErrPositionIsNotActive          = errorsmod.Register(ModuleName, 128, "Position is not active")                           // 포지션 상태가 Active 가 아닌 경우
	ErrFailureMintStAsset           = errorsmod.Register(ModuleName, 129, "Failure issue staked assets to mint")              // Staking 시 StToken 보증 토큰 발급 시 문제가 있을 때
	ErrValidatorNotFound            = errorsmod.Register(ModuleName, 130, "Validator not found")                              // validator 를 찾을 수 없을 때
	ErrFailedDeleteValidator        = errorsmod.Register(ModuleName, 140, "Failed to delete validator")                       // HostZone 에 등록된 Validator 를 삭제할 때
	ErrInvalidNumValidator          = errorsmod.Register(ModuleName, 141, "Invalid number of validator error")                // Rebalance 할 때 가용한 Validator 수 범위 안에 없을 때
	ErrWeightsNotDifferent          = errorsmod.Register(ModuleName, 142, "Validator weights haven't changed")
	ErrInvalidICAChannel            = errorsmod.Register(ModuleName, 143, "Invalid ICA Channel Connection, Channel, Port")

	ErrInvalidUserRedemptionRecord = errorsmod.Register(ModuleName, 199, "user redemption record error") // @Deprecated Claim 단계가 없어짐

)
