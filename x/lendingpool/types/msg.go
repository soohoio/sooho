package types

import (
	"fmt"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
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
func NewMsgCreatePool(creator, denom string, interestModel InterestModelI) (*MsgCreatePool, error) {
	msg := &MsgCreatePool{
		Creator: creator,
		Denom:   denom,
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

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgCreatePool) ValidateBasic() error {
	fmt.Println("model: ", msg.GetInterestModel())
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
	fmt.Println("model: ", msg.GetInterestModel())
	if err := msg.GetInterestModel().ValidateBasic(); err != nil {
		return err
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
