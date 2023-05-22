package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgChangeValidatorWeight = "change_validator_weight"

var _ sdk.Msg = &MsgChangeValidatorWeight{}

func NewMsgChangeValidatorWeight(creator string, hostZone string, address string, weight uint64) *MsgChangeValidatorWeight {
	return &MsgChangeValidatorWeight{
		Creator:  creator,
		HostZone: hostZone,
		ValAddr:  address,
		Weight:   weight,
	}
}

func (msg *MsgChangeValidatorWeight) Route() string {
	return RouterKey
}

func (msg *MsgChangeValidatorWeight) Type() string {
	return TypeMsgChangeValidatorWeight
}

func (msg *MsgChangeValidatorWeight) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.HostZone == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "host zone can not be empty")
	}

	if msg.ValAddr == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "validator address can not be empty")
	}

	if msg.Weight <= 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "weight value must be greater than 0")
	}

	return nil
}

func (msg *MsgChangeValidatorWeight) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChangeValidatorWeight) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
