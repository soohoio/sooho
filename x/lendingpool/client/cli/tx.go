package cli

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/soohoio/stayking/v3/x/lendingpool/interestmodels/tripleslope"
	"github.com/soohoio/stayking/v3/x/lendingpool/types"
	"github.com/spf13/cobra"
	"os"
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
		NewLiquidateCmd(),
	)

	return lendingpoolTxCmd
}

func NewCreatePoolCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool [base-denom] [interset-model] [max-debt-ratio]",
		Args:  cobra.ExactArgs(3),
		Short: "create a new lending pool initialized with a base denom and an interest model",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a lending pool.

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

			var interestModel types.InterestModelI
			interestModel, err = getInterestModel(cliCtx.Codec, args[1])
			if err != nil {
				return err
			}

			maxDebtRatio, err := sdk.NewDecFromStr(args[2])

			msg, err := types.NewMsgCreatePool(fromAddr, denom, maxDebtRatio, interestModel)
			if err != nil {
				return err
			}

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

// NewLiquidateCmd is command which triggers liquidation, just for unit test, not for production
func NewLiquidateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquidate [loan id]",
		Short: "liquidate loan, accepts loan id as a input.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(cliCtx, cmd.Flags()).WithTxConfig(cliCtx.TxConfig).WithAccountRetriever(cliCtx.AccountRetriever)

			fromAddr := cliCtx.GetFromAddress().String()

			loanId, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgLiquidate(fromAddr, uint64(loanId))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxWithFactory(cliCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

type interestModel struct {
	Type    string
	Content json.RawMessage
}

func getInterestModel(cdc codec.Codec, path string) (types.InterestModelI, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var im interestModel
	err = json.Unmarshal(contents, &im)
	if err != nil {
		return nil, err
	}

	switch im.Type {
	case tripleslope.InterestModelTypeTripleSlope:
		var model tripleslope.TripleSlope
		err = cdc.UnmarshalJSON(im.Content, &model)
		return &model, err
	default:
		return nil, types.ErrInvalidInterestModel
	}
}
