package levstakeibc

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/soohoio/stayking/v2/x/icacallbacks"
	//icqtypes "github.com/cosmos/ibc-go/v5/modules/types"
	//errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v5/modules/core/exported"
	//"github.com/soohoio/stayking/v2/x/icacallbacks"
	icacallbacktypes "github.com/soohoio/stayking/v2/x/icacallbacks/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/keeper"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

type IBCModule struct {
	keeper keeper.Keeper
}

func NewIBCModule(k keeper.Keeper) IBCModule {
	return IBCModule{
		keeper: k,
	}
}

func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	im.keeper.Logger(ctx).Info(fmt.Sprintf("LevstakeIBC >> OnChanOpenInit: portID %s, channelID %s, oerderType %s, version %s", portID, channelID, order, version))

	if err := im.keeper.ClaimCapability(ctx, channelCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
		return version, err
	}

	return version, nil

}

// it must not be implemented in ICS27 (ICA)
func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (version string, err error) {
	panic("UNIMPLEMENTED")
}

func (im IBCModule) OnChanOpenAck(
	ctx sdk.Context,
	portID string,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	im.keeper.Logger(ctx).Info(fmt.Sprintf("LevstakeIBC >> OnChanOpenAck: portID %s, channelID %s, counterpartyChannelID %s, counterpartyVersion %s", portID, channelID, counterpartyChannelID, counterpartyVersion))
	contConnId, err := im.keeper.GetConnectionId(ctx, portID)
	if err != nil {
		ctx.Logger().Error(fmt.Sprintf("Error: %s", err.Error()))
	}
	address, found := im.keeper.ICAControllerKeeper.GetInterchainAccountAddress(ctx, contConnId, portID)
	if !found {
		ctx.Logger().Error(fmt.Sprintf("There is no ica address for controllerConnId: %s, portId: %s", contConnId, portID))
	}
	hostChainId, err := im.keeper.GetChainID(ctx, contConnId)
	if err != nil {
		ctx.Logger().Error(fmt.Sprintf("Unable to obtain counterparty chain for connection: %s, port: %s, err: %s", contConnId, portID, err.Error()))
		return nil
	}
	zoneInfo, found := im.keeper.GetHostZone(ctx, hostChainId)
	if !found {
		ctx.Logger().Error(fmt.Sprintf("Expected to find zone info for %v", hostChainId))
		return nil
	}
	ctx.Logger().Info(fmt.Sprintf("Found matching address for chain: %s, address %s, port %s", zoneInfo.ChainId, address, portID))

	// addresses
	delegationICA, err := icatypes.NewControllerPortID(types.FormatICAAccountOwner(hostChainId, types.ICAType_DELEGATION))
	if err != nil {
		return err
	}
	withdrawalICA, err := icatypes.NewControllerPortID(types.FormatICAAccountOwner(hostChainId, types.ICAType_WITHDRAWAL))
	if err != nil {
		return err
	}
	redemptionICA, err := icatypes.NewControllerPortID(types.FormatICAAccountOwner(hostChainId, types.ICAType_REDEMPTION))
	if err != nil {
		return err
	}
	feeICA, err := icatypes.NewControllerPortID(types.FormatICAAccountOwner(hostChainId, types.ICAType_FEE))
	if err != nil {
		return err
	}
	icqICA, err := icatypes.NewControllerPortID(types.FormatICAAccountOwner(hostChainId, types.ICAType_ICQ))
	if err != nil {
		return err
	}

	switch {
	case portID == delegationICA:
		zoneInfo.DelegationAccount = &types.ICAAccount{Address: address, Target: types.ICAType_DELEGATION}
	case portID == withdrawalICA:
		zoneInfo.WithdrawalAccount = &types.ICAAccount{Address: address, Target: types.ICAType_WITHDRAWAL}
	case portID == redemptionICA:
		zoneInfo.RedemptionAccount = &types.ICAAccount{Address: address, Target: types.ICAType_REDEMPTION}
	case portID == feeICA:
		zoneInfo.FeeAccount = &types.ICAAccount{Address: address, Target: types.ICAType_FEE}
	case portID == icqICA:
		zoneInfo.IcqAccount = &types.ICAAccount{Address: address, Target: types.ICAType_ICQ}
	default:
		ctx.Logger().Error(fmt.Sprintf("Missing portId: %s", portID))
	}

	im.keeper.SetHostZone(ctx, zoneInfo)
	return nil
}

// it must not be implemented
func (im IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID string,
	channelID string,
) error {
	panic("implement me")
}

// it must not be implemented
func (im IBCModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	panic("UNIMPLEMENTED")
}

func (im IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {

	// WARNING: For some reason, in IBCv3 the ICA controller module does not call the underlying OnChanCloseConfirm (this function)
	// So, we need to put logic that _should_ execute upon channel closure in the OnTimeoutPacket function
	// This works because ORDERED channels are always closed when a timeout occurs, but if we migrate to using ORDERED channels that don't
	// close on timeout, we will need to move this logic to the OnChanCloseConfirm function
	// relevant IBCv3 code: https://github.com/cosmos/ibc-go/blob/5c0bf8b8a0f79643e36be98fb9883ea163d2d93a/modules/apps/27-interchain-accounts/controller/ibc_module.go#L123
	return nil
}

// it must not be implemented
func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	panic("UNIMPLEMENTED")
}

//func (im IBCModule) OnAcknowledgementPacket(
//	ctx sdk.Context,
//	modulePacket channeltypes.Packet,
//	acknowledgement []byte,
//	relayer sdk.AccAddress,
//) error {
//	im.keeper.Logger(ctx).Info(fmt.Sprintf("OnAcknowledgementPacket (Levstakeibc) - packet: %+v, relayer: %v", modulePacket, relayer))
//	var ack channeltypes.Acknowledgement
//	if err := ibctransfertypes.ModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
//		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet acknowledgement: %v", err)
//	}
//
//	return im.keeper.OnAcknowledgementPacket(ctx, modulePacket, ack)
//
//}

func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	im.keeper.Logger(ctx).Info(fmt.Sprintf("1: OnAcknowledgementPacket (Levstakeibc) - packet: %+v, relayer: %v", modulePacket, relayer))
	ackResponse, err := icacallbacks.UnpackAcknowledgementResponse(ctx, im.keeper.Logger(ctx), acknowledgement, true)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to unpack message data from acknowledgement, Sequence %d, from %s %s, to %s %s: %s",
			modulePacket.Sequence, modulePacket.SourceChannel, modulePacket.SourcePort, modulePacket.DestinationChannel, modulePacket.DestinationPort, err.Error())
		im.keeper.Logger(ctx).Error(errMsg)
		return errorsmod.Wrapf(icacallbacktypes.ErrInvalidAcknowledgement, errMsg)
	}

	ackInfo := fmt.Sprintf("sequence #%d, from %s %s, to %s %s",
		modulePacket.Sequence, modulePacket.SourceChannel, modulePacket.SourcePort, modulePacket.DestinationChannel, modulePacket.DestinationPort)
	im.keeper.Logger(ctx).Info(fmt.Sprintf("Acknowledgement was successfully unmarshalled: ackInfo: %s", ackInfo))
	eventType := "ack"
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			eventType,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyAck, ackInfo),
		),
	)
	err = im.keeper.ICACallbacksKeeper.CallRegisteredICACallback(ctx, modulePacket, ackResponse)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to call registered callback from stakeibc OnAcknowledgePacket | Sequence %d, from %s %s, to %s %s",
			modulePacket.Sequence, modulePacket.SourceChannel, modulePacket.SourcePort, modulePacket.DestinationChannel, modulePacket.DestinationPort)
		im.keeper.Logger(ctx).Error(errMsg)
		return errorsmod.Wrapf(icacallbacktypes.ErrCallbackFailed, errMsg)
	}

	return nil
}

func (im IBCModule) OnTimeoutPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	im.keeper.Logger(ctx).Info(fmt.Sprintf("OnTimeoutPacket: packet %v, relayer %v", modulePacket, relayer))
	ackResponse := icacallbacktypes.AcknowledgementResponse{Status: icacallbacktypes.AckResponseStatus_TIMEOUT}
	err := im.keeper.ICACallbacksKeeper.CallRegisteredICACallback(ctx, modulePacket, &ackResponse)
	if err != nil {
		return err
	}
	return nil
}

// ###################################################################################
// 	Helper functions
// ###################################################################################

func (im IBCModule) NegotiateAppVersion(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionID string,
	portID string,
	counterparty channeltypes.Counterparty,
	proposedVersion string,
) (version string, err error) {
	return proposedVersion, nil
}
