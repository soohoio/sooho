package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	admintypes "github.com/soohoio/stayking/v2/x/admin/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ibctransfertypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	"github.com/spf13/cast"

	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func (k msgServer) ClearBalance(goCtx context.Context, msg *types.MsgClearBalance) (*types.MsgClearBalanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// admin address check
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if !k.AdminKeeper.IsAdmin(ctx, creator) {
		return nil, admintypes.ErrNotAdmin
	}
	zone, found := k.GetHostZone(ctx, msg.ChainId)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrHostZoneNotFound, "host zone not found by chain id %s", msg.ChainId)
	}
	feeAccount := zone.GetFeeAccount()
	if feeAccount == nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidAccount, "fee account not found by chainId %v", msg.ChainId)
	}

	sourcePort := ibctransfertypes.PortID
	// Should this be a param?
	// I think as long as we have a timeout on this, it should be hard to attack (even if someone send a tx on a bad channel, it would be reverted relatively quickly)
	sourceChannel := msg.Channel
	coinString := cast.ToString(msg.Amount) + zone.GetHostDenom()
	tokens, err := sdk.ParseCoinNormalized(coinString)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("failed to parse coin (%s)", coinString))
		return nil, sdkerrors.Wrapf(err, "failed to parse coin (%s)", coinString)
	}
	sender := feeAccount.GetAddress()
	// KeyICATimeoutNanos are for our StayKing ICA calls, KeyFeeTransferTimeoutNanos is for the IBC transfer
	feeTransferTimeoutNanos := k.GetParam(ctx, types.KeyFeeTransferTimeoutNanos)
	timeoutTimestamp := cast.ToUint64(ctx.BlockTime().UnixNano()) + feeTransferTimeoutNanos
	msgs := []sdk.Msg{
		&ibctransfertypes.MsgTransfer{
			SourcePort:       sourcePort,
			SourceChannel:    sourceChannel,
			Token:            tokens,
			Sender:           sender,
			Receiver:         types.StakingFeeAccount,
			TimeoutTimestamp: timeoutTimestamp,
		},
	}

	connectionId := zone.GetConnectionId()

	icaTimeoutNanos := k.GetParam(ctx, types.KeyICATimeoutNanos)
	icaTimeoutNanos = cast.ToUint64(ctx.BlockTime().UnixNano()) + icaTimeoutNanos

	_, err = k.SubmitTxs(ctx, connectionId, msgs, *feeAccount, icaTimeoutNanos, "", nil)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to submit txs")
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeClearBalance,
			sdk.NewAttribute(types.AttributeKeyRecipientChain, msg.ChainId),
			sdk.NewAttribute(types.AttributeKeyConnectionId, connectionId),
			sdk.NewAttribute(types.AttributeKeyChannelId, sourceChannel),
			sdk.NewAttribute(types.AttributeKeyFromAddress, sender),
			sdk.NewAttribute(types.AttributeKeyToAddress, types.StakingFeeAccount),
			sdk.NewAttribute(types.AttributeKeyTransferTokenAmount, tokens.String()),
		),
	)
	return &types.MsgClearBalanceResponse{}, nil
}
