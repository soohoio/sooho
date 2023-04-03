package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSendQueryBalance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-query-balance [channel-id] [address] [denom]",
		Short: "Query the balances of an account on the remote chain via ICQ",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}


			msg := types.NewMsgSendQueryBalance(
				clientCtx.GetFromAddress().String(),
				args[0], // channel id
				args[1], // address
				args[2],
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "send query all balances")

	return cmd
}
