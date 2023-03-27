package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/soohoio/stayking/v2/x/lendingpool/types"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func NewTxCmd() *cobra.Command {
	lendingpoolTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "lending pool transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	lendingpoolTxCmd.AddCommand(
		NewCreatePoolCmd(),
		NewDepositCmd(),
		NewWithdrawCmd(),
	)

	return lendingpoolTxCmd
}
func NewCreatePoolCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool [base-denom] [interset-model]",
		Args:  cobra.ExactArgs(2),
		Short: "create a new lending pool initialized with a base denom and an interest model",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a Shield pool. Can only be executed from the Shield admin address.

Example:
$ %s tx lendingpool create-pool [base-denom] [interest-model]
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

			// TODO: implement general interest model
			// interestModel := args[1]
			interestRate, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePool(fromAddr, denom, interestRate)

			return tx.GenerateOrBroadcastTxWithFactory(cliCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewDepositCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [pool id] [amount]",
		Short: "deposit asset into a lending pool",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(cliCtx, cmd.Flags()).WithTxConfig(cliCtx.TxConfig).WithAccountRetriever(cliCtx.AccountRetriever)

			fromAddr := cliCtx.GetFromAddress().String()

			poolID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDeposit(fromAddr, uint64(poolID), amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxWithFactory(cliCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewWithdrawCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw [pool id] [amount]",
		Short: "withdraw asset from a lending pool. Accepts ib tokens as input.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(cliCtx, cmd.Flags()).WithTxConfig(cliCtx.TxConfig).WithAccountRetriever(cliCtx.AccountRetriever)

			fromAddr := cliCtx.GetFromAddress().String()

			poolID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdraw(fromAddr, uint64(poolID), amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxWithFactory(cliCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
