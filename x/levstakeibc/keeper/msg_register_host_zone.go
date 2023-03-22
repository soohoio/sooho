package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func (m msgServer) RegisterHostZone(goCtx context.Context, msg *types.MsgRegisterHostZone) (*types.MsgRegisterHostZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fmt.Sprintf("========================= input args =======================")
	fmt.Sprintf("connection-id : %s ", msg.ConnectionId)
	fmt.Sprintf("transferchannelid : %s ", msg.TransferChannelId)
	fmt.Sprintf("hostdenom : %s ", msg.HostDenom)
	fmt.Sprintf("ibcdenom : %s ", msg.IbcDenom)
	fmt.Sprintf("bech32prefix : %s ", msg.Bech32Prefix)
	fmt.Sprintf("unbondingfrequency : %d ", msg.UnbondingFrequency)
	fmt.Sprintf("creator : %s ", msg.Creator)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgRegisterHostZoneResponse{}, nil
}
