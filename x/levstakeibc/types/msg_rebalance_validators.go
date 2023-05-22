package types

import (
	errorsmod "cosmossdk.io/errors"
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRebalanceValidators = "rebalance_validators"

var _ sdk.Msg = &MsgRebalanceValidators{}

func NewMsgRebalanceValidators(creator string, hostZone string, numValidators uint64) *MsgRebalanceValidators {
	return &MsgRebalanceValidators{
		Creator:      creator,
		HostZone:     hostZone,
		NumRebalance: numValidators,
	}
}

func (msg *MsgRebalanceValidators) Route() string {
	return RouterKey
}

func (msg *MsgRebalanceValidators) Type() string {
	return TypeMsgRebalanceValidators
}

func (msg *MsgRebalanceValidators) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRebalanceValidators) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRebalanceValidators) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.HostZone == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "host zone can not be empty")
	}

	if (msg.NumRebalance <= 0) || (msg.NumRebalance > 10) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid number of validators to rebalance (%d)", msg.NumRebalance))
	}
	return nil
}
