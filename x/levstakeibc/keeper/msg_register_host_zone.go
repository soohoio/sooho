package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func (m msgServer) RegisterHostZone(goCtx context.Context, msg *types.MsgRegisterHostZone) (*types.MsgRegisterHostZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	m.Keeper.Logger(ctx).Info("========================= input args =======================")
	m.Keeper.Logger(ctx).Info("connection-id : %s ", msg.ConnectionId)
	m.Keeper.Logger(ctx).Info("transferchannelid : %s ", msg.TransferChannelId)
	m.Keeper.Logger(ctx).Info("hostdenom : %s ", msg.HostDenom)
	m.Keeper.Logger(ctx).Info("ibcdenom : %s ", msg.IbcDenom)
	m.Keeper.Logger(ctx).Info("bech32prefix : %s ", msg.Bech32Prefix)
	m.Keeper.Logger(ctx).Info("unbondingfrequency : %d ", msg.UnbondingFrequency)
	m.Keeper.Logger(ctx).Info("creator : %s ", msg.Creator)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgRegisterHostZoneResponse{}, nil
}
