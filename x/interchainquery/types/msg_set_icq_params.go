package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSSetIcqParams = "set_icq_params"

var _ sdk.Msg = &MsgSetIcqParams{}

func NewMsgSetIcqParams(creator string, channelId string, poolId string, routesPoolId string, tokenInDenom string, tokenOut string, path string) *MsgSetIcqParams {
	return &MsgSetIcqParams{
		Creator:                creator,
		PriceQueryChannelId:    channelId,
		PriceQueryPoolId:       poolId,
		PriceQueryRoutesPoolId: routesPoolId,
		PriceQueryTokenInDenom: tokenInDenom,
		PriceQueryTokenOut:     tokenOut,
		PriceQueryPath:         path,
	}
}

func (msg *MsgSetIcqParams) Route() string {
	return RouterKey
}

func (msg *MsgSetIcqParams) Type() string {
	return TypeMsgSSetIcqParams
}

func (msg *MsgSetIcqParams) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetIcqParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetIcqParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
