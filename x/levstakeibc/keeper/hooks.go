package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	epochstypes "github.com/soohoio/stayking/v2/x/epochs/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"github.com/spf13/cast"
)

var _ epochstypes.EpochHooks = Hooks{}

// Hooks wrapper struct for incentives keeper
type Hooks struct {
	k Keeper
}

func (h Hooks) BeforeEpochStart(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {
	h.k.BeforeEpochStart(ctx, epochInfo)
}

func (h Hooks) AfterEpochEnd(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {
	h.k.AfterEpochEnd(ctx, epochInfo)
}

func (k Keeper) BeforeEpochStart(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {
	epochNumber, err := k.UpdateEpochTracker(ctx, epochInfo)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Unable to update epoch tracker, err: %s", err.Error()))
		return
	}

	if epochInfo.Identifier == epochstypes.DAY_EPOCH {
		// unbonding batch process
		k.InitiateAllHostZoneUnbondings(ctx, epochNumber)
		// Check previous epochs to see if unbondings finished, and sweep the tokens if so
		k.SweepAllUnbondedTokens(ctx)
		// clean up unused unbonding records
		k.CleanupEpochUnbondingRecords(ctx)
		// create an empty unbonding record for this epoch term
		k.CreateEpochUnbondingRecord(ctx, epochNumber)
	}

	if epochInfo.Identifier == epochstypes.STAYKING_EPOCH {

		// Delegation 의 Reward 받을 곳 지정
		k.SetWithdrawalAddress(ctx)

		k.CreateDepositRecordsForEpoch(ctx, epochNumber)
		depositRecords := k.RecordsKeeper.GetAllDepositRecord(ctx)

		// update redemption rate
		if epochNumber%k.GetParam(ctx, types.KeyRedemptionRateInterval) == 0 {
			hostZones := k.UpdateRedemptionRates(ctx, depositRecords)
			updatedForLoanTotalValue := k.UpdateLoanTotalValue(ctx, hostZones)
			if updatedForLoanTotalValue {
				k.Logger(ctx).Info("LOAN TOTAL VALUE UPDATED", fmt.Sprintf("epoch number : %v", epochNumber))
			}
		}

		// stayking > hostzone ( transfer ibc token )
		if epochNumber%k.GetParam(ctx, types.KeyDepositInterval) == 0 {
			k.TransferExistingDepositsToHostZones(ctx, epochNumber, depositRecords)
		}
		// hostzone > validator ( staking action )
		if epochNumber%k.GetParam(ctx, types.KeyDelegateInterval) == 0 {
			k.StakeExistingDepositsOnHostZones(ctx, epochNumber, depositRecords)
		}
		// reinvest
		if epochNumber%k.GetParam(ctx, types.KeyReinvestInterval) == 0 {
			k.ReinvestRewards(ctx)
		}
	}
}

func (k Keeper) AfterEpochEnd(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {}

func (k Keeper) Hooks() epochstypes.EpochHooks {
	return Hooks{k}
}

// Update the epoch information in the stakeibc epoch tracker
func (k Keeper) UpdateEpochTracker(ctx sdk.Context, epochInfo epochstypes.EpochInfo) (epochNumber uint64, err error) {
	epochNumber, err = cast.ToUint64E(epochInfo.CurrentEpoch)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Could not convert epoch number to uint64: %v", err))
		return 0, err
	}
	epochDurationNano, err := cast.ToUint64E(epochInfo.Duration.Nanoseconds())
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Could not convert epoch duration to uint64: %v", err))
		return 0, err
	}
	nextEpochStartTime, err := cast.ToUint64E(epochInfo.CurrentEpochStartTime.Add(epochInfo.Duration).UnixNano())
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Could not convert epoch duration to uint64: %v", err))
		return 0, err
	}
	epochTracker := types.EpochTracker{
		EpochIdentifier:    epochInfo.Identifier,
		EpochNumber:        epochNumber,
		Duration:           epochDurationNano,
		NextEpochStartTime: nextEpochStartTime,
	}
	k.SetEpochTracker(ctx, epochTracker)

	return epochNumber, nil
}
