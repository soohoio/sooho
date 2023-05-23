package types

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
)

const (
	TypeMsgCreatePool = "create_pool"
	TypeMsgDeletePool = "delete_pool"
	TypeMsgUpdatePool = "update_pool"
	TypeMsgDeposit    = "deposit"
	TypeMsgWithdraw   = "withdraw"
	TypeMsgLiquidate  = "liquidate"
)

var (
	_ sdk.Msg = &MsgCreatePool{}
	_ sdk.Msg = &MsgDeletePool{}
	_ sdk.Msg = &MsgUpdatePool{}
	_ sdk.Msg = &MsgDeposit{}
	_ sdk.Msg = &MsgWithdraw{}
	_ sdk.Msg = &MsgLiquidate{}
)

// NewMsgCreatePool creates a new NewMsgCreatePool instance.
func NewMsgCreatePool(creator, denom string, maxDebtRatio sdk.Dec, interestModel InterestModelI) (*MsgCreatePool, error) {
	msg := &MsgCreatePool{
		Creator:      creator,
		Denom:        denom,
		MaxDebtRatio: maxDebtRatio,
	}
	err := msg.SetInterestModel(interestModel)
	return msg, err
}

// NewMsgDeletePool deletes a new NewMsgDeletePool instance.
func NewMsgDeletePool(creator string, poolId uint64) *MsgDeletePool {
	return &MsgDeletePool{
		Creator: creator,
		PoolId:  poolId,
	}
}

// NewMsgUpdatePool updates a new NewMsgUpdatePool instance.
func NewMsgUpdatePool(creator string, poolId uint64, denom string, maxDebtRatio sdk.Dec, interestModel InterestModelI) (*MsgUpdatePool, error) {
	msg := &MsgUpdatePool{
		Creator:      creator,
		PoolId:       poolId,
		Denom:        denom,
		MaxDebtRatio: maxDebtRatio,
	}
	err := msg.SetInterestModel(interestModel)
	return msg, err
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

func (m MsgCreatePool) GetInterestModel() InterestModelI {
	model, ok := m.InterestModel.GetCachedValue().(InterestModelI)
	if !ok {
		return nil
	}
	return model
}

func (m *MsgCreatePool) SetInterestModel(model InterestModelI) error {
	msg, ok := model.(proto.Message)
	if !ok {
		return fmt.Errorf("can't proto marshal %T", msg)
	}
	a, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return err
	}
	m.InterestModel = a
	return nil
}

// Route implements the sdk.Msg interface.
func (msg MsgDeletePool) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgDeletePool) Type() string { return TypeMsgDeletePool }

// GetSigners implements the sdk.Msg interface.
func (msg MsgDeletePool) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgDeletePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// Route implements the sdk.Msg interface.
func (msg MsgUpdatePool) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgUpdatePool) Type() string { return TypeMsgUpdatePool }

// GetSigners implements the sdk.Msg interface.
func (msg MsgUpdatePool) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgUpdatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}
func (m MsgUpdatePool) GetInterestModel() InterestModelI {
	model, ok := m.InterestModel.GetCachedValue().(InterestModelI)
	if !ok {
		return nil
	}
	return model
}

func (m *MsgUpdatePool) SetInterestModel(model InterestModelI) error {
	msg, ok := model.(proto.Message)
	if !ok {
		return fmt.Errorf("can't proto marshal %T", msg)
	}
	a, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return err
	}
	m.InterestModel = a
	return nil
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgCreatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Denom == "" || sdk.ValidateDenom(msg.Denom) != nil {
		return ErrInvalidDenom
	}

	if err = msg.GetInterestModel().ValidateBasic(); err != nil {
		return err
	}

	return nil
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgDeletePool) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}
	if from.Empty() {
		return ErrEmptyCreator
	}

	if msg.PoolId == 0 {
		return ErrInvalidPoolID
	}
	return nil
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgUpdatePool) ValidateBasic() error {
	from, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}
	if msg.PoolId == 0 {
		return ErrInvalidPoolID
	}
	if from.Empty() {
		return ErrEmptyCreator
	}

	if msg.Denom == "" || sdk.ValidateDenom(msg.Denom) != nil {
		return ErrInvalidDenom
	}
	if err = msg.GetInterestModel().ValidateBasic(); err != nil {
		return err
	}
	return nil
}

func (m MsgCreatePool) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var model InterestModelI
	return unpacker.UnpackAny(m.InterestModel, &model)
}

func (m MsgUpdatePool) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var model InterestModelI
	return unpacker.UnpackAny(m.InterestModel, &model)
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
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.PoolId == 0 {
		return ErrInvalidPoolID
	}

	// only accept one coin denom at a time
	if len(msg.Amount) != 1 {
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
func (msg MsgWithdraw) Type() string { return TypeMsgWithdraw }

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
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.PoolId == 0 {
		return ErrInvalidPoolID
	}

	// only accept one coin denom at a time
	if len(msg.Amount) != 1 {
		return ErrInvalidWithdrawCoins
	}
	return nil
}

// NewMsgLiquidate creates a new NewMsgLiquidate instance.
func NewMsgLiquidate(from string, loanID uint64) *MsgLiquidate {
	return &MsgLiquidate{
		From:   from,
		LoanId: loanID,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgLiquidate) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgLiquidate) Type() string { return TypeMsgLiquidate }

// GetSigners implements the sdk.Msg interface.
func (msg MsgLiquidate) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgLiquidate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgLiquidate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.LoanId == 0 {
		return errorsmod.Wrapf(ErrInvalidLoanId, "Loan Id can not equal to 0")
	}

	return nil
}
