package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

const TypeMsgQuerySendEstimateSwapExactAmountOut = "query_estimate_swap_exact_amount_out"

var _ sdk.Msg = &SendEstimateSwapExactAmountOutRequest{}

func NewSendEstimateSwapExactAmountOut(creator string, poolId uint64, routes []SwapAmountOutRoute, tokenOut string,page *query.PageRequest) *SendEstimateSwapExactAmountOutRequest {
	return &SendEstimateSwapExactAmountOutRequest{
		Sender:	 creator,
		PoolId:   poolId,
		Routes:   routes,
		TokenOut: tokenOut,
		Pagination: page,
	}
}

func (msg *SendEstimateSwapExactAmountOutRequest) Route() string {
	return RouterKey
}

func (msg *SendEstimateSwapExactAmountOutRequest) Type() string {
	return TypeMsgQuerySendEstimateSwapExactAmountOut
}

func (msg *SendEstimateSwapExactAmountOutRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *SendEstimateSwapExactAmountOutRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *SendEstimateSwapExactAmountOutRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
