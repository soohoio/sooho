package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	"github.com/soohoio/stayking/v2/utils"
	epochstypes "github.com/soohoio/stayking/v2/x/epochs/types"
	icacallbackstypes "github.com/soohoio/stayking/v2/x/icacallbacks/types"
	icqtypes "github.com/soohoio/stayking/v2/x/interchainquery/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func (k Keeper) SubmitTxsEpoch(
	ctx sdk.Context,
	connectionId string,
	msgs []sdk.Msg,
	account types.ICAAccount,
	epochType string,
	callbackId string,
	callbackArgs []byte,
) (uint64, error) {
	timeoutNanosUint64, err := k.GetICATimeoutNanos(ctx, epochType)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Failed to get ICA timeout nanos for epochType %s using param, error: %s", epochType, err.Error()))
		return 0, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Failed to convert timeoutNanos to uint64, error: %s", err.Error())
	}
	sequence, err := k.SubmitTxs(ctx, connectionId, msgs, account, timeoutNanosUint64, callbackId, callbackArgs)
	if err != nil {
		return 0, err
	}
	return sequence, nil
}

// SubmitTxs submits an ICA transaction containing multiple messages
func (k Keeper) SubmitTxs(
	ctx sdk.Context,
	connectionId string,
	msgs []sdk.Msg,
	account types.ICAAccount,
	timeoutTimestamp uint64,
	callbackId string,
	callbackArgs []byte,
) (uint64, error) {
	chainId, err := k.GetChainID(ctx, connectionId)
	if err != nil {
		return 0, err
	}
	owner := types.FormatICAAccountOwner(chainId, account.Target)
	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return 0, err
	}

	k.Logger(ctx).Info(utils.LogWithHostZone(chainId, "  Submitting ICA Tx on %s, %s with TTL: %d", portID, connectionId, timeoutTimestamp))
	for _, msg := range msgs {
		k.Logger(ctx).Info(utils.LogWithHostZone(chainId, "    Msg: %+v", msg))
	}

	channelID, found := k.ICAControllerKeeper.GetActiveChannelID(ctx, connectionId, portID)
	if !found {
		return 0, sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", portID)
	}

	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channelID))
	if !found {
		return 0, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	data, err := icatypes.SerializeCosmosTx(k.cdc, msgs)
	if err != nil {
		return 0, err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}
	sequence, err := k.ICAControllerKeeper.SendTx(ctx, chanCap, connectionId, portID, packetData, timeoutTimestamp)
	if err != nil {
		return 0, err
	}

	// Store the callback data
	if callbackId != "" && callbackArgs != nil {
		callback := icacallbackstypes.CallbackData{
			CallbackKey:  icacallbackstypes.PacketID(portID, channelID, sequence),
			PortId:       portID,
			ChannelId:    channelID,
			Sequence:     sequence,
			CallbackId:   callbackId,
			CallbackArgs: callbackArgs,
		}
		k.Logger(ctx).Info(utils.LogWithHostZone(chainId, "Storing callback data: %+v", callback))
		k.ICACallbacksKeeper.SetCallbackData(ctx, callback)
	}

	return sequence, nil
}

func (k Keeper) SubmitICQWithWithdrawalBalance(ctx sdk.Context, hostZone types.HostZone) error {
	k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Submitting ICQ for withdrawal account balance"))

	withdrawalIca := hostZone.WithdrawalAccount
	if withdrawalIca == nil || withdrawalIca.Address == "" {
		k.Logger(ctx).Error(fmt.Sprintf("Zone %s is missing a withdrawal address!", hostZone.ChainId))
	}

	k.Logger(ctx).Info(utils.LogWithHostZone("[CUSTOM DEBUG] withdrawalIca:: ", withdrawalIca.String()))
	k.Logger(ctx).Info(utils.LogWithHostZone("[CUSTOM DEBUG] withdrawalIca.Address:: ", withdrawalIca.Address))

	_, addr, _ := bech32.DecodeAndConvert(withdrawalIca.Address)
	data := bankTypes.CreateAccountBalancesPrefix(addr)

	k.Logger(ctx).Info(fmt.Sprintf("[CUSTOM DEBUG] addr:: %v", addr))
	k.Logger(ctx).Info(fmt.Sprintf("[CUSTOM DEBUG] data:: %v", data))

	// get ttl, the end of the ICA buffer window
	ttl, err := k.GetICATimeoutNanos(ctx, epochstypes.STAYKING_EPOCH)
	k.Logger(ctx).Info(fmt.Sprintf("[CUSTOM DEBUG] epochType:: %s", epochstypes.STAYKING_EPOCH))
	k.Logger(ctx).Info(fmt.Sprintf("[CUSTOM DEBUG] ttl:: %d", ttl))
	k.Logger(ctx).Info(fmt.Sprintf("[CUSTOM DEBUG] HostDenom:: %s", hostZone.HostDenom))

	if err != nil {
		errMsg := fmt.Sprintf("Failed to get ICA timeout nanos for epochType %s using param, error: %s", epochstypes.STAYKING_EPOCH, err.Error())
		k.Logger(ctx).Error(errMsg)
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, errMsg)
	}

	err = k.InterchainQueryKeeper.MakeRequest(
		ctx,
		types.ModuleName,
		ICQCallbackID_WithdrawalBalance,
		hostZone.ChainId,
		hostZone.ConnectionId,
		// use "bank" store to access acct balances which live in the bank module
		// use "key" suffix to retrieve a proof alongside the query result
		icqtypes.BANK_STORE_QUERY_WITH_PROOF,
		append(data, []byte(hostZone.HostDenom)...),
		ttl, // ttl
	)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Error querying for withdrawal balance, error: %s", err.Error()))
		return err
	}
	return nil
}
