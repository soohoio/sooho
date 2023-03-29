package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLeverageStake = "leverage_stake"

var _ sdk.Msg = &MsgLeverageStake{}

const (
	NotLeverageType = iota
	LeverageType    = iota + 1
)

func NewMsgLeverageStake(creator string, equity sdk.Int, hostDenom string, leverageRatio sdk.Dec) *MsgLeverageStake {
	return &MsgLeverageStake{
		Creator:       creator,
		HostDenom:     hostDenom,
		Equity:        equity,
		LeverageRatio: leverageRatio,
	}
}

func (msg *MsgLeverageStake) GetStakeType(leverageRatio sdk.Dec) int {
	if leverageRatio.GT(sdk.NewDec(1)) {
		return NotLeverageType
	} else if leverageRatio.Equal(sdk.NewDec(1)) {
		return LeverageType
	}

	panic("-")
}

func (msg *MsgLeverageStake) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Equity.LTE(sdk.ZeroInt()) {
		return errorsmod.Wrapf(ErrInvalidAmount, "amount leverage staked must be positive and nonzero")
	}

	if msg.HostDenom == "" {
		return errorsmod.Wrapf(ErrRequiredFieldEmpty, "host denom can not be empty")
	}

	if msg.LeverageRatio.LT(sdk.NewDec(1)) {
		return errorsmod.Wrapf(ErrLeverageRatio, "leverage ratio must be greater than 1.0 and equal 1.0")
	}

	return nil
}

func (msg *MsgLeverageStake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}
