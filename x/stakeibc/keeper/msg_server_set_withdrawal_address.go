package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	"github.com/soohoio/stayking/utils"
	"github.com/soohoio/stayking/x/stakeibc/types"
	"github.com/spf13/cast"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetWithdrawalAddress(goCtx context.Context, msg *types.MsgSetWithdrawalAddress) (*types.MsgSetWithdrawalAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info(fmt.Sprintf("setWithdrawal Address by cmd: %s", msg.String()))
	for _, hostZone := range k.GetAllHostZone(ctx) {
		owner := types.FormatICAAccountOwner(hostZone.ChainId, types.ICAAccountType_DELEGATION)
		portID, err := icatypes.NewControllerPortID(owner)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "%s has no associated portId", owner)
		}
		connectionId, err := k.GetConnectionId(ctx, portID)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidChainID, "%s has no associated connection", portID)
		}
		delegationIca := hostZone.DelegationAccount
		if delegationIca == nil || delegationIca.Address == "" {
			k.Logger(ctx).Error(fmt.Sprintf("Zone %s is missing a delegation address!", hostZone.ChainId))
			return nil, nil
		}
		withdrawalIca := hostZone.WithdrawalAccount
		if withdrawalIca == nil || withdrawalIca.Address == "" {
			k.Logger(ctx).Error(fmt.Sprintf("Zone %s is missing a withdrawal address!", hostZone.ChainId))
			return nil, nil
		}
		withdrawalIcaAddr := hostZone.WithdrawalAccount.Address

		k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "Withdrawal Address: %s, Delegator Address: %s", withdrawalIcaAddr, delegationIca.Address))

		icaTimeOutNanos := k.GetParam(ctx, types.KeyICATimeoutNanos)
		icaTimeOutNanos = cast.ToUint64(ctx.BlockTime().UnixNano()) + icaTimeOutNanos
		msgs := []sdk.Msg{
			&distributiontypes.MsgSetWithdrawAddress{
				DelegatorAddress: delegationIca.Address,
				WithdrawAddress:  withdrawalIcaAddr,
			},
		}
		_, err = k.SubmitTxs(ctx, connectionId, msgs, *delegationIca, icaTimeOutNanos, "", nil)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Failed to SubmitTxs for %s, %s, %s", connectionId, hostZone.ChainId, msgs)
		}

		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Unable to set withdrawal address on %s, err: %s", hostZone.ChainId, err))
		}
	}

	// The relevant ICA is the delegate account

	return &types.MsgSetWithdrawalAddressResponse{}, nil
}
