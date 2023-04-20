package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/soohoio/stayking/v2/utils"
)

const TypeMsgEstimateSwapExactAmountOut = "send_estimate_swap_exact_amount_out"

var _ sdk.Msg = &MsgSendEstimateSwapExactAmountOutRequest{}

func NewMsgSendEstimateSwapExactAmountOut(creator string, poolId uint64, routes []SwapAmountOutRoute, tokenOut string, channelId string) *MsgSendEstimateSwapExactAmountOutRequest {
	return &MsgSendEstimateSwapExactAmountOutRequest{
		Creator:   creator,
		PoolId:    poolId,
		Routes:    routes,
		TokenOut:  tokenOut,
		ChannelId: channelId,
	}
}

func (msg *MsgSendEstimateSwapExactAmountOutRequest) Route() string {
	return RouterKey
}

func (msg *MsgSendEstimateSwapExactAmountOutRequest) Type() string {
	return TypeMsgEstimateSwapExactAmountOut
}

func (msg *MsgSendEstimateSwapExactAmountOutRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendEstimateSwapExactAmountOutRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendEstimateSwapExactAmountOutRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if err := utils.ValidateAdminAddress(msg.Creator); err != nil {
		return err
	}
	return nil
}
