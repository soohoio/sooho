package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAdjustPosition = "adjust_position"

var _ sdk.Msg = &MsgAdjustPosition{}

func NewMsgAdjustPosition(creator string, positionId uint64, collateral sdk.Int, debt sdk.Int, hostDenom string) *MsgAdjustPosition {
	return &MsgAdjustPosition{
		Creator:    creator,
		PositionId: positionId,
		Collateral: collateral,
		Debt:       debt,
		HostDenom:  hostDenom,
	}
}

func (msg *MsgAdjustPosition) Route() string {
	return RouterKey
}

func (msg *MsgAdjustPosition) Type() string {
	return TypeMsgAdjustPosition
}

func (msg *MsgAdjustPosition) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Collateral.LT(sdk.ZeroInt()) || msg.Debt.LT(sdk.ZeroInt()) {
		return errorsmod.Wrapf(ErrInvalidAmount, "either collateral or debt should be >= 0")
	}

	if msg.Collateral.Equal(sdk.ZeroInt()) && msg.Debt.Equal(sdk.ZeroInt()) {
		return errorsmod.Wrapf(ErrInvalidAmount, "both collateral and debt should not be zero")
	}

	if msg.HostDenom == "" {
		return errorsmod.Wrapf(ErrRequiredFieldEmpty, "host denom can not be empty")
	}

	return nil
}

func (msg *MsgAdjustPosition) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}
