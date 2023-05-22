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
	"github.com/soohoio/stayking/v3/x/interchainquery/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func (k Keeper) EstimateSwapExactAmountOut(ctx sdk.Context, epochNumber uint64) error {

	k.Logger(ctx).Info(fmt.Sprintf("[EstimateSwapExactAmountOut Query is triggered by Epoch]"))
	params := k.GetParams(ctx)
	k.Logger(ctx).Info(fmt.Sprintf("Estimate params channelid:%v Poolid:%v RoutespoolId:%v TokenInDenom:%v TokenOut:%v ", params.PriceQueryChannelId, params.PriceQueryPoolId, params.PriceQueryRoutesPoolId, params.PriceQueryTokenInDenom, params.PriceQueryTokenOut))
	channelId := params.PriceQueryChannelId
	poolID, err := strconv.Atoi(params.PriceQueryPoolId)
	if err != nil {
		return err
	}
	routes := []types.SwapAmountOutRoute{}
	pID, err := strconv.Atoi(params.PriceQueryRoutesPoolId)
	if err != nil {
		return err
	}
	tokenInDenom := params.PriceQueryTokenInDenom
	tokenOut := params.PriceQueryTokenOut
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
			Path: params.PriceQueryPath,
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
