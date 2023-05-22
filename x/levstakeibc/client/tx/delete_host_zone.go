package tx

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/soohoio/stayking/v3/x/levstakeibc/types"
	"github.com/spf13/cobra"
)

func CmdDeleteHostZone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-host-zone [chain-id]",
		Short: "Broadcast message update-host-zone",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainId := args[0]

			msg := types.NewMsgDeleteHostZone(
				clientCtx.GetFromAddress().String(),
				chainId,
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
