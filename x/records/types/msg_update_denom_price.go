package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateDenomPrice = "update_denom_price"

var _ sdk.Msg = &MsgUpdateDenomPrice{}

func NewMsgUpdateDenomPrice(creator string, baseDenom string, targetDenom string, denomPrice sdk.Int) *MsgUpdateDenomPrice {
	return &MsgUpdateDenomPrice{
		Creator:     creator,
		BaseDenom:   baseDenom,
		TargetDenom: targetDenom,
		DenomPrice:  denomPrice,
	}
}

func (m *MsgUpdateDenomPrice) Route() string {
	return RouterKey
}

func (m *MsgUpdateDenomPrice) Type() string {
	return TypeMsgUpdateDenomPrice
}

func (m *MsgUpdateDenomPrice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgUpdateDenomPrice) ValidateBasic() error {
	if m.GetBaseDenom() == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid request (base_denom empty)")
	}

	if m.GetTargetDenom() == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid request (target_denom empty)")
	}

	if m.DenomPrice.LTE(sdk.NewInt(0)) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "denom price should be greater than zero(0)")
	}

	return nil
}

func (m *MsgUpdateDenomPrice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic("invalid address")
	}
	return []sdk.AccAddress{creator}
}
