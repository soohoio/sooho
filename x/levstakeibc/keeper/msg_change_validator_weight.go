package keeper

import (
	"context"
	"fmt"
	admintypes "github.com/soohoio/stayking/v2/x/admin/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func (k msgServer) ChangeValidatorWeight(goCtx context.Context, msg *types.MsgChangeValidatorWeight) (*types.MsgChangeValidatorWeightResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// admin address check
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if !k.AdminKeeper.IsAdmin(ctx, creator) {
		return nil, admintypes.ErrNotAdmin
	}

	hostZone, found := k.GetHostZone(ctx, msg.HostZone)
	if !found {
		k.Logger(ctx).Error(fmt.Sprintf("Host Zone %s not found", msg.HostZone))
		return nil, types.ErrInvalidHostZone
	}

	validators := hostZone.Validators
	for _, validator := range validators {
		if validator.GetAddress() == msg.ValAddr {

			// when changing a weight from 0 to non-zero, make sure we have space in the val set for this new validator
			if validator.Weight == 0 && msg.Weight > 0 {
				err := k.ConfirmValSetHasSpace(ctx, validators)
				if err != nil {
					return nil, sdkerrors.Wrap(types.ErrMaxNumValidators, "cannot set val weight from zero to nonzero on host zone")
				}
			}
			validator.Weight = msg.Weight
			k.SetHostZone(ctx, hostZone)
			return &types.MsgChangeValidatorWeightResponse{}, nil

		}
	}

	k.Logger(ctx).Error(fmt.Sprintf("Validator %s not found on Host Zone %s", msg.ValAddr, msg.HostZone))
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeChangeValidatorWeight,
			sdk.NewAttribute(types.AttributeKeyRecipientChain, hostZone.ChainId),
			sdk.NewAttribute(types.AttributeKeyAddress, msg.ValAddr),
			sdk.NewAttribute(types.AttributeKeyValidatorWeight, string(msg.Weight)),

		),
	)
	return nil, types.ErrValidatorNotFound
}
