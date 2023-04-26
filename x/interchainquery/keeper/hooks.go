package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"

	epochstypes "github.com/soohoio/stayking/v2/x/epochs/types"
	"github.com/soohoio/stayking/v2/x/interchainquery/types"
)

func (k Keeper) BeforeEpochStart(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {
	// Update the stakeibc epoch tracker
	epochNumber, err := k.UpdateEpochTracker(ctx, epochInfo)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Unable to update epoch tracker, err: %s", err.Error()))
		return
	}

	// PriceQuery Epoch
	if epochInfo.Identifier == epochstypes.PRICE_QUERY_EPOCH {
		// Initiate unbondings from any hostZone where it's appropriate
		k.EstimateSwapExactAmountOut(ctx, epochNumber)
	}

}

func (k Keeper) AfterEpochEnd(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {}

// Hooks wrapper struct for incentives keeper
type Hooks struct {
	k Keeper
}

var _ epochstypes.EpochHooks = Hooks{}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// epochs hooks
func (h Hooks) BeforeEpochStart(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {
	h.k.BeforeEpochStart(ctx, epochInfo)
}

func (h Hooks) AfterEpochEnd(ctx sdk.Context, epochInfo epochstypes.EpochInfo) {
	h.k.AfterEpochEnd(ctx, epochInfo)
}

// Update the epoch information in the epoch tracker
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
