package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func (k msgServer) ExitLeverageStake(goCtx context.Context, msg *types.MsgExitLeverageStake) (*types.MsgExitLeverageStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger(ctx).Info(fmt.Sprintf("Exit Leverage stake: %s", msg.String()))

	err := k.UnStakeWithLeverage(ctx, msg.GetCreator(), msg.GetPositionId(), msg.GetChainId(), msg.GetReceiver())

	if err != nil {
		return nil, err
	}

	k.Logger(ctx).Info(fmt.Sprintf("executed Exit Leverage stake: %s", msg.String()))

	return &types.MsgExitLeverageStakeResponse{}, nil

}