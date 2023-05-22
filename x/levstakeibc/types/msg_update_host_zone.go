package types

import (
	errorsmod "cosmossdk.io/errors"
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
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.ConnectionId == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "connection id cannot be empty")
	}
	if !strings.HasPrefix(msg.ConnectionId, "connection") {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "connection id must begin with 'connection'")
	}

	if msg.HostDenom == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "host denom cannot be empty")
	}
	if err := sdk.ValidateDenom(msg.HostDenom); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "host denom is not valid")
	}

	if strings.TrimSpace(msg.Bech32Prefix) == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "bech32 prefix must be non-empty")
	}

	if msg.IbcDenom == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "ibc denom cannot be empty")
	}
	if !strings.HasPrefix(msg.IbcDenom, "ibc") {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "ibc denom must begin with 'ibc'")
	}

	err = ibctransfertypes.ValidateIBCDenom(msg.IbcDenom)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "ibc denom is not valid name")
	}

	if msg.TransferChannelId == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "transfer channel id cannot be empty")
	}

	if !strings.HasPrefix(msg.TransferChannelId, "channel") {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "transfer channel id must begin with 'channel'")
	}

	if msg.UnbondingFrequency < 1 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "unbonding frequency must be greater than zero")
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
