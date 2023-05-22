package cli

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/soohoio/stayking/v3/x/records/types"
)

var DefaultRelativePacketTimeoutTimestamp = cast.ToUint64((time.Duration(10) * time.Minute).Nanoseconds())

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// TODO: 임시 denom price 저장하기 ( 향후 삭제 필요 )
	cmd.AddCommand(CmdUpdateDenomPriceRecord())

	return cmd
}
