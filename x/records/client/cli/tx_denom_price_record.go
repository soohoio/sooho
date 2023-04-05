package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func CmdUpdateDenomPriceRecord() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "update-denom-price [base-denom] [target-denom] [price]",
		Short: "broadcast tx message update-denom-price",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argBaseDenom := args[0]
			argTargetDenom := args[1]
			argPrice, err := types.NewDecFromStr(args[2])

			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			//msg := types.

			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
