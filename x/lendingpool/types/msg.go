package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgCreatePool = "create_pool"
	TypeMsgDeposit    = "deposit"
	TypeMsgWithdraw   = "withdraw"
)

var (
	_ sdk.Msg = &MsgCreatePool{}
	_ sdk.Msg = &MsgDeposit{}
	_ sdk.Msg = &MsgWithdraw{}
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

// NewMsgDeposit creates a new NewMsgCreatePool instance.
func NewMsgDeposit(from string, poolID uint64, amount sdk.Coins) *MsgDeposit {
	return &MsgDeposit{
		From:   from,
		PoolId: poolID,
		Amount: amount,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgDeposit) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgDeposit) Type() string { return TypeMsgDeposit }

// GetSigners implements the sdk.Msg interface.
func (msg MsgDeposit) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgDeposit) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return err
	}
	if from.Empty() {
		return ErrInvalidDepositor
	}

	if msg.PoolId == 0 {
		return ErrInvalidPoolID
	}

	// only accept one coin denom at a time
	if len(msg.Amount) > 1 || len(msg.Amount) == 0 {
		return ErrInvalidDepositCoins
	}
	return nil
}

// NewMsgWithdraw creates a new NewMsgCreatePool instance.
func NewMsgWithdraw(from string, poolID uint64, amount sdk.Coins) *MsgWithdraw {
	return &MsgWithdraw{
		From:   from,
		PoolId: poolID,
		Amount: amount,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgWithdraw) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgWithdraw) Type() string { return TypeMsgDeposit }

// GetSigners implements the sdk.Msg interface.
func (msg MsgWithdraw) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgWithdraw) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return err
	}
	if from.Empty() {
		return ErrInvalidDepositor
	}

	if msg.PoolId == 0 {
		return ErrInvalidPoolID
	}

	// only accept one coin denom at a time
	if len(msg.Amount) > 1 || len(msg.Amount) == 0 {
		return ErrInvalidDepositCoins
	}
	return nil
}
