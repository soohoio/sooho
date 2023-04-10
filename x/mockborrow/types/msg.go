package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgBorrow = "borrow"
	TypeMsgRepay  = "repay"
)

var (
	_ sdk.Msg = &MsgBorrow{}
	_ sdk.Msg = &MsgRepay{}
)

// NewMsgBorrow creates a new NewMsgBorrow instance.
func NewMsgBorrow(from, denom string, collateral, borrow sdk.Coins) *MsgBorrow {
	msg := &MsgBorrow{
		From:         from,
		Denom:        denom,
		Collateral:   collateral,
		BorrowAmount: borrow,
	}
	return msg
}

// Route implements the sdk.Msg interface.
func (msg MsgBorrow) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgBorrow) Type() string { return TypeMsgBorrow }

// GetSigners implements the sdk.Msg interface.
func (msg MsgBorrow) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgBorrow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgBorrow) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return err
	}
	if from.Empty() {
		return ErrInvalidFromAddr
	}

	if msg.Denom == "" {
		return ErrEmptyDenom
	}

	if len(msg.Collateral) == 0 {
		return ErrEmptyCollateral
	}

	if len(msg.BorrowAmount) == 0 {
		return ErrEmptyBorrowAmount
	}
	return nil
}

// NewMsgRepay creates a new NewMsgBorrow instance.
func NewMsgRepay(from string, id uint64, amount sdk.Coins) *MsgRepay {
	msg := &MsgRepay{
		From:   from,
		Id:     id,
		Amount: amount,
	}
	return msg
}

// Route implements the sdk.Msg interface.
func (msg MsgRepay) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgRepay) Type() string { return TypeMsgRepay }

// GetSigners implements the sdk.Msg interface.
func (msg MsgRepay) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgRepay) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgRepay) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return err
	}
	if from.Empty() {
		return ErrInvalidFromAddr
	}

	if msg.Id == 0 {
		return ErrInvalidLoanId
	}

	if len(msg.Amount) == 0 {
		return ErrInvalidAmount
	}

	return nil
}
