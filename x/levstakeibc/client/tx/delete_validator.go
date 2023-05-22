package tx

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

func CmdDeleteValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-validator [host-zone] [address]",
		Short: "Broadcast message delete-validator",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argHostZone := args[0]
			argAddress := args[1]

			msg := types.NewMsgDeleteValidator(
				clientCtx.GetFromAddress().String(),
				argHostZone,
				argAddress,
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
