package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRestoreInterchainAccount = "restore_interchain_account"

var _ sdk.Msg = &MsgRestoreInterchainAccount{}

func NewMsgRestoreInterchainAccount(creator string, chainId string, accountType ICAType) *MsgRestoreInterchainAccount {
	return &MsgRestoreInterchainAccount{
		Creator:     creator,
		ChainId:     chainId,
		AccountType: accountType,
	}
}

func (msg *MsgRestoreInterchainAccount) Route() string {
	return RouterKey
}

func (msg *MsgRestoreInterchainAccount) Type() string {
	return TypeMsgRestoreInterchainAccount
}

func (msg *MsgRestoreInterchainAccount) ValidateBasic() error {

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.ChainId == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "chain id cannot be empty")
	}

	if msg.GetAccountType() < 0 || msg.GetAccountType() > 2 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid account type (%v)", msg.GetAccountType())
	}

	return nil
}

func (msg *MsgRestoreInterchainAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRestoreInterchainAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}
