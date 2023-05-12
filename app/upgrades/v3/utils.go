package v3

import (
	levstakeibctypes "github.com/soohoio/stayking/v2/x/levstakeibc/types"
	stakeibctypes "github.com/soohoio/stayking/v2/x/stakeibc/types"
)

func NewHostZones(zones []stakeibctypes.HostZone) []levstakeibctypes.HostZone {
	var res []levstakeibctypes.HostZone
	for _, z := range zones {
		migratedZone := levstakeibctypes.HostZone{
			ChainId:               z.ChainId,
			ConnectionId:          z.ConnectionId,
			Bech32Prefix:          z.Bech32Prefix,
			TransferChannelId:     z.TransferChannelId,
			WithdrawalAccount:     newICAAccount(z.WithdrawalAccount),
			FeeAccount:            newICAAccount(z.FeeAccount),
			DelegationAccount:     newICAAccount(z.DelegationAccount),
			HostDenom:             z.HostDenom,
			IbcDenom:              z.IbcDenom,
			LastRedemptionRate:    z.LastRedemptionRate,
			RedemptionRate:        z.RedemptionRate,
			StakedBal:             z.StakedBal,
			Address:               z.Address,
			UnbondingFrequency:    z.UnbondingFrequency,
			Validators:            newValidators(z.Validators),
			BlacklistedValidators: newValidators(z.BlacklistedValidators),
		}
		res = append(res, migratedZone)
	}
	return res
}

func newICAAccount(icaAccount *stakeibctypes.ICAAccount) *levstakeibctypes.ICAAccount {
	if icaAccount == nil {
		return nil
	}
	return &levstakeibctypes.ICAAccount{
		Address: icaAccount.Address,
		Target:  levstakeibctypes.ICAType(icaAccount.Target),
	}
}

func newValidators(validators []*stakeibctypes.Validator) []*levstakeibctypes.Validator {
	var res []*levstakeibctypes.Validator
	for _, v := range validators {
		var exchangeRate *levstakeibctypes.ValidatorExchangeRate
		if v.InternalExchangeRate == nil {
			exchangeRate = nil
		} else {
			exchangeRate = &levstakeibctypes.ValidatorExchangeRate{
				InternalTokensToSharesRate: v.InternalExchangeRate.InternalTokensToSharesRate,
				EpochNumber:                v.InternalExchangeRate.EpochNumber,
			}
		}
		newVal := &levstakeibctypes.Validator{
			Name:                 v.Name,
			Address:              v.Address,
			Status:               levstakeibctypes.Validator_ValidatorStatus(v.Status),
			CommissionRate:       v.CommissionRate,
			DelegationAmt:        v.DelegationAmt,
			Weight:               v.Weight,
			InternalExchangeRate: exchangeRate,
		}
		res = append(res, newVal)
	}
	return res
}

func NewParams(params stakeibctypes.Params) levstakeibctypes.Params {
	defaultGenesis := levstakeibctypes.DefaultGenesis()
	safetyPriceExpirationTime := defaultGenesis.GetParams().SafetyMarkPriceExpirationTime
	liquidationPerformanceFee := defaultGenesis.GetParams().LiquidationPerformanceFee
	return levstakeibctypes.Params{
		RewardsInterval:                  params.RewardsInterval,
		DepositInterval:                  params.DepositInterval,
		DelegateInterval:                 params.DelegateInterval,
		ReinvestInterval:                 params.ReinvestInterval,
		RedemptionRateInterval:           params.RedemptionRateInterval,
		StaykingCommission:               params.StaykingCommission,
		ZoneComAddress:                   params.ZoneComAddress,
		IbcTimeoutBlocks:                 params.IbcTimeoutBlocks,
		IbcTransferTimeoutNanos:          params.IbcTransferTimeoutNanos,
		FeeTransferTimeoutNanos:          params.FeeTransferTimeoutNanos,
		IcaTimeoutNanos:                  params.IcaTimeoutNanos,
		MaxStakeIcaCallsPerEpoch:         params.MaxStakeIcaCallsPerEpoch,
		ValidatorRebalancingThreshold:    params.ValidatorRebalancingThreshold,
		BufferSize:                       params.BufferSize,
		SafetyMinRedemptionRateThreshold: params.SafetyMinRedemptionRateThreshold,
		SafetyMaxRedemptionRateThreshold: params.SafetyMaxRedemptionRateThreshold,
		SafetyNumValidators:              params.SafetyNumValidators,
		SafetyMaxSlashPercent:            params.SafetyMaxSlashPercent,
		SafetyMarkPriceExpirationTime:    safetyPriceExpirationTime,
		LiquidationPerformanceFee:        liquidationPerformanceFee,
	}
}

func NewEpochTrackers(ets []stakeibctypes.EpochTracker) []levstakeibctypes.EpochTracker {
	var res []levstakeibctypes.EpochTracker
	for _, et := range ets {
		newET := levstakeibctypes.EpochTracker{
			EpochIdentifier:    et.EpochIdentifier,
			EpochNumber:        et.EpochNumber,
			NextEpochStartTime: et.NextEpochStartTime,
			Duration:           et.Duration,
		}
		res = append(res, newET)
	}
	return res
}
