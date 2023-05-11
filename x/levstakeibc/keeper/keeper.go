package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/keeper"
	ibctransferkeeper "github.com/cosmos/ibc-go/v5/modules/apps/transfer/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v5/modules/core/keeper"
	ibctmtypes "github.com/cosmos/ibc-go/v5/modules/light-clients/07-tendermint/types"
	"github.com/soohoio/stayking/v2/utils"
	adminkeeper "github.com/soohoio/stayking/v2/x/admin/keeper"
	icacallbackskeeper "github.com/soohoio/stayking/v2/x/icacallbacks/keeper"
	icqkeeper "github.com/soohoio/stayking/v2/x/interchainquery/keeper"
	lendingpoolmodulekeeper "github.com/soohoio/stayking/v2/x/lendingpool/keeper"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	recordsmodulekeeper "github.com/soohoio/stayking/v2/x/records/keeper"
	"github.com/spf13/cast"
	"github.com/tendermint/tendermint/libs/log"
	"sort"
)

type Keeper struct {
	cdc                   codec.BinaryCodec
	storeKey              storetypes.StoreKey
	memKey                storetypes.StoreKey
	paramstore            paramtypes.Subspace
	accountKeeper         types.AccountKeeper
	bankKeeper            bankkeeper.Keeper
	TransferKeeper        ibctransferkeeper.Keeper
	scopedKeeper          capabilitykeeper.ScopedKeeper
	InterchainQueryKeeper icqkeeper.Keeper
	StakingKeeper         stakingkeeper.Keeper
	IBCKeeper             ibckeeper.Keeper
	ICAControllerKeeper   icacontrollerkeeper.Keeper
	ICACallbacksKeeper    icacallbackskeeper.Keeper
	RecordsKeeper         recordsmodulekeeper.Keeper
	LendingPoolKeeper     lendingpoolmodulekeeper.Keeper
	AdminKeeper           types.AdminKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	TransferKeeper ibctransferkeeper.Keeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	InterchainQueryKeeper icqkeeper.Keeper,
	StakingKeeper stakingkeeper.Keeper,
	ibcKeeper ibckeeper.Keeper,
	icaControllerKeeper icacontrollerkeeper.Keeper,
	icaCallbacksKeeper icacallbackskeeper.Keeper,
	recordsKeeper recordsmodulekeeper.Keeper,
	lendingKeeper lendingpoolmodulekeeper.Keeper,
	adminKeeper adminkeeper.Keeper,
) Keeper {
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:                   cdc,
		storeKey:              storeKey,
		memKey:                memKey,
		paramstore:            ps,
		accountKeeper:         accountKeeper,
		bankKeeper:            bankKeeper,
		TransferKeeper:        TransferKeeper,
		scopedKeeper:          scopedKeeper,
		InterchainQueryKeeper: InterchainQueryKeeper,
		StakingKeeper:         StakingKeeper,
		IBCKeeper:             ibcKeeper,
		ICAControllerKeeper:   icaControllerKeeper,
		ICACallbacksKeeper:    icaCallbacksKeeper,
		RecordsKeeper:         recordsKeeper,
		LendingPoolKeeper:     lendingKeeper,
		AdminKeeper:           adminKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) ClaimCapability(ctx sdk.Context, channelCap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, channelCap, name)
}

func (k Keeper) GetConnectionId(ctx sdk.Context, portId string) (string, error) {
	interchainAccounts := k.ICAControllerKeeper.GetAllInterchainAccounts(ctx)

	for _, interchainAccount := range interchainAccounts {
		if interchainAccount.PortId == portId {
			return interchainAccount.ConnectionId, nil
		}
	}

	return "", fmt.Errorf("PortId %s has no associated connectionId", portId)
}

func (k Keeper) GetChainID(ctx sdk.Context, connectionId string) (string, error) {
	conn, found := k.IBCKeeper.ConnectionKeeper.GetConnection(ctx, connectionId)
	if !found {
		errMsg := fmt.Sprintf("invalid connection id, %s not found", connectionId)
		k.Logger(ctx).Error(errMsg)
		return "", fmt.Errorf(errMsg)
	}
	clientState, found := k.IBCKeeper.ClientKeeper.GetClientState(ctx, conn.ClientId)
	if !found {
		errMsg := fmt.Sprintf("client id %s not found for connection %s", conn.ClientId, connectionId)
		k.Logger(ctx).Error(errMsg)
		return "", fmt.Errorf(errMsg)
	}
	client, ok := clientState.(*ibctmtypes.ClientState)
	if !ok {
		errMsg := fmt.Sprintf("invalid client state for client %s on connection %s", conn.ClientId, connectionId)
		k.Logger(ctx).Error(errMsg)
		return "", fmt.Errorf(errMsg)
	}

	return client.ChainId, nil
}

// GetHostZoneFromHostDenom returns a HostZone from a HostDenom
func (k Keeper) GetHostZoneFromHostDenom(ctx sdk.Context, denom string) (*types.HostZone, error) {
	var matchZone types.HostZone
	k.IterateHostZones(ctx, func(ctx sdk.Context, index int64, zoneInfo types.HostZone) error {
		if zoneInfo.HostDenom == denom {
			matchZone = zoneInfo
			return nil
		}
		return nil
	})
	if matchZone.ChainId != "" {
		return &matchZone, nil
	}
	return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "No HostZone for %s found", denom)
}

// IterateHostZones iterates zones
func (k Keeper) IterateHostZones(ctx sdk.Context, fn func(ctx sdk.Context, index int64, zoneInfo types.HostZone) error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.HostZoneKey))

	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		k.Logger(ctx).Debug(fmt.Sprintf("Iterating HostZone %d", i))
		zone := types.HostZone{}
		k.cdc.MustUnmarshal(iterator.Value(), &zone)

		error := fn(ctx, i, zone)

		if error != nil {
			break
		}
		i++
	}
}

// This function returns a map from Validator Address to how many extra tokens need to be given to that validator
//
//	positive implies extra tokens need to be given,
//	negative impleis tokens need to be taken away
func (k Keeper) GetValidatorDelegationAmtDifferences(ctx sdk.Context, hostZone types.HostZone) (map[string]sdk.Int, error) {
	validators := hostZone.GetValidators()
	delegationDelta := make(map[string]sdk.Int)
	totalDelegatedAmt := k.GetTotalValidatorDelegations(hostZone)
	targetDelegation, err := k.GetTargetValAmtsForHostZone(ctx, hostZone, totalDelegatedAmt)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error getting target val amts for host zone %s", hostZone.ChainId))
		return nil, err
	}
	for _, validator := range validators {
		targetDelForVal := targetDelegation[validator.GetAddress()]
		delegationDelta[validator.GetAddress()] = targetDelForVal.Sub(validator.DelegationAmt)
	}
	return delegationDelta, nil
}

func (k Keeper) GetTargetValAmtsForHostZone(ctx sdk.Context, hostZone types.HostZone, finalDelegation sdk.Int) (map[string]sdk.Int, error) {
	// Confirm the expected delegation amount is greater than 0
	if finalDelegation == sdk.ZeroInt() {
		k.Logger(ctx).Error(fmt.Sprintf("Cannot calculate target delegation if final amount is 0 %s", hostZone.ChainId))
		return nil, types.ErrNoValidatorWeights
	}

	// Sum the total weight across all validators
	totalWeight := k.GetTotalValidatorWeight(hostZone)
	if totalWeight == 0 {
		k.Logger(ctx).Error(fmt.Sprintf("No non-zero validators found for host zone %s", hostZone.ChainId))
		return nil, types.ErrNoValidatorWeights
	}
	k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Total Validator Weight: %d", totalWeight))

	// sort validators by weight ascending, this is inplace sorting!
	validators := hostZone.Validators

	for i, j := 0, len(validators)-1; i < j; i, j = i+1, j-1 {
		validators[i], validators[j] = validators[j], validators[i]
	}

	sort.SliceStable(validators, func(i, j int) bool { // Do not use `Slice` here, it is stochastic
		return validators[i].Weight < validators[j].Weight
	})

	// Assign each validator their portion of the delegation (and give any overflow to the last validator)
	targetUnbondingsByValidator := make(map[string]sdk.Int)
	totalAllocated := sdk.ZeroInt()
	for i, validator := range validators {
		// For the last element, we need to make sure that the totalAllocated is equal to the finalDelegation
		if i == len(validators)-1 {
			targetUnbondingsByValidator[validator.Address] = finalDelegation.Sub(totalAllocated)
		} else {
			// 여기는 validator 하나 이기 때문에 안탄다...
			delegateAmt := sdk.NewIntFromUint64(validator.Weight).Mul(finalDelegation).Quo(sdk.NewIntFromUint64(totalWeight))
			totalAllocated = totalAllocated.Add(delegateAmt)
			targetUnbondingsByValidator[validator.Address] = delegateAmt
		}
	}

	return targetUnbondingsByValidator, nil
}

func (k Keeper) GetTotalValidatorDelegations(hostZone types.HostZone) sdk.Int {
	validators := hostZone.GetValidators()
	total_delegation := sdk.ZeroInt()
	for _, validator := range validators {
		total_delegation = total_delegation.Add(validator.DelegationAmt)
	}
	return total_delegation
}

func (k Keeper) GetTotalValidatorWeight(hostZone types.HostZone) uint64 {
	validators := hostZone.GetValidators()
	total_weight := uint64(0)
	for _, validator := range validators {
		total_weight += validator.Weight
	}
	return total_weight
}

func (k Keeper) GetICATimeoutNanos(ctx sdk.Context, epochType string) (uint64, error) {
	epochTracker, found := k.GetEpochTracker(ctx, epochType)
	if !found {
		k.Logger(ctx).Error(fmt.Sprintf("Failed to get epoch tracker for %s", epochType))
		return 0, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Failed to get epoch tracker for %s", epochType)
	}
	// BUFFER by 5% of the epoch length
	bufferSizeParam := k.GetParam(ctx, types.KeyBufferSize)
	bufferSize := epochTracker.Duration / bufferSizeParam
	// buffer size should not be negative or longer than the epoch duration
	if bufferSize > epochTracker.Duration {
		k.Logger(ctx).Error(fmt.Sprintf("Invalid buffer size %d", bufferSize))
		return 0, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Invalid buffer size %d", bufferSize)
	}
	timeoutNanos := epochTracker.NextEpochStartTime - bufferSize
	timeoutNanosUint64, err := cast.ToUint64E(timeoutNanos)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Failed to convert timeoutNanos to uint64, error: %s", err.Error()))
		return 0, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Failed to convert timeoutNanos to uint64, error: %s", err.Error())
	}
	return timeoutNanosUint64, nil
}

func (k Keeper) AddDelegationToValidator(ctx sdk.Context, hostZone types.HostZone, validatorAddress string, amount sdk.Int, callbackId string) (success bool) {
	for _, validator := range hostZone.Validators {
		if validator.Address == validatorAddress {
			k.Logger(ctx).Info(utils.LogCallbackWithHostZone(hostZone.ChainId, callbackId,
				"  Validator %s, Current Delegation: %v, Delegation Change: %v", validator.Address, validator.DelegationAmt, amount))

			if amount.GTE(sdk.ZeroInt()) {
				validator.DelegationAmt = validator.DelegationAmt.Add(amount)
				return true
			}
			absAmt := amount.Abs()
			if absAmt.GT(validator.DelegationAmt) {
				k.Logger(ctx).Error(fmt.Sprintf("Delegation amount %v is greater than validator %s delegation amount %v", absAmt, validatorAddress, validator.DelegationAmt))
				return false
			}
			validator.DelegationAmt = validator.DelegationAmt.Sub(absAmt)
			return true
		}
	}

	k.Logger(ctx).Error(fmt.Sprintf("Could not find validator %s on host zone %s", validatorAddress, hostZone.ChainId))
	return false
}

func (k Keeper) GetLightClientTimeSafely(ctx sdk.Context, connectionID string) (uint64, error) {
	// get light client's latest height
	conn, found := k.IBCKeeper.ConnectionKeeper.GetConnection(ctx, connectionID)
	if !found {
		errMsg := fmt.Sprintf("invalid connection id, %s not found", connectionID)
		k.Logger(ctx).Error(errMsg)
		return 0, fmt.Errorf(errMsg)
	}
	// TODO(TEST-112) make sure to update host LCs here!
	latestConsensusClientState, found := k.IBCKeeper.ClientKeeper.GetLatestClientConsensusState(ctx, conn.ClientId)
	if !found {
		errMsg := fmt.Sprintf("client id %s not found for connection %s", conn.ClientId, connectionID)
		k.Logger(ctx).Error(errMsg)
		return 0, fmt.Errorf(errMsg)
	} else {
		latestTime := latestConsensusClientState.GetTimestamp()
		return latestTime, nil
	}
}
