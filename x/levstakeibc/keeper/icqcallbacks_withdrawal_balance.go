package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/soohoio/stayking/v2/utils"
	epochstypes "github.com/soohoio/stayking/v2/x/epochs/types"
	icqkeeper "github.com/soohoio/stayking/v2/x/interchainquery/keeper"
	icqtypes "github.com/soohoio/stayking/v2/x/interchainquery/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"github.com/spf13/cast"
)

// WithdrawalBalanceCallback is a callback handler for WithdrawalBalance queries.
// The query response will return the withdrawal account balance
// If the balance is non-zero, ICA MsgSends are submitted to transfer from the withdrawal account
//
//	to the delegation account (for reinvestment) and fee account (for commission)
//
// Note: for now, to get proofs in your ICQs, you need to query the entire store on the host zone! e.g. "store/bank/key"
func WithdrawalBalanceCallback(k Keeper, ctx sdk.Context, args []byte, query icqtypes.Query) error {
	k.Logger(ctx).Info(utils.LogICQCallbackWithHostZone(query.ChainId, ICQCallbackID_WithdrawalBalance,
		"[ICQ RESPONSE] Starting withdrawal balance callback, QueryId: %vs, QueryType: %s, Connection: %s", query.Id, query.QueryType, query.ConnectionId))

	// Confirm host exists
	chainId := query.ChainId
	hostZone, found := k.GetHostZone(ctx, chainId)
	if !found {
		return errorsmod.Wrapf(types.ErrHostZoneNotFound, "not found a host zone by chain id %s", chainId)
	}

	// Unmarshal the query response args to determine the balance
	withdrawalBalanceAmount, err := icqkeeper.UnmarshalAmountFromBalanceQuery(k.cdc, args)
	if err != nil {
		return errorsmod.Wrap(err, "unable to determine balance from query response")
	}
	k.Logger(ctx).Info(utils.LogICQCallbackWithHostZone(chainId, ICQCallbackID_WithdrawalBalance,
		"[ICQ RESPONSE] Query response - Withdrawal Balance: %v %s", withdrawalBalanceAmount, hostZone.HostDenom))

	// Confirm the balance is greater than zero
	if withdrawalBalanceAmount.LTE(sdkmath.ZeroInt()) {
		k.Logger(ctx).Info(utils.LogICQCallbackWithHostZone(chainId, ICQCallbackID_WithdrawalBalance,
			"[ICQ RESPONSE] No balance to transfer for address: %v, balance: %v", hostZone.WithdrawalAccount.GetAddress(), withdrawalBalanceAmount))
		return nil
	}

	// Get the host zone's ICA accounts
	withdrawalAccount := hostZone.WithdrawalAccount
	if withdrawalAccount == nil || withdrawalAccount.Address == "" {
		return errorsmod.Wrapf(types.ErrHostZoneICAAccountNotFound, "no withdrawal account found for %s", chainId)
	}
	delegationAccount := hostZone.DelegationAccount
	if delegationAccount == nil || delegationAccount.Address == "" {
		return errorsmod.Wrapf(types.ErrHostZoneICAAccountNotFound, "no delegation account found for %s", chainId)
	}
	feeAccount := hostZone.FeeAccount
	if feeAccount == nil || feeAccount.Address == "" {
		return errorsmod.Wrapf(types.ErrHostZoneICAAccountNotFound, "no fee account found for %s", chainId)
	}

	// Determine the stayking commission rate to the relevant portion can be sent to the fee account
	params := k.GetParams(ctx)
	staykingCommissionInt, err := cast.ToInt64E(params.StaykingCommission)
	if err != nil {
		return err
	}

	// check that stayking commission is between 0 and 1
	staykingCommission := sdk.NewDec(staykingCommissionInt).Quo(sdk.NewDec(100))
	if staykingCommission.LT(sdk.ZeroDec()) || staykingCommission.GT(sdk.OneDec()) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Aborting withdrawal balance callback -- Stayking commission must be between 0 and 1!")
	}

	feeAmount := staykingCommission.Mul(sdk.NewDecFromInt(withdrawalBalanceAmount)).TruncateInt()
	reinvestAmount := withdrawalBalanceAmount.Sub(feeAmount)

	if !feeAmount.Add(reinvestAmount).Equal(withdrawalBalanceAmount) {
		k.Logger(ctx).Error(fmt.Sprintf("Error with withdraw logic: %v, Fee Portion: %v, Reinvest Portion %v", withdrawalBalanceAmount, feeAmount, reinvestAmount))
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Failed to subdivide rewards to commission and delegationAccount")
	}

	feeCoin := sdk.NewCoin(hostZone.HostDenom, feeAmount)
	reinvestCoin := sdk.NewCoin(hostZone.HostDenom, reinvestAmount)

	var msgs []sdk.Msg
	if feeCoin.Amount.GT(sdk.ZeroInt()) {
		msgs = append(msgs, &banktypes.MsgSend{
			FromAddress: withdrawalAccount.Address,
			ToAddress:   feeAccount.Address,
			Amount:      sdk.NewCoins(feeCoin),
		})
		k.Logger(ctx).Info(utils.LogICQCallbackWithHostZone(chainId, ICQCallbackID_WithdrawalBalance,
			"[ICQ RESPONSE] Preparing MsgSends of %v from the withdrawal account to the fee account (for commission)", feeCoin.String()))
	}
	if reinvestCoin.Amount.GT(sdk.ZeroInt()) {
		msgs = append(msgs, &banktypes.MsgSend{
			FromAddress: withdrawalAccount.Address,
			ToAddress:   delegationAccount.Address,
			Amount:      sdk.NewCoins(reinvestCoin),
		})
		k.Logger(ctx).Info(utils.LogICQCallbackWithHostZone(chainId, ICQCallbackID_WithdrawalBalance,
			"[ICQ RESPONSE] Preparing MsgSends of %v from the withdrawal account to the delegation account (for reinvestment)", reinvestCoin.String()))
	}

	// add callback data before calling reinvestment ICA
	reinvestCallback := types.ReinvestCallback{
		ReinvestAmount: reinvestCoin,
		HostZoneId:     hostZone.ChainId,
	}
	k.Logger(ctx).Info(utils.LogICQCallbackWithHostZone(chainId, ICQCallbackID_WithdrawalBalance, "Marshalling ReinvestCallback args: %v", reinvestCallback))
	marshalledCallbackArgs, err := k.MarshalReinvestCallbackArgs(ctx, reinvestCallback)
	if err != nil {
		return err
	}

	// Send the transaction through SubmitTx
	_, err = k.SubmitTxsEpoch(ctx, hostZone.ConnectionId, msgs, *withdrawalAccount, epochstypes.STAYKING_EPOCH, ICACallbackID_Reinvest, marshalledCallbackArgs)
	if err != nil {
		return errorsmod.Wrapf(types.ErrICATxFailed, "Failed to SubmitTxs, Messages: %v, err: %s", msgs, err.Error())
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute("hostZone", hostZone.ChainId),
			sdk.NewAttribute("totalWithdrawalBalance", withdrawalBalanceAmount.String()),
		),
	)

	return nil
}
