package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/soohoio/stayking/v3/x/interchainquery/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSetIcqParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-icq-params [channel-id] [pool-id] [routes-id] [token-in-denom] [tokenout] [path]",
		Short: "Query the balances of an account on the remote chain via ICQ",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetIcqParams(
				clientCtx.GetFromAddress().String(),
				args[0], // channel id
				args[1], // address
				args[2],
				args[3],
				args[4],
				args[5],
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "set icq params")

	return cmd
}
