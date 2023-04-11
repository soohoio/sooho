package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRedeemStake = "redeem_stake"

var _ sdk.Msg = &MsgRedeemStake{}

func NewMsgRedeemStake(creator string, stTokenAmount sdk.Int, chainId string, receiver string) *MsgRedeemStake {
	return &MsgRedeemStake{
		Creator:       creator,
		StTokenAmount: stTokenAmount,
		ChainId:       chainId,
		Receiver:      receiver,
	}
}

func (m *MsgRedeemStake) Route() string {
	return RouterKey
}

func (m *MsgRedeemStake) Type() string {
	return TypeMsgRedeemStake
}

func (m *MsgRedeemStake) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if m.GetReceiver() == "" {
		return errorsmod.Wrapf(ErrRequiredFieldEmpty, "receiver cannot be empty")
	}
	if m.StTokenAmount.LTE(sdk.ZeroInt()) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid st token amount (%v)", m.StTokenAmount)
	}
	if m.GetChainId() == "" {
		return errorsmod.Wrapf(ErrRequiredFieldEmpty, "host zone cannot be empty")
	}
	return nil
}

func (m *MsgRedeemStake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (m *MsgRedeemStake) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}
