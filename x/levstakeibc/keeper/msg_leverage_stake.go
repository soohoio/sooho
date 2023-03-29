package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func (k msgServer) LeverageStake(goCtx context.Context, msg *types.MsgLeverageStake) (*types.MsgLeverageStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	leverageRatio := msg.LeverageRatio

	msg.GetStakeType(leverageRatio)
}
