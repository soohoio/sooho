package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	admintypes "github.com/soohoio/stayking/v3/x/admin/types"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	"github.com/soohoio/stayking/v3/x/interchainquery/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func (k msgServer) SendEstimateSwapExactAmountOut(goCtx context.Context, msg *types.MsgSendEstimateSwapExactAmountOutRequest) (*types.MsgSendEstimateSwapExactAmountOutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info(fmt.Sprintf("[Send EstimateSwapExactAmountOut Query DEBUG] Get chan cap: %v", host.ChannelCapabilityPath("interchainquery", msg.ChannelId)))
	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath("interchainquery", msg.ChannelId))
	if !found {
		return nil, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	// admin address check
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if !k.AdminKeeper.IsAdmin(ctx, creator) {
		return nil, admintypes.ErrNotAdmin
	}

	q := types.EstimateSwapExactAmountOutRequest{
		PoolId:   msg.PoolId,
		Routes:   msg.Routes,
		TokenOut: msg.TokenOut,
	}
	reqs := []abcitypes.RequestQuery{
		{
			Path: "/osmosis.poolmanager.v1beta1.Query/EstimateSwapExactAmountOut",
			Data: k.cdc.MustMarshal(&q),
		},
	}

	// timeoutTimestamp set to max value with the unsigned bit shifted to sastisfy hermes timestamp conversion
	// it is the responsibility of the auth module developer to ensure an appropriate timeout timestamp
	timeoutTimestamp := ctx.BlockTime().Add(time.Minute).UnixNano()
	seq, err := k.SendQuery(ctx, "interchainquery", msg.ChannelId, chanCap, reqs, clienttypes.ZeroHeight(), uint64(timeoutTimestamp))
	if err != nil {
		return nil, err
	}

	k.SetQuerySwapRequest(ctx, seq, q)

	return &types.MsgSendEstimateSwapExactAmountOutResponse{
		Sequence: seq,
	}, nil
}
