package cli

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/soohoio/stayking/v2/x/records/types"
)

func CmdListDenomPriceRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-denom-price-record",
		Short: "list all denom price records",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDenomPriceRecordRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DenomPriceRecordAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
