package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	"github.com/golang/protobuf/proto"
	"github.com/soohoio/stayking/v2/utils"
	icacallbackstypes "github.com/soohoio/stayking/v2/x/icacallbacks/types"
	lendingpooltypes "github.com/soohoio/stayking/v2/x/lendingpool/types"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	//recordstypes "github.com/soohoio/stayking/v2/x/records/types"
)

// Marshalls delegate callback arguments
func (k Keeper) MarshalTransferUndelegatedTokensArgs(ctx sdk.Context, TransferUndelegatedTokensCallback types.TransferUndelegatedTokensCallback) ([]byte, error) {
	out, err := proto.Marshal(&TransferUndelegatedTokensCallback)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("MarshalTransferUndelegatedTokensCallbackArgs %v", err.Error()))
		return nil, err
	}
	return out, nil
}

// Unmarshalls delegate callback arguments into a DelegateCallback struct
func (k Keeper) UnmarshalTransferUndelegatedTokensCallbackArgs(ctx sdk.Context, transferUndelegatedTokensCallback []byte) (*types.TransferUndelegatedTokensCallback, error) {
	unmarshalledTransferUndelegatedTokensCallback := types.TransferUndelegatedTokensCallback{}
	if err := proto.Unmarshal(transferUndelegatedTokensCallback, &unmarshalledTransferUndelegatedTokensCallback); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("UnmarshalTransferUndelegatedTokensCallbackArgs %v", err.Error()))
		return nil, err
	}
	return &unmarshalledTransferUndelegatedTokensCallback, nil
}

func TransferUndelegatedTokensCallback(k Keeper, ctx sdk.Context, packet channeltypes.Packet, ackResponse *icacallbackstypes.AcknowledgementResponse, args []byte) error {
	// Deserialize the callback args
	transferUndelegatedTokensCallback, err := k.UnmarshalTransferUndelegatedTokensCallbackArgs(ctx, args)
	if err != nil {
		return errorsmod.Wrapf(types.ErrUnmarshalFailure, fmt.Sprintf("Unable to unmarshal transferUndelegatedTokens callback args: %s", err.Error()))
	}
	chainId := transferUndelegatedTokensCallback.HostZoneId
	k.Logger(ctx).Info(utils.LogICACallbackWithHostZone(chainId, ICACallbackID_TransferUndelegatedTokens,
		"Starting TransferUndelegatedTokens callback for Position Record: %d", transferUndelegatedTokensCallback.PositionId))

	// Confirm chainId and deposit record Id exist
	hostZone, found := k.GetHostZone(ctx, chainId)
	if !found {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "host zone not found %s", chainId)
	}
	positionId := transferUndelegatedTokensCallback.PositionId
	position, found := k.GetPosition(ctx, positionId)
	if !found {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "position not found %d", positionId)
	}

	// Check for timeout (ack nil)
	if ackResponse.Status == icacallbackstypes.AckResponseStatus_TIMEOUT {
		k.Logger(ctx).Error(utils.LogICACallbackStatusWithHostZone(chainId, ICACallbackID_TransferUndelegatedTokens,
			icacallbackstypes.AckResponseStatus_TIMEOUT, packet))

		// Reset deposit record status
		//depositRecord.Status = recordstypes.DepositRecord_DELEGATION_QUEUE
		//k.RecordsKeeper.SetDepositRecord(ctx, depositRecord)
		return nil
	}

	// Check for a failed transaction (ack error)
	// Reset the deposit record status upon failure
	if ackResponse.Status == icacallbackstypes.AckResponseStatus_FAILURE {
		k.Logger(ctx).Error(utils.LogICACallbackStatusWithHostZone(chainId, ICACallbackID_Delegate,
			icacallbackstypes.AckResponseStatus_FAILURE, packet))

		// Reset deposit record status
		//depositRecord.Status = recordstypes.DepositRecord_DELEGATION_QUEUE
		//k.RecordsKeeper.SetDepositRecord(ctx, depositRecord)
		return nil
	}

	k.Logger(ctx).Info(utils.LogICACallbackStatusWithHostZone(chainId, ICACallbackID_Delegate,
		icacallbackstypes.AckResponseStatus_SUCCESS, packet))

	hostZoneAddress, err := sdk.AccAddressFromBech32(hostZone.Address)
	nativeCoin := sdk.Coins{sdk.Coin{Denom: position.Denom, Amount: position.NativeTokenAmount}}
	// Transfer undelegated tokens to module account
	lendingpoolAddress := k.accountKeeper.GetModuleAddress(lendingpooltypes.ModuleName)
	k.Logger(ctx).Info(fmt.Sprintf("[TransferUndeleagtedTokens for leveraged position] LendingPool Address %v", lendingpoolAddress))
	k.bankKeeper.SendCoinsFromAccountToModule(ctx, hostZoneAddress, lendingpooltypes.ModuleName, nativeCoin)
	_, err = k.LendingPoolKeeper.Repay(ctx, position.LoanId, sdk.NewCoins(sdk.NewCoin(hostZone.GetIbcDenom(), position.NativeTokenAmount)))
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Repay failed for loan id %v", position.LoanId))
	}
	k.RemovePosition(ctx, position.Id)
	// 처리 완료 후 record 삭제 -> staking 완료
	k.Logger(ctx).Info(fmt.Sprintf("[TransferUndeleagtedTokens for leveraged position] success on %s", chainId))
	return nil
}
