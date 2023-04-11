package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	"github.com/soohoio/stayking/v2/x/interchainquery/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func (k Keeper) EstimateSwapExactAmountOut(ctx sdk.Context, epochNumber uint64) error {

	k.Logger(ctx).Info(fmt.Sprintf("[EstimateSwapExactAmountOut Query is triggered by Epoch]"))

	//@TODO remove hard coded variables, it must be able to be controlled by parameters
	channelId := "channel-1"
	poolID, err := strconv.Atoi("1")
	if err != nil {
		return err
	}
	routes := []types.SwapAmountOutRoute{}
	//@TODO multihop poolId
	pID, err := strconv.Atoi("1")
	if err != nil {
		return err
	}
	//@TODO it must be EVMOS
	tokenInDenom := "uosmo"
	tokenOut := "1uosmo"
	routes = append(routes, types.SwapAmountOutRoute{
		PoolId:       uint64(pID),
		TokenInDenom: tokenInDenom,
	})
	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath("interchainquery", channelId))
	if !found {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	q := types.EstimateSwapExactAmountOutRequest{
		PoolId:   uint64(poolID),
		Routes:   routes,
		TokenOut: tokenOut,
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
	seq, err := k.SendQuery(ctx, "interchainquery", channelId, chanCap, reqs, clienttypes.ZeroHeight(), uint64(timeoutTimestamp))
	if err != nil {
		return err
	}

	k.SetQuerySwapRequest(ctx, seq, q)

	return nil
}
