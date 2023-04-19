package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeExitLeverageStake = "exit_leverage_stake"

var _ sdk.Msg = &MsgExitLeverageStake{}

func NewMsgExitLeverageStake(creator string, positionId uint64, chainId string, receiver string) *MsgExitLeverageStake {
	return &MsgExitLeverageStake{
		Creator:    creator,
		PositionId: positionId,
		ChainId:    chainId,
		Receiver:   receiver,
	}
}

func (m *MsgExitLeverageStake) Route() string {
	return RouterKey
}

func (m *MsgExitLeverageStake) Type() string {
	return TypeExitLeverageStake
}

func (m *MsgExitLeverageStake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (m *MsgExitLeverageStake) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if m.ChainId == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid request Position Id should not be empty")
	}
	if m.Receiver == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid request Position Id should not be empty")
	}
	return nil
}
