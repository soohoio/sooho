package tx

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/soohoio/stayking/v3/x/levstakeibc/types"
	"github.com/spf13/cobra"
)

func CmdRestoreInterchainAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restore-interchain-account [chain-id] [account-type]",
		Short: "Broadcast message restore-interchain-account",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			argChainId := args[0]
			argAccountType := types.ICAType_value[args[1]]

			msg := types.NewMsgRestoreInterchainAccount(
				creator,
				argChainId,
				types.ICAType(argAccountType),
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
