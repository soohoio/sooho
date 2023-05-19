package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	admintypes "github.com/soohoio/stayking/v2/x/admin/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func (k msgServer) DeleteHostZone(goCtx context.Context, msg *types.MsgDeleteHostZone) (*types.MsgDeleteHostZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// admin address check
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if !k.AdminKeeper.IsAdmin(ctx, creator) {
		return nil, admintypes.ErrNotAdmin
	}

	hostZone, found := k.GetHostZone(ctx, msg.ChainId)
	if !found {
		errMsg := fmt.Sprintf("cannot find host zone for chain id :%v", msg.ChainId)
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(types.ErrFailedDeleteValidator, errMsg)
	}
	stSupply := k.bankKeeper.GetSupply(ctx, types.StAssetDenomFromHostZoneDenom(hostZone.HostDenom)).Amount

	err = k.BurnTokens(ctx, hostZone, stSupply)
	if err != nil {
		errMsg := fmt.Sprintf("cannot burn stTokens :%v%s for hostZone", stSupply, hostZone.HostDenom, msg.ChainId)
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(types.ErrFailedDeleteValidator, errMsg)
	}

	k.RemoveHostZone(ctx, hostZone)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateZone,
			sdk.NewAttribute(types.AttributeKeyRecipientChain, msg.ChainId),
		),
	)

	return &types.MsgDeleteHostZoneResponse{}, nil
}

// RemoveHostZone removes a hostzone from the store
func (k Keeper) RemoveHostZone(ctx sdk.Context, hostZone types.HostZone) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.HostZoneKey))
	store.Delete([]byte(hostZone.ChainId))
}
