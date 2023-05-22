package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimUndelegatedTokens = "claim_undelegated_tokens"

var _ sdk.Msg = &MsgClaimUndelegatedTokens{}

func NewMsgClaimUndelegatedTokens(creator string, hostZoneId string, epoch uint64, sender string) *MsgClaimUndelegatedTokens {
	return &MsgClaimUndelegatedTokens{
		Creator:    creator,
		HostZoneId: hostZoneId,
		Epoch:      epoch,
		Sender:     sender,
	}
}

func (m *MsgClaimUndelegatedTokens) Route() string {
	return RouterKey
}

func (m *MsgClaimUndelegatedTokens) Type() string {
	return TypeMsgClaimUndelegatedTokens
}

func (m *MsgClaimUndelegatedTokens) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (m *MsgClaimUndelegatedTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.GetCreator())
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(m.GetSender())
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if m.GetHostZoneId() == "" {
		return errorsmod.Wrapf(ErrRequiredFieldEmpty, "chainId cannot be empty")
	}

	if m.GetEpoch() == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid Epoch Number: %d", m.GetEpoch())
	}

	return nil
}
