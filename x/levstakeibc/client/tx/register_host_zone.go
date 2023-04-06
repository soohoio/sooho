package tx

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"github.com/spf13/cobra"
	"strconv"
)

func CmdRegisterHostZone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-host-zone [connection-id] [bech32prefix] [host-denom] [ibc-denom] [channel-id] [unbonding-frequency]",
		Short: "Broadcast message register-host-zone",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			connectionId := args[0]
			hostDenom := args[1]
			bech32prefix := args[2]
			ibcDenom := args[3]
			channelId := args[4]
			unbondingFrequency, err := strconv.ParseUint(args[5], 10, 64)
			if err != nil {
				return err
			}
			msg := types.NewMsgRegisterHostZone(
				connectionId,
				bech32prefix,
				hostDenom,
				ibcDenom,
				clientCtx.GetFromAddress().String(),
				channelId,
				unbondingFrequency,
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
