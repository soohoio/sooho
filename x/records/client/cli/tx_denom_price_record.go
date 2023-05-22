package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v3/x/records/types"
	"github.com/spf13/cobra"
)

func CmdUpdateDenomPriceRecord() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "update-denom-price [base-denom] [target-denom] [price]",
		Short: "broadcast tx message update-denom-price",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argBaseDenom := args[0]
			argTargetDenom := args[1]
			argPrice, _ := sdk.NewIntFromString(args[2])
			msg := types.NewMsgUpdateDenomPrice(
				clientCtx.GetFromAddress().String(),
				argBaseDenom,
				argTargetDenom,
				argPrice,
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
