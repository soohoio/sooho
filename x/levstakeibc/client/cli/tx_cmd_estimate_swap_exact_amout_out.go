package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdEstimateSwapExactAmountOut() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "estimate-swap-exact-amount-out <poolID> <sender> <tokenOut> <swap-route-pool-ids> <TokenInDenom>",
		Short: "Query estimate-swap-exact-amount-out",
		Long: `Query estimate-swap-exact-amount-out.{{.ExampleHeader}}
{{.CommandPrefix}} estimate-swap-exact-amount-out 1 osm11vmx8jtggpd9u7qr0t8vxclycz85u925sazglr7 1000stake --swap-route-pool-ids=2 --swap-route-pool-ids=3`,
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolID, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			routes:= []types.SwapAmountOutRoute{}
			pID, err := strconv.Atoi(args[3])
			if err != nil {
				return err
			}
			routes = append(routes, types.SwapAmountOutRoute{
				PoolId: uint64(pID),
				TokenInDenom: args[4],
			})
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			msg := types.NewSendEstimateSwapExactAmountOut(
				clientCtx.GetFromAddress().String(),
				uint64(poolID), //pool id
				routes,
				args[2],
				pageReq,

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
