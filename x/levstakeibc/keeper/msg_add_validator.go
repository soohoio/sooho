package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	admintypes "github.com/soohoio/stayking/v2/x/admin/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"github.com/spf13/cast"
	"math"
)

// Appends a validator to host zone (if the host zone is not already at capacity)
// If the validator is added through governance, the weight is equal to the minimum weight across the set
// If the validator is added through an admin transactions, the weight is specified in the message
func (k msgServer) AddValidator(_ctx context.Context, msg *types.MsgAddValidator) (*types.MsgAddValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(_ctx)

	// admin address check
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if !k.AdminKeeper.IsAdmin(ctx, creator) {
		return nil, admintypes.ErrNotAdmin
	}

	// Get the corresponding host zone
	hostZone, found := k.GetHostZone(ctx, msg.HostZone)
	if !found {
		errMsg := fmt.Sprintf("Host Zone (%s) not found", msg.HostZone)
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrap(types.ErrHostZoneNotFound, errMsg)
	}

	// Get max number of validators and confirm we won't exceed it
	err = k.ConfirmValSetHasSpace(ctx, hostZone.Validators)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrMaxNumValidators, "cannot add validator on host zone")
	}

	// Check that we don't already have this validator
	// Grab the minimum weight in the process (to assign to validator's added through governance)
	var minWeight uint64 = math.MaxUint64
	for _, validator := range hostZone.Validators {
		if validator.Address == msg.Address {
			errMsg := fmt.Sprintf("Validator address (%s) already exists on Host Zone (%s)", msg.Address, msg.HostZone)
			k.Logger(ctx).Error(errMsg)
			return nil, errorsmod.Wrap(types.ErrValidatorAlreadyExists, errMsg)
		}
		if validator.Name == msg.Name {
			errMsg := fmt.Sprintf("Validator name (%s) already exists on Host Zone (%s)", msg.Name, msg.HostZone)
			k.Logger(ctx).Error(errMsg)
			return nil, errorsmod.Wrap(types.ErrValidatorAlreadyExists, errMsg)
		}
		// Store the min weight to assign to new validator added through governance (ignore zero-weight validators)
		if validator.Weight < minWeight && validator.Weight > 0 {
			minWeight = validator.Weight
		}
	}

	// If the validator was added via governance, set the weight to the min validator weight of the host zone
	valWeight := msg.Weight
	//if fromGovernance {
	//	valWeight = minWeight
	//}

	// Finally, add the validator to the host
	hostZone.Validators = append(hostZone.Validators, &types.Validator{
		Name:           msg.Name,
		Address:        msg.Address,
		Status:         types.Validator_ACTIVE,
		CommissionRate: msg.Commission,
		DelegationAmt:  sdk.ZeroInt(),
		Weight:         valWeight,
	})

	k.SetHostZone(ctx, hostZone)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddValidator,
			sdk.NewAttribute(types.AttributeKeyRecipientChain, msg.HostZone),
			sdk.NewAttribute(types.AttributeKeyAccountName, msg.Name),
			sdk.NewAttribute(types.AttributeKeyAddress, msg.Address),
		),
	)

	return &types.MsgAddValidatorResponse{}, nil
}

func (k Keeper) ConfirmValSetHasSpace(ctx sdk.Context, validators []*types.Validator) error {

	// get max val parameter
	maxNumVals, err := cast.ToIntE(k.GetParam(ctx, types.KeySafetyNumValidators))
	if err != nil {
		errMsg := fmt.Sprintf("Error getting safety max num validators | err: %s", err.Error())
		k.Logger(ctx).Error(errMsg)
		return errorsmod.Wrap(types.ErrMaxNumValidators, errMsg)
	}

	// count up the number of validators with non-zero weights
	numNonzeroWgtValidators := 0
	for _, validator := range validators {
		if validator.Weight > 0 {
			numNonzeroWgtValidators++
		}
	}

	// check if the number of validators with non-zero weights is below than the max
	if numNonzeroWgtValidators >= maxNumVals {
		errMsg := fmt.Sprintf("Attempting to add new validator but already reached max number of validators (%d)", maxNumVals)
		k.Logger(ctx).Error(errMsg)
		return errorsmod.Wrap(types.ErrMaxNumValidators, errMsg)
	}

	return nil
}
