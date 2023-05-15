package keeper

import (
	"context"
	"fmt"
	admintypes "github.com/soohoio/stayking/v2/x/admin/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func (k msgServer) DeleteValidator(goCtx context.Context, msg *types.MsgDeleteValidator) (*types.MsgDeleteValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// admin address check
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if !k.AdminKeeper.IsAdmin(ctx, creator) {
		return nil, admintypes.ErrNotAdmin
	}

	// Removes a validator from a host zone
	// The validator must be zero-weight and have no delegations in order to be removed

	hostZone, found := k.GetHostZone(ctx, msg.HostZone)
	if !found {
		errMsg := fmt.Sprintf("HostZone (%s) not found", msg.HostZone)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrapf(types.ErrHostZoneNotFound, errMsg)
	}
	for i, val := range hostZone.Validators {
		if val.GetAddress() == msg.ValAddr {
			if val.DelegationAmt.IsZero() && val.Weight == 0 {
				hostZone.Validators = append(hostZone.Validators[:i], hostZone.Validators[i+1:]...)
				k.SetHostZone(ctx, hostZone)
				ctx.EventManager().EmitEvent(
					sdk.NewEvent(
						sdk.EventTypeMessage,
						sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					),
				)
				ctx.EventManager().EmitEvent(
					sdk.NewEvent(
						types.EventTypeDeleteValidator,
						sdk.NewAttribute(types.AttributeKeyRecipientChain, msg.HostZone),
						sdk.NewAttribute(types.AttributeKeyAccountName, val.Name),
						sdk.NewAttribute(types.AttributeKeyAddress, msg.ValAddr),
					),
				)

				return &types.MsgDeleteValidatorResponse{}, nil
			}
			errMsg := fmt.Sprintf("Validator (%s) has non-zero delegation (%v) or weight (%d)", msg.ValAddr, val.DelegationAmt, val.Weight)
			k.Logger(ctx).Error(errMsg)
			return nil, sdkerrors.Wrapf(types.ErrDeleteValidatorFailed, errMsg)
		}
	}
	errMsg := fmt.Sprintf("Validator address (%s) not found on host zone (%s)", msg.ValAddr, msg.HostZone)
	k.Logger(ctx).Error(errMsg)
	return nil, sdkerrors.Wrapf(types.ErrDeleteValidatorFailed, errMsg)

}
