package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/soohoio/stayking/v3/x/interchainquery/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSendEstimateSwapExactAmountOut() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-swap-exact-amount-out <poolID> <tokenOut> <swap-route-pool-ids> <TokenInDenom> <channel-id>",
		Short: "Query estimate-swap-exact-amount-out",
		Long: `Query estimate-swap-exact-amount-out.{{.ExampleHeader}}
{{.CommandPrefix}} estimate-swap-exact-amount-out 1 osm11vmx8jtggpd9u7qr0t8vxclycz85u925sazglr7 1000stake --swap-route-pool-ids=2 --swap-route-pool-ids=3`,
		Args: cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			routes := []types.SwapAmountOutRoute{}
			pID, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}
			routes = append(routes, types.SwapAmountOutRoute{
				PoolId:       uint64(pID),
				TokenInDenom: args[3],
			})

			msg := types.NewMsgSendEstimateSwapExactAmountOut(
				clientCtx.GetFromAddress().String(),
				uint64(poolID), //pool id
				routes,
				args[1],
				args[4],
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "send query all balances")

	return cmd
}
