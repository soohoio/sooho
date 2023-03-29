package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func (k msgServer) RegisterHostZone(goCtx context.Context, msg *types.MsgRegisterHostZone) (*types.MsgRegisterHostZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 기존에 만들어진 HostChain 간의 ConnectionId 를 통해 Chain-ID 를 불러 온다.
	connectionEnd, found := k.IBCKeeper.ConnectionKeeper.GetConnection(ctx, msg.ConnectionId)
	if !found {
		errMsg := fmt.Sprintf("invalid connection id, %s not found", msg.ConnectionId)
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(types.ErrFailedToRegisterHostZone, errMsg)
	}

	chainId, err := k.GetChainID(ctx, msg.ConnectionId)
	if err != nil {
		errMsg := fmt.Sprintf("unable to obtain chain id from connection %s, err: %s", msg.ConnectionId, err.Error())
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(types.ErrFailedToRegisterHostZone, errMsg)
	}

	_, found = k.GetHostZone(ctx, chainId)
	if found {
		errMsg := fmt.Sprintf("invalid chain id, zone for %s already registered", chainId)
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(types.ErrFailedToRegisterHostZone, errMsg)
	}

	hostZones := k.GetAllHostZone(ctx)
	for _, hostZone := range hostZones {
		if hostZone.HostDenom == msg.HostDenom {
			errMsg := fmt.Sprintf("host denom %s already registered", msg.HostDenom)
			k.Logger(ctx).Error(errMsg)
			return nil, errorsmod.Wrapf(types.ErrFailedToRegisterHostZone, errMsg)
		}
		if hostZone.ConnectionId == msg.ConnectionId {
			errMsg := fmt.Sprintf("connectionId %s already registered", msg.ConnectionId)
			k.Logger(ctx).Error(errMsg)
			return nil, errorsmod.Wrapf(types.ErrFailedToRegisterHostZone, errMsg)
		}
		if hostZone.Bech32Prefix == msg.Bech32Prefix {
			errMsg := fmt.Sprintf("bech32prefix %s already registered", msg.Bech32Prefix)
			k.Logger(ctx).Error(errMsg)
			return nil, errorsmod.Wrapf(types.ErrFailedToRegisterHostZone, errMsg)
		}
	}

	zoneAddress := types.NewZoneAddress(chainId)
	acc := k.accountKeeper.NewAccount(
		ctx,
		authtypes.NewModuleAccount(
			authtypes.NewBaseAccountWithAddress(zoneAddress),
			zoneAddress.String(),
		),
	)
	k.accountKeeper.SetAccount(ctx, acc)

	hostZone := types.HostZone{
		ChainId:            chainId,
		ConnectionId:       msg.ConnectionId,
		Bech32Prefix:       msg.Bech32Prefix,
		TransferChannelId:  msg.TransferChannelId,
		HostDenom:          msg.HostDenom,
		IbcDenom:           msg.IbcDenom,
		LastRedemptionRate: sdk.NewDec(1),
		RedemptionRate:     sdk.NewDec(1),
		StakedBal:          sdk.NewInt(0),
		Address:            zoneAddress.String(),
		UnbondingFrequency: msg.UnbondingFrequency,
	}
	k.SetHostZone(ctx, hostZone)

	appVersion := string(icatypes.ModuleCdc.MustMarshalJSON(&icatypes.Metadata{
		Version:                icatypes.Version,
		ControllerConnectionId: hostZone.ConnectionId,
		HostConnectionId:       connectionEnd.Counterparty.ConnectionId,
		Encoding:               icatypes.EncodingProtobuf,
		TxType:                 icatypes.TxTypeSDKMultiMsg,
	}))

	delegateAccount := types.FormatICAAccountOwner(chainId, types.ICAType_DELEGATION)
	if err := k.ICAControllerKeeper.RegisterInterchainAccount(ctx, hostZone.ConnectionId, delegateAccount, appVersion); err != nil {
		errMsg := fmt.Sprintf("unable to register delegation account, err: %s", err.Error())
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(types.ErrFailedToRegisterHostZone, errMsg)
	}

	withdrawalAccount := types.FormatICAAccountOwner(chainId, types.ICAType_WITHDRAWAL)
	if err := k.ICAControllerKeeper.RegisterInterchainAccount(ctx, hostZone.ConnectionId, withdrawalAccount, appVersion); err != nil {
		errMsg := fmt.Sprintf("unable to register withdrawal account, err: %s", err.Error())
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(types.ErrFailedToRegisterHostZone, errMsg)
	}

	redemptionAccount := types.FormatICAAccountOwner(chainId, types.ICAType_REDEMPTION)
	if err := k.ICAControllerKeeper.RegisterInterchainAccount(ctx, hostZone.ConnectionId, redemptionAccount, appVersion); err != nil {
		errMsg := fmt.Sprintf("unable to register redemption account, err: %s", err.Error())
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(types.ErrFailedToRegisterHostZone, errMsg)
	}

	feeAccount := types.FormatICAAccountOwner(chainId, types.ICAType_FEE)
	if err := k.ICAControllerKeeper.RegisterInterchainAccount(ctx, hostZone.ConnectionId, feeAccount, appVersion); err != nil {
		errMsg := fmt.Sprintf("unable to register fee account, err: %s", err.Error())
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(types.ErrFailedToRegisterHostZone, errMsg)
	}

	icqAccount := types.FormatICAAccountOwner(chainId, types.ICAType_ICQ)
	if err := k.ICAControllerKeeper.RegisterInterchainAccount(ctx, hostZone.ConnectionId, icqAccount, appVersion); err != nil {
		errMsg := fmt.Sprintf("unable to register icq account, err: %s", err.Error())
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(types.ErrFailedToRegisterHostZone, errMsg)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRegisterZone,
			sdk.NewAttribute(types.AttributeKeyRecipientChain, chainId),
			sdk.NewAttribute(types.AttributeKeyConnectionId, msg.ConnectionId),
		),
	)

	return &types.MsgRegisterHostZoneResponse{}, nil
}

func (k Keeper) SetHostZone(ctx sdk.Context, hostZone types.HostZone) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.HostZoneKey))
	b := k.cdc.MustMarshal(&hostZone)
	store.Set([]byte(hostZone.ChainId), b)
}
