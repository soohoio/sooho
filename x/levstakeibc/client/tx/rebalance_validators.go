package tx

import (
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/soohoio/stayking/v3/x/levstakeibc/types"
)

var _ = strconv.Itoa(0)

func CmdRebalanceValidators() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rebalance-validators [host-zone] [num-to-rebalance]",
		Short: "Broadcast message rebalanceValidators",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argHostZone := args[0]
			argNumValidators, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "number of validators parse error: (%v)", err.Error())
			}

			msg := types.NewMsgRebalanceValidators(
				clientCtx.GetFromAddress().String(),
				argHostZone,
				argNumValidators,
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
