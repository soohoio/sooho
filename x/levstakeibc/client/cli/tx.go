package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/soohoio/stayking/v2/x/levstakeibc/client/tx"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"github.com/spf13/cobra"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		tx.CmdRegisterHostZone(),
		tx.CmdLeverageStake(),
		tx.CmdAddValidator(),
		tx.CmdRedeemStake(),
		tx.CmdExitLeverageStake(),
	)

	return cmd
}
