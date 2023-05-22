package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddValidator = "register_interchain_account"

var _ sdk.Msg = &MsgAddValidator{}

func NewMsgAddValidator(creator string, hostZone string, name string, address string, commission uint64, weight uint64) *MsgAddValidator {
	return &MsgAddValidator{
		Creator:    creator,
		HostZone:   hostZone,
		Name:       name,
		Address:    address,
		Commission: commission,
		Weight:     weight,
	}
}

func (msg *MsgAddValidator) Route() string {
	return RouterKey
}

func (msg *MsgAddValidator) Type() string {
	return TypeMsgAddValidator
}

func (msg *MsgAddValidator) ValidateBasic() error {

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.HostZone == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "host zone can not be empty")
	}

	if msg.Name == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "validator name can not be empty")
	}

	if msg.Address == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "validator address can not be empty")
	}

	if msg.Commission > 100 || msg.Commission <= 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "commission value must be between 0 and 100")
	}

	if msg.Weight <= 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "weight value must be greater than 0")
	}

	return nil
}

func (msg *MsgAddValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}
