package tx

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"github.com/spf13/cobra"
)

func CmdRestoreInterchainAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "",
		Short: "restore-interchain-account [chain-id] [account-type]",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			argChannelId := args[0]
			argAccountType := types.ICAType_value[args[1]]

			msg := types.NewMsgRestoreInterchainAccount(
				creator,
				argChannelId,
				types.ICAType(argAccountType),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}
