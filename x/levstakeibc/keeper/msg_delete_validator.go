package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
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
		errMsg := fmt.Sprintf("host zone (%s) not found", msg.HostZone)
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(types.ErrHostZoneNotFound, errMsg)
	}
	for i, val := range hostZone.Validators {
		if val.GetAddress() == msg.ValAddr {
			if val.DelegationAmt.IsZero() && val.Weight == 0 {
				hostZone.Validators = append(hostZone.Validators[:i], hostZone.Validators[i+1:]...)
				k.SetHostZone(ctx, hostZone)
				return &types.MsgDeleteValidatorResponse{}, nil
			}
			errMsg := fmt.Sprintf("Validator (%s) has non-zero delegation (%v) or weight (%d)", msg.ValAddr, val.DelegationAmt, val.Weight)
			k.Logger(ctx).Error(errMsg)
			return nil, sdkerrors.Wrapf(types.ErrFailedDeleteValidator, errMsg)
		}
	}
	errMsg := fmt.Sprintf("Validator address (%s) not found on host zone (%s)", msg.ValAddr, msg.HostZone)
	k.Logger(ctx).Error(errMsg)
	return nil, sdkerrors.Wrapf(types.ErrFailedDeleteValidator, errMsg)

	return &types.MsgDeleteValidatorResponse{}, nil
}
