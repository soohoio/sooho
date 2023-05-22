package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v3/x/levstakeibc/types"
)

func (k msgServer) ExitLeverageStake(goCtx context.Context, msg *types.MsgExitLeverageStake) (*types.MsgExitLeverageStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger(ctx).Info(fmt.Sprintf("Exit Leverage stake: %s", msg.String()))

	err := k.UnStakeWithLeverage(ctx, msg.GetCreator(), msg.GetPositionId(), msg.GetChainId(), msg.GetReceiver())

	if err != nil {
		k.Logger(ctx).Error("[CUSTOM DEBUG] ExitLeverageStake error reason : " + err.Error())
		return nil, err
	}

	k.Logger(ctx).Info(fmt.Sprintf("executed Exit Leverage stake: %s", msg.String()))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)
	return &types.MsgExitLeverageStakeResponse{}, nil

}
