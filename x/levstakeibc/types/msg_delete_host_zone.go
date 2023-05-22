package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteHostZone = "update_host_zone"

var _ sdk.Msg = &MsgDeleteHostZone{}

func NewMsgDeleteHostZone(creator string, chainId string) *MsgDeleteHostZone {
	return &MsgDeleteHostZone{
		Creator: creator,
		ChainId: chainId,
	}
}

func (msg *MsgDeleteHostZone) Route() string {
	return RouterKey
}

func (msg *MsgDeleteHostZone) Type() string {
	return TypeMsgDeleteHostZone
}

func (msg *MsgDeleteHostZone) ValidateBasic() error {

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.ChainId == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "chain id cannot be empty")
	}

	return nil
}

func (msg *MsgDeleteHostZone) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}
