package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/soohoio/stayking/x/stakeibc/types"
)

var _ = strconv.Itoa(0)

func CmdSetWithdrawalAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-withdrawal-address [delegatorAddress] [withdrawAddress]",
		Short: "Broadcast message set-withdrawal-address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			delegatorAddress := args[0]
			withdrawAddress := args[1]



			msg := types.NewMsgSetWithdrawalAddress(
				delegatorAddress,
				withdrawAddress,
			)

			//if err := msg.ValidateBasic(); err != nil {
			//	return err
			//}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
