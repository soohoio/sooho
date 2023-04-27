package tx

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"github.com/spf13/cobra"
	"strconv"
)

func CmdAdjustPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "adjust-position [positionId] [collateral] [debt] [hostDenom]",
		Short: "Broadcast tx message adjust-position",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argPositionId, err := strconv.ParseUint(args[0], 0, 64)

			if err != nil {
				errorsmod.Wrap(sdkerrors.ErrInvalidType, "the positionId can not convert string to uint")
			}

			collateral, ok := sdk.NewIntFromString(args[1])

			if !ok {
				errorsmod.Wrap(sdkerrors.ErrInvalidType, "the collateral can not convert string to uint")
			}

			debt, ok := sdk.NewIntFromString(args[2])

			if !ok {
				errorsmod.Wrap(sdkerrors.ErrInvalidType, "the debt can not convert string to uint")
			}

			hostDenom := args[3]

			msg := types.NewMsgAdjustPosition(
				clientCtx.GetFromAddress().String(),
				argPositionId,
				collateral,
				debt,
				hostDenom,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
