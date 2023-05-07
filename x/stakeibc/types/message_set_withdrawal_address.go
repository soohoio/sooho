package types

import (
	//distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	//"github.com/cosmos/cosmos-sdk/x/distribution/types"
	//"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetWithdrawalAddress = "set_withdrawal_address"

var _ sdk.Msg = &MsgSetWithdrawalAddress{}

func NewMsgSetWithdrawalAddress(creator string) *MsgSetWithdrawalAddress {
	return &MsgSetWithdrawalAddress{
		Creator: creator,
	}
}

func (msg *MsgSetWithdrawalAddress) Route() string {
	return RouterKey
}

func (msg *MsgSetWithdrawalAddress) Type() string {
	return TypeMsgSetWithdrawalAddress
}

func (msg *MsgSetWithdrawalAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetWithdrawalAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetWithdrawalAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	//if err := utils.ValidateAdminAddress(msg.Creator); err != nil {
	//	return err
	//}
	return nil
}
