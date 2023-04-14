package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/soohoio/stayking/v2/x/mockborrow/types"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func NewTxCmd() *cobra.Command {
	lendingpoolTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "mockborrow transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	lendingpoolTxCmd.AddCommand(
		NewBorrowCmd(),
		NewRepayCmd(),
	)

	return lendingpoolTxCmd
}

func NewBorrowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "borrow [denom] [collateral amt] [borrow amt]",
		Args:  cobra.ExactArgs(3),
		Short: "borrow from lending pool and does nothing with it",
		Long: strings.TrimSpace(
			fmt.Sprintf(`borrow from lending pool and does nothing with it

Example:
$ %s tx mockborrow borrow [denom] [collateral amt] [borrow amt]
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(cliCtx, cmd.Flags()).WithTxConfig(cliCtx.TxConfig).WithAccountRetriever(cliCtx.AccountRetriever)

			fromAddr := cliCtx.GetFromAddress().String()
			denom := args[0]

			collateralAmount, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			borrowAmount, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgBorrow(fromAddr, denom, collateralAmount, borrowAmount)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxWithFactory(cliCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewRepayCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repay [loan id] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "repay for a registered loan",
		Long: strings.TrimSpace(
			fmt.Sprintf(`repay for a registered loan. 
accepts loan ID and corresponding coins and reduces the total size of the loan.

Example:
$ %s tx mockborrow borrow [loan id] [amount]
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(cliCtx, cmd.Flags()).WithTxConfig(cliCtx.TxConfig).WithAccountRetriever(cliCtx.AccountRetriever)

			fromAddr := cliCtx.GetFromAddress().String()
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("pool id %s is invalid", args[0])
			}

			amount, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgRepay(fromAddr, id, amount)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxWithFactory(cliCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
