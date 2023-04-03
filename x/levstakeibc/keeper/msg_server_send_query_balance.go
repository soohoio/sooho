package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func (k msgServer) SendQueryBalance(goCtx context.Context, msg *types.MsgSendQueryBalance) (*types.MsgSendQueryBalanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info(fmt.Sprintf("[SendQuery DEBUG] Get chan cap: %v", host.ChannelCapabilityPath("icacontroller-osmosis-localnet.ICQ", msg.ChannelId)))
	k.Logger(ctx).Info(fmt.Sprintf("[SendQuery DEBUG] Get chan cap: %v", host.ChannelCapabilityPath("icacontroller.osmosis-localnet.ICQ", msg.ChannelId)))
	k.Logger(ctx).Info(fmt.Sprintf("[SendQuery DEBUG] Get chan cap: %v", host.ChannelCapabilityPath("icacontroller-levstakeibc.osmosis-localnet.ICQ", msg.ChannelId)))
	k.Logger(ctx).Info(fmt.Sprintf("[SendQuery DEBUG] Get chan cap: %v", host.ChannelCapabilityPath("icacontroller-levstakeibc.osmosis-localnet.ICQ", msg.ChannelId)))
	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath("icacontroller-levstakeibc.osmosis-localnet.ICQ", msg.ChannelId))
	if !found {
		return nil, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	q := banktypes.QueryBalanceRequest{
		Address:    msg.Address,
		Denom: 		msg.Denom,
	}
	reqs := []abcitypes.RequestQuery{
		{
			Path: "/cosmos.bank.v1beta1.Query/Balance",
			Data: k.cdc.MustMarshal(&q),
		},
	}

	// timeoutTimestamp set to max value with the unsigned bit shifted to sastisfy hermes timestamp conversion
	// it is the responsibility of the auth module developer to ensure an appropriate timeout timestamp
	timeoutTimestamp := ctx.BlockTime().Add(time.Minute).UnixNano()
	seq, err := k.SendQuery(ctx, "icacontroller-levstakeibc.osmosis-localnet.ICQ", msg.ChannelId, chanCap, reqs, clienttypes.ZeroHeight(), uint64(timeoutTimestamp))
	if err != nil {
		return nil, err
	}

	k.SetQueryBalanceRequest(ctx, seq, q)

	return &types.MsgSendQueryBalanceResponse{
		Sequence: seq,
	}, nil
}
