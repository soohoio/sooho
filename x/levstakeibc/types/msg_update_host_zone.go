package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ibctransfertypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	"strings"
)

const TypeMsgUpdateHostZone = "update_host_zone"

var _ sdk.Msg = &MsgUpdateHostZone{}

func NewMsgUpdateHostZone(connectionId string, bech32Prefix string, hostDenom string, ibcDenom string, creator string, transferChannelId string, unbondingFrequency uint64) *MsgUpdateHostZone {
	return &MsgUpdateHostZone{
		ConnectionId:       connectionId,
		Bech32Prefix:       bech32Prefix,
		HostDenom:          hostDenom,
		IbcDenom:           ibcDenom,
		Creator:            creator,
		TransferChannelId:  transferChannelId,
		UnbondingFrequency: unbondingFrequency,
	}
}

func (msg *MsgUpdateHostZone) Route() string {
	return RouterKey
}

func (msg *MsgUpdateHostZone) Type() string {
	return TypeMsgUpdateHostZone
}

func (msg *MsgUpdateHostZone) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.HostDenom == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "host denom cannot be empty")
	}

	if err := sdk.ValidateDenom(msg.HostDenom); err != nil {
		return err
	}

	if msg.IbcDenom == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "ibc denom cannot be empty")
	}

	if !strings.HasPrefix(msg.IbcDenom, "ibc") {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "ibc denom must begin with 'ibc'")
	}

	err = ibctransfertypes.ValidateIBCDenom(msg.IbcDenom)
	if err != nil {
		return err
	}

	if strings.TrimSpace(msg.Bech32Prefix) == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "bech32 prefix must be non-empty")
	}

	if msg.ConnectionId == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "connection id cannot be empty")
	}
	if !strings.HasPrefix(msg.ConnectionId, "connection") {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "connection id must begin with 'connection'")
	}

	if msg.TransferChannelId == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "transfer channel id cannot be empty")
	}
	if !strings.HasPrefix(msg.TransferChannelId, "channel") {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "transfer channel id must begin with 'channel'")
	}

	if msg.UnbondingFrequency < 1 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "unbonding frequency must be greater than zero")
	}

	return nil
}

func (msg *MsgUpdateHostZone) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}
