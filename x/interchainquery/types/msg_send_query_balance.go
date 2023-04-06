package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendQueryBalance = "send_query_balance"

var _ sdk.Msg = &MsgSendQueryBalance{}

func NewMsgSendQueryBalance(creator, channelId string, addr string, denom string) *MsgSendQueryBalance {
	return &MsgSendQueryBalance{
		Creator:   creator,
		ChannelId: channelId,
		Address:   addr,
		Denom:     denom,
	}
}

func (msg *MsgSendQueryBalance) Route() string {
	return RouterKey
}

func (msg *MsgSendQueryBalance) Type() string {
	return TypeMsgSendQueryBalance
}

func (msg *MsgSendQueryBalance) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendQueryBalance) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendQueryBalance) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
