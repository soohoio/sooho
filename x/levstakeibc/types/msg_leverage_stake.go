package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLeverageStake = "leverage_stake"

var _ sdk.Msg = &MsgLeverageStake{}

func NewMsgLeverageStake(creator string, equity sdk.Int, hostDenom string, leverageRatio sdk.Dec) *MsgLeverageStake {
	return &MsgLeverageStake{
		Creator:       creator,
		HostDenom:     hostDenom,
		Equity:        equity,
		LeverageRatio: leverageRatio,
	}
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
		return errorsmod.Wrapf()
	}

	return nil
}

func (msg *MsgLeverageStake) GetSigners() []sdk.AccAddress {
	//TODO implement me
	panic("implement me")
}
