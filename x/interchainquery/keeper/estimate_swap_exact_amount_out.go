package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	"github.com/soohoio/stayking/v2/x/interchainquery/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"strconv"
)

func (k Keeper) EstimateSwapExactAmountOut(ctx sdk.Context, epochNumber uint64) error {

	k.Logger(ctx).Info(fmt.Sprintf("[EstimateSwapExactAmountOut Query is triggered by Epoch]"))

	channelId := types.PriceQueryChannelId
	poolID, err := strconv.Atoi(types.PriceQueryPoolId)
	if err != nil {
		return err
	}
	routes := []types.SwapAmountOutRoute{}
	pID, err := strconv.Atoi(types.PriceQueryRoutesPoolId)
	if err != nil {
		return err
	}
	tokenInDenom := types.PriceQueryTokenInDenom
	tokenOut := types.PriceQueryTokenOut
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
	timeoutTimestamp := ctx.BlockTime().Add(180000000000).UnixNano()
	seq, err := k.SendQuery(ctx, "interchainquery", channelId, chanCap, reqs, clienttypes.ZeroHeight(), uint64(timeoutTimestamp))
	if err != nil {
		return err
	}

	k.SetQuerySwapRequest(ctx, seq, q)

	return nil
}
