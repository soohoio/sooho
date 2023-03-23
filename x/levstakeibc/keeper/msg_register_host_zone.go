package keeper

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func (m msgServer) RegisterHostZone(goCtx context.Context, msg *types.MsgRegisterHostZone) (*types.MsgRegisterHostZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	m.Keeper.Logger(ctx).Info("========================= input args =======================")
	m.Keeper.Logger(ctx).Info(fmt.Sprintf("connection-id : %s \n", msg.ConnectionId))
	m.Keeper.Logger(ctx).Info(fmt.Sprintf("transfer-channel-id : %s \n", msg.TransferChannelId))
	m.Keeper.Logger(ctx).Info(fmt.Sprintf("hostdenom : %s \n", msg.HostDenom))
	m.Keeper.Logger(ctx).Info(fmt.Sprintf("ibcdenom : %s \n", msg.IbcDenom))
	m.Keeper.Logger(ctx).Info(fmt.Sprintf("bech32prefix : %s \n", msg.Bech32Prefix))
	m.Keeper.Logger(ctx).Info(fmt.Sprintf("unbondingfrequency : %d \n", msg.UnbondingFrequency))
	m.Keeper.Logger(ctx).Info(fmt.Sprintf("creator : %s \n", msg.Creator))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	// TODO : chain id 별로 zone 어드레스 다르게 가져오기
	zoneAddress := types.NewZoneAddress("localstayking")
	acc := m.accountKeeper.NewAccount(
		ctx,
		authtypes.NewModuleAccount(
			authtypes.NewBaseAccountWithAddress(zoneAddress),
			zoneAddress.String(),
		),
	)
	m.accountKeeper.SetAccount(ctx, acc)

	hostZone := types.HostZone{
		//ChainId:            "", // ibc 구성 후 넣어야 함
		ConnectionId:       msg.ConnectionId,
		Bech32Prefix:       msg.Bech32Prefix,
		TransferChannelId:  msg.TransferChannelId,
		HostDenom:          msg.HostDenom,
		IbcDenom:           msg.IbcDenom,
		LastRedemptionRate: sdk.NewDec(1),
		RedemptionRate:     sdk.NewDec(1),
		StakedBal:          sdk.NewInt(0),
		Address:            "",
		UnbondingFrequency: msg.UnbondingFrequency,
	}

	return &types.MsgRegisterHostZoneResponse{}, nil
}

func (k Keeper) SetHostZone(ctx sdk.Context, hostZone types.HostZone) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HostZoneKey))
}
