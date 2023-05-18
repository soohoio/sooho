package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	admintypes "github.com/soohoio/stayking/v2/x/admin/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func (k msgServer) UpdateHostZone(goCtx context.Context, msg *types.MsgUpdateHostZone) (*types.MsgUpdateHostZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// admin address check
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if !k.AdminKeeper.IsAdmin(ctx, creator) {
		return nil, admintypes.ErrNotAdmin
	}

	chainId, err := k.GetChainID(ctx, msg.ConnectionId)
	if err != nil {
		errMsg := fmt.Sprintf("unable to obtain chain id from connection %s, err: %s", msg.ConnectionId, err.Error())
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(types.ErrFailedToUpdateHostZone, errMsg)
	}

	hostZone, found := k.GetHostZone(ctx, chainId)
	if !found {
		errMsg := fmt.Sprintf("Please register hostzone first for chain id :%v", chainId)
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(types.ErrFailedToUpdateHostZone, errMsg)
	}

	//update hostZone Values
	hostZone.ChainId = chainId
	hostZone.ConnectionId = msg.ConnectionId
	hostZone.Bech32Prefix = msg.Bech32Prefix
	hostZone.TransferChannelId = msg.TransferChannelId
	hostZone.HostDenom = msg.HostDenom
	hostZone.IbcDenom = msg.IbcDenom
	hostZone.UnbondingFrequency = msg.UnbondingFrequency

	k.SetHostZone(ctx, hostZone)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateZone,
			sdk.NewAttribute(types.AttributeKeyRecipientChain, chainId),
			sdk.NewAttribute(types.AttributeKeyConnectionId, msg.ConnectionId),
		),
	)

	return &types.MsgUpdateHostZoneResponse{}, nil
}
