package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	admintypes "github.com/soohoio/stayking/v2/x/admin/types"
	epochtypes "github.com/soohoio/stayking/v2/x/epochs/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	recordstypes "github.com/soohoio/stayking/v2/x/records/types"
)

func (k msgServer) RegisterHostZone(goCtx context.Context, msg *types.MsgRegisterHostZone) (*types.MsgRegisterHostZoneResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// admin address check
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if !k.AdminKeeper.IsAdmin(ctx, creator) {
		return nil, admintypes.ErrNotAdmin
	}

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

	appVersion := string(icatypes.ModuleCdc.MustMarshalJSON(&icatypes.Metadata{
		Version:                icatypes.Version,
		ControllerConnectionId: msg.ConnectionId,
		HostConnectionId:       connectionEnd.Counterparty.ConnectionId,
		Encoding:               icatypes.EncodingProtobuf,
		TxType:                 icatypes.TxTypeSDKMultiMsg,
	}))

	delegateAccount := types.FormatICAAccountOwner(chainId, types.ICAType_DELEGATION)
	portID, err := icatypes.NewControllerPortID(delegateAccount)

	delegateICA, found := k.ICAControllerKeeper.GetInterchainAccountAddress(ctx, msg.ConnectionId, portID)

	if found {
		hostZone.DelegationAccount = &types.ICAAccount{Address: delegateICA, Target: types.ICAType_DELEGATION}
	} else {
		if err := k.ICAControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, delegateAccount, appVersion); err != nil {
			errMsg := fmt.Sprintf("unable to register delegation account, err: %s", err.Error())
			k.Logger(ctx).Error(errMsg)
			return nil, errorsmod.Wrapf(types.ErrFailedToRegisterHostZone, errMsg)
		}
	}

	withdrawalAccount := types.FormatICAAccountOwner(chainId, types.ICAType_WITHDRAWAL)
	portID, err = icatypes.NewControllerPortID(withdrawalAccount)

	withdrawalICA, found := k.ICAControllerKeeper.GetInterchainAccountAddress(ctx, msg.ConnectionId, portID)
	if found {
		hostZone.WithdrawalAccount = &types.ICAAccount{Address: withdrawalICA, Target: types.ICAType_WITHDRAWAL}
	} else {
		if err := k.ICAControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, withdrawalAccount, appVersion); err != nil {
			errMsg := fmt.Sprintf("unable to register withdrawal account, err: %s", err.Error())
			k.Logger(ctx).Error(errMsg)
			return nil, errorsmod.Wrapf(types.ErrFailedToRegisterHostZone, errMsg)
		}
	}

	feeAccount := types.FormatICAAccountOwner(chainId, types.ICAType_FEE)
	portID, err = icatypes.NewControllerPortID(feeAccount)

	feeICA, found := k.ICAControllerKeeper.GetInterchainAccountAddress(ctx, msg.ConnectionId, portID)
	if found {
		hostZone.FeeAccount = &types.ICAAccount{Address: feeICA, Target: types.ICAType_FEE}
	} else {
		if err := k.ICAControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, feeAccount, appVersion); err != nil {
			errMsg := fmt.Sprintf("unable to register fee account, err: %s", err.Error())
			k.Logger(ctx).Error(errMsg)
			return nil, errorsmod.Wrapf(types.ErrFailedToRegisterHostZone, errMsg)
		}
	}

	k.SetHostZone(ctx, hostZone)

	// add this host zone to unbonding hostZones, otherwise users won't be able to unbond
	// for this host zone until the following day
	dayEpochTracker, found := k.GetEpochTracker(ctx, epochtypes.DAY_EPOCH)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrEpochNotFound, "epoch tracker (%s) not found", epochtypes.DAY_EPOCH)
	}
	epochUnbondingRecord, found := k.RecordsKeeper.GetEpochUnbondingRecord(ctx, dayEpochTracker.EpochNumber)
	if !found {
		errMsg := "unable to find latest epoch unbonding record"
		k.Logger(ctx).Error(errMsg)
		return nil, errorsmod.Wrapf(recordstypes.ErrEpochUnbondingRecordNotFound, errMsg)
	}
	hostZoneUnbonding := &recordstypes.HostZoneUnbonding{
		NativeTokenAmount: sdk.ZeroInt(),
		StTokenAmount:     sdk.ZeroInt(),
		Denom:             hostZone.HostDenom,
		HostZoneId:        hostZone.ChainId,
		Status:            recordstypes.HostZoneUnbonding_UNBONDING_QUEUE,
	}
	updatedEpochUnbondingRecord, success := k.RecordsKeeper.AddHostZoneToEpochUnbondingRecord(ctx, epochUnbondingRecord.EpochNumber, chainId, hostZoneUnbonding)
	if !success {
		return nil, errorsmod.Wrapf(types.ErrFailureUpdateUnbondingRecord, "Failed to set host zone epoch unbonding record: epochNumber %d, chainId %s, hostZoneUnbonding %v.",
			epochUnbondingRecord.EpochNumber, chainId, hostZoneUnbonding)
	}
	k.RecordsKeeper.SetEpochUnbondingRecord(ctx, *updatedEpochUnbondingRecord)

	// create an empty deposit record for the host zone
	staykingEpochTracker, found := k.GetEpochTracker(ctx, epochtypes.STAYKING_EPOCH)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrEpochNotFound, "epoch tracker (%s) not found", epochtypes.STAYKING_EPOCH)
	}
	depositRecord := recordstypes.DepositRecord{
		Id:                 0,
		Amount:             sdk.ZeroInt(),
		Denom:              hostZone.HostDenom,
		HostZoneId:         hostZone.ChainId,
		Status:             recordstypes.DepositRecord_TRANSFER_QUEUE,
		DepositEpochNumber: staykingEpochTracker.EpochNumber,
	}
	k.RecordsKeeper.AppendDepositRecord(ctx, depositRecord)

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
