package tx

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/soohoio/stayking/v3/x/levstakeibc/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdExitLeverageStake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exit-leverage-stake [positionId] [chainId] [receiver]",
		Short: "Broadcast tx message exit-leverage-stake",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argPositionId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argChainId := args[1]
			argReceiver := args[2]

			msg := types.NewMsgExitLeverageStake(
				clientCtx.GetFromAddress().String(),
				argPositionId,
				argChainId,
				argReceiver,
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
