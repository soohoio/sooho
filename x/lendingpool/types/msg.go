package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgCreatePool = "create_pool"
)

// NewMsgCreatePool creates a new NewMsgCreatePool instance.
func NewMsgCreatePool(creator, denom string, interestRate sdk.Dec) *MsgCreatePool {
	return &MsgCreatePool{
		Creator:      creator,
		Denom:        denom,
		InterestRate: interestRate,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgCreatePool) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgCreatePool) Type() string { return TypeMsgCreatePool }

// GetSigners implements the sdk.Msg interface.
func (msg MsgCreatePool) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgCreatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgCreatePool) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}
	if from.Empty() {
		return ErrEmptyCreator
	}

	if msg.Denom == "" {
		return ErrEmptyDenom
	}
	if msg.InterestRate.LTE(sdk.ZeroDec()) {
		return ErrInvalidInterestRate
	}
	return nil
}
