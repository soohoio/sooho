package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	"github.com/soohoio/stayking/v2/x/interchainquery/types"
	recordstypes "github.com/soohoio/stayking/v2/x/records/types"
	icqtypes "github.com/strangelove-ventures/async-icq/v5/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (k Keeper) SendQuery(
	ctx sdk.Context,
	sourcePort,
	sourceChannel string,
	chanCap *capabilitytypes.Capability,
	reqs []abci.RequestQuery,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) (uint64, error) {
	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return 0, sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	data, err := icqtypes.SerializeCosmosQuery(reqs)
	if err != nil {
		return 0, sdkerrors.Wrap(err, "could not serialize reqs into cosmos query")
	}
	icqPacketData := icqtypes.InterchainQueryPacketData{
		Data: data,
	}

	return k.createOutgoingPacket(ctx, sourcePort, sourceChannel, destinationPort, destinationChannel, chanCap, icqPacketData, timeoutTimestamp)
}

func (k Keeper) createOutgoingPacket(
	ctx sdk.Context,
	sourcePort,
	sourceChannel,
	destinationPort,
	destinationChannel string,
	chanCap *capabilitytypes.Capability,
	icqPacketData icqtypes.InterchainQueryPacketData,
	timeoutTimestamp uint64,
) (uint64, error) {
	if err := icqPacketData.ValidateBasic(); err != nil {
		return 0, sdkerrors.Wrap(err, "invalid interchain query packet data")
	}

	// get the next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return 0, sdkerrors.Wrapf(channeltypes.ErrSequenceSendNotFound, "failed to retrieve next sequence send for channel %s on port %s", sourceChannel, sourcePort)
	}

	packet := channeltypes.NewPacket(
		icqPacketData.GetBytes(),
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		clienttypes.ZeroHeight(),
		timeoutTimestamp,
	)

	if err := k.ics4Wrapper.SendPacket(ctx, chanCap, packet); err != nil {
		return 0, err
	}

	return packet.Sequence, nil
}

func (k Keeper) OnAcknowledgementPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	ack channeltypes.Acknowledgement,
) error {
	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:

		var ackData icqtypes.InterchainQueryPacketAck
		if err := icqtypes.ModuleCdc.UnmarshalJSON(resp.Result, &ackData); err != nil {
			return sdkerrors.Wrap(err, "failed to unmarshal interchain query packet ack")
		}
		resps, err := icqtypes.DeserializeCosmosResponse(ackData.Data)
		if err != nil {
			return sdkerrors.Wrap(err, "could not deserialize data to cosmos response")
		}

		if len(resps) < 1 {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no responses in interchain query packet ack")
		}

		var r types.EstimateSwapExactAmountOutResponse
		if err := k.cdc.Unmarshal(resps[0].Value, &r); err != nil {
			return sdkerrors.Wrapf(err, "failed to unmarshal interchain query response to type %T", resp)
		}

		var packetData icqtypes.InterchainQueryPacketData
		if err := types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &packetData); err != nil {
			return sdkerrors.Wrapf(err, "could not unmarshal icq packet data")
		}

		k.SetQueryResponse(ctx, modulePacket.Sequence, r)
		k.SetLastQueryPacketSeq(ctx, modulePacket.Sequence)

		if err := k.handleOsmosisPriceQueryResponse(ctx, resps[0]); err != nil {
			return sdkerrors.Wrapf(err, "could not handle icq response of request")
		}
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeQueryResult,
				sdk.NewAttribute(types.AttributeKeyAckSuccess, string(resp.Result)),
			),
		)
	case *channeltypes.Acknowledgement_Error:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeQueryResult,
				sdk.NewAttribute(types.AttributeKeyAckError, resp.Error),
			),
		)

		k.Logger(ctx).Error("interchain query response", "sequence", modulePacket.Sequence, "error", resp.Error)
	}
	return nil
}

func (k Keeper) handleOsmosisPriceQueryResponse(ctx sdk.Context, resp abci.ResponseQuery) error {
	if resp.IsErr() {
		return errorsmod.Wrapf(types.ErrFailedICQResponse, "icq response failed with code %d", resp.GetCode())
	}

	var qresp types.EstimateSwapExactAmountOutResponse
	k.cdc.MustUnmarshal(resp.GetValue(), &qresp)
	var denomPriceRecord recordstypes.DenomPriceRecord
	//@TODO EVMOS 외 다른 자산 추가시 분기처리 필요
	denomPriceRecord.BaseDenom = "uosmo"
	denomPriceRecord.TargetDenom = "USDC" //@TODO axl USDC denom ibc/ 포맷으로 변경
	denomPriceRecord.DenomPrice = qresp.TokenInAmount
	denomPriceRecord.Timestamp = uint64(ctx.BlockTime().UnixNano())
	k.RecordsKeeper.SetDenomPriceRecord(ctx, denomPriceRecord)
	return nil
}
