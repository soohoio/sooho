package keeper

import (
	"context"
	"fmt"

	"github.com/soohoio/stayking/x/stakeibc/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetWithdrawalAddress(goCtx context.Context, msg *types.MsgSetWithdrawalAddress) (*types.MsgSetWithdrawalAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info(fmt.Sprintf("setWithdrawal Address: "))
	//k.Logger(ctx).Info(fmt.Sprintf("setWithdrawal Address: %s", msg.String()))

	return &types.MsgSetWithdrawalAddressResponse{}, nil
}
