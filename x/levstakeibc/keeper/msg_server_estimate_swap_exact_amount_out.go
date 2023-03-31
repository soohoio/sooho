package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func (k msgServer) SendEstimateSwapExactAmountOut(goCtx context.Context, msg *types.SendEstimateSwapExactAmountOutRequest) (*types.SendEstimateSwapExactAmountOutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info(fmt.Sprintf("[SendQuery DEBUG] Get chan cap: %v", host.ChannelCapabilityPath("icacontroller-osmosis-localnet.ICQ", msg.Sender)))
	k.Logger(ctx).Info(fmt.Sprintf("[SendQuery DEBUG] Get chan cap: %v", host.ChannelCapabilityPath("icacontroller.osmosis-localnet.ICQ", string(msg.PoolId))))

	k.Logger(ctx).Info(fmt.Sprintf("[SendQuery DEBUG] Get chan cap: %v", host.ChannelCapabilityPath("icacontroller-levstakeibc.osmosis-localnet.ICQ", msg.TokenOut)))
	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath("icacontroller-levstakeibc.osmosis-localnet.ICQ", "channel-9"))
	if !found {
		return nil, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	q := types.EstimateSwapExactAmountOutRequest{
		PoolId: msg.PoolId,
		Routes: msg.Routes,
		TokenOut: msg.TokenOut,
	}

	reqs := []abcitypes.RequestQuery{
		{
			Path: "/osmosis.poolmanager.v1beta1.Query/EstimateSinglePoolSwapExactAmountOut",
			Data: k.cdc.MustMarshal(&q),
		},
	}

	// timeoutTimestamp set to max value with the unsigned bit shifted to sastisfy hermes timestamp conversion
	// it is the responsibility of the auth module developer to ensure an appropriate timeout timestamp
	timeoutTimestamp := ctx.BlockTime().Add(time.Minute).UnixNano()
	seq, err := k.InterchainQueryKeeper.SendQuery(ctx, "icacontroller-levstakeibc.osmosis-localnet.ICQ", "channel-9", chanCap, reqs, clienttypes.ZeroHeight(), uint64(timeoutTimestamp))
	if err != nil {
		return nil, err
	}

	k.InterchainQueryKeeper.SetQuerySwapRequest(ctx, seq, q)

	return &types.SendEstimateSwapExactAmountOutResponse{
		Sequence: seq,
	}, nil
}
