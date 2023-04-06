package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string) *cobra.Command {

	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Query Commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdGetAllHostZone())
	cmd.AddCommand(CmdGetHostZone())

	return cmd
}
