package types

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"gopkg.in/yaml.v2"
)

var (
	DefaultDepositInterval        uint64 = 1
	DefaultDelegateInterval       uint64 = 1
	DefaultReinvestInterval       uint64 = 1
	DefaultRewardsInterval        uint64 = 1
	DefaultRedemptionRateInterval uint64 = 1

	DefaultStaykingCommission               uint64 = 10
	DefaultValidatorRebalancingThreshold    uint64 = 100 // divide by 10,000, so 100 = 1%
	DefaultICATimeoutNanos                  uint64 = 600000000000
	DefaultBufferSize                       uint64 = 5             // 1/5=20% of the epoch
	DefaultIbcTimeoutBlocks                 uint64 = 360           // 300 blocks (5secs / 1block ~= 25 minutes
	DefaultIBCTransferTimeoutNanos          uint64 = 1800000000000 // 30 minutes
	DefaultFeeTransferTimeoutNanos          uint64 = 1800000000000 // 30 minutes
	DefaultSafetyMinRedemptionRateThreshold uint64 = 90            // divide by 100, so 90 = 0.9
	DefaultSafetyMaxRedemptionRateThreshold uint64 = 150           // divide by 100, so 150 = 1.5
	DefaultMaxStakeICACallsPerEpoch         uint64 = 100
	DefaultSafetyNumValidators              uint64 = 10
	DefaultSafetyMarkPriceExpirationTime    uint64 = 300_000_000_000

	KeyDepositInterval                  = []byte("DepositInterval")
	KeyDelegateInterval                 = []byte("DelegateInterval")
	KeyReinvestInterval                 = []byte("ReinvestInterval")
	KeyRewardsInterval                  = []byte("RewardsInterval")
	KeyRedemptionRateInterval           = []byte("RedemptionRateInterval")
	KeyStaykingCommission               = []byte("StaykingCommission")
	KeyValidatorRebalancingThreshold    = []byte("ValidatorRebalancingThreshold")
	KeyICATimeoutNanos                  = []byte("ICATimeoutNanos")
	KeyFeeTransferTimeoutNanos          = []byte("FeeTransferTimeoutNanos")
	KeyBufferSize                       = []byte("BufferSize")
	KeyIbcTimeoutBlocks                 = []byte("IbcTimeoutBlocks")
	KeySafetyMinRedemptionRateThreshold = []byte("SafetyMinRedemptionRateThreshold")
	KeySafetyMaxRedemptionRateThreshold = []byte("SafetyMaxRedemptionRateThreshold")
	KeyMaxStakeICACallsPerEpoch         = []byte("MaxStakeICACallsPerEpoch")
	KeyIBCTransferTimeoutNanos          = []byte("IBCTransferTimeoutNanos")
	KeySafetyNumValidators              = []byte("SafetyNumValidators")
	KeySafetyMarkPriceExpirationTime    = []byte("SafetyMarkPriceExpirationTime")
)

// NewParams creates a new Params instance
func NewParams(
	deposit_interval uint64,
	delegate_interval uint64,
	rewards_interval uint64,
	redemption_rate_interval uint64,
	stayking_commission uint64,
	reinvest_interval uint64,
	validator_rebalancing_threshold uint64,
	ica_timeout_nanos uint64,
	buffer_size uint64,
	ibc_timeout_blocks uint64,
	fee_transfer_timeout_nanos uint64,
	max_stake_ica_calls_per_epoch uint64,
	safety_min_redemption_rate_threshold uint64,
	safety_max_redemption_rate_threshold uint64,
	ibc_transfer_timeout_nanos uint64,
	safety_num_validators uint64,
	safety_markprice_expiration_time uint64,
) Params {
	return Params{
		DepositInterval:                  deposit_interval,
		DelegateInterval:                 delegate_interval,
		RewardsInterval:                  rewards_interval,
		RedemptionRateInterval:           redemption_rate_interval,
		StaykingCommission:               stayking_commission,
		ReinvestInterval:                 reinvest_interval,
		ValidatorRebalancingThreshold:    validator_rebalancing_threshold,
		IcaTimeoutNanos:                  ica_timeout_nanos,
		BufferSize:                       buffer_size,
		IbcTimeoutBlocks:                 ibc_timeout_blocks,
		FeeTransferTimeoutNanos:          fee_transfer_timeout_nanos,
		MaxStakeIcaCallsPerEpoch:         max_stake_ica_calls_per_epoch,
		SafetyMinRedemptionRateThreshold: safety_min_redemption_rate_threshold,
		SafetyMaxRedemptionRateThreshold: safety_max_redemption_rate_threshold,
		IbcTransferTimeoutNanos:          ibc_transfer_timeout_nanos,
		SafetyNumValidators:              safety_num_validators,
		SafetyMarkPriceExpirationTime:    safety_markprice_expiration_time,
	}
}

func DefaultParams() Params {
	return NewParams(
		DefaultDepositInterval,
		DefaultDelegateInterval,
		DefaultRewardsInterval,
		DefaultRedemptionRateInterval,
		DefaultStaykingCommission,
		DefaultReinvestInterval,
		DefaultValidatorRebalancingThreshold,
		DefaultICATimeoutNanos,
		DefaultBufferSize,
		DefaultIbcTimeoutBlocks,
		DefaultFeeTransferTimeoutNanos,
		DefaultMaxStakeICACallsPerEpoch,
		DefaultSafetyMinRedemptionRateThreshold,
		DefaultSafetyMaxRedemptionRateThreshold,
		DefaultIBCTransferTimeoutNanos,
		DefaultSafetyNumValidators,
		DefaultSafetyMarkPriceExpirationTime,
	)
}

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDepositInterval, &p.DepositInterval, isPositive),
		paramtypes.NewParamSetPair(KeyDelegateInterval, &p.DelegateInterval, isPositive),
		paramtypes.NewParamSetPair(KeyRewardsInterval, &p.RewardsInterval, isPositive),
		paramtypes.NewParamSetPair(KeyRedemptionRateInterval, &p.RedemptionRateInterval, isPositive),
		paramtypes.NewParamSetPair(KeyStaykingCommission, &p.StaykingCommission, isCommission),
		paramtypes.NewParamSetPair(KeyReinvestInterval, &p.ReinvestInterval, isPositive),
		paramtypes.NewParamSetPair(KeyValidatorRebalancingThreshold, &p.ValidatorRebalancingThreshold, isThreshold),
		paramtypes.NewParamSetPair(KeyICATimeoutNanos, &p.IcaTimeoutNanos, isPositive),
		paramtypes.NewParamSetPair(KeyBufferSize, &p.BufferSize, isPositive),
		paramtypes.NewParamSetPair(KeyIbcTimeoutBlocks, &p.IbcTimeoutBlocks, isPositive),
		paramtypes.NewParamSetPair(KeyFeeTransferTimeoutNanos, &p.FeeTransferTimeoutNanos, validTimeoutNanos),
		paramtypes.NewParamSetPair(KeyMaxStakeICACallsPerEpoch, &p.MaxStakeIcaCallsPerEpoch, isPositive),
		paramtypes.NewParamSetPair(KeySafetyMinRedemptionRateThreshold, &p.SafetyMinRedemptionRateThreshold, validMinRedemptionRateThreshold),
		paramtypes.NewParamSetPair(KeySafetyMaxRedemptionRateThreshold, &p.SafetyMaxRedemptionRateThreshold, validMaxRedemptionRateThreshold),
		paramtypes.NewParamSetPair(KeyIBCTransferTimeoutNanos, &p.IbcTransferTimeoutNanos, validTimeoutNanos),
		paramtypes.NewParamSetPair(KeySafetyNumValidators, &p.SafetyNumValidators, isPositive),
		paramtypes.NewParamSetPair(KeySafetyMarkPriceExpirationTime, &p.SafetyMarkPriceExpirationTime, isPositive),
	}
}

func isThreshold(i interface{}) error {
	ival, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("parameter not accepted: %T", i)
	}

	if ival <= 0 {
		return fmt.Errorf("parameter must be positive: %d", ival)
	}
	if ival > 10000 {
		return fmt.Errorf("parameter must be less than 10,000: %d", ival)
	}
	return nil
}

func validMaxRedemptionRateThreshold(i interface{}) error {
	ival, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("parameter not accepted: %T", i)
	}

	maxVal := uint64(1000) // divide by 100, so 1000 => 10

	if ival > maxVal {
		return fmt.Errorf("parameter must be l.t. 1000: %d", ival)
	}

	return nil
}

func validMinRedemptionRateThreshold(i interface{}) error {
	ival, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("parameter not accepted: %T", i)
	}

	minVal := uint64(75) // divide by 100, so 75 => 0.75

	if ival < minVal {
		return fmt.Errorf("parameter must be g.t. 75: %d", ival)
	}

	return nil
}

func validTimeoutNanos(i interface{}) error {
	ival, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("parameter not accepted: %T", i)
	}

	tenMin := uint64(600000000000)
	oneHour := uint64(600000000000 * 6)

	if ival < tenMin {
		return fmt.Errorf("parameter must be g.t. 600000000000ns: %d", ival)
	}
	if ival > oneHour {
		return fmt.Errorf("parameter must be less than %dns: %d", oneHour, ival)
	}
	return nil
}

func isCommission(i interface{}) error {
	ival, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("commission not accepted: %T", i)
	}

	if ival > 100 {
		return fmt.Errorf("commission must be less than 100: %d", ival)
	}
	return nil
}

func isPositive(i interface{}) error {
	ival, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("parameter not accepted: %T", i)
	}

	if ival <= 0 {
		return fmt.Errorf("parameter must be positive: %d", ival)
	}
	return nil
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
