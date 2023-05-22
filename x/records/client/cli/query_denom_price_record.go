package cli

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/soohoio/stayking/v3/x/records/types"
	"github.com/spf13/cobra"
)

func CmdListDenomPriceRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-denom-price",
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

func CmdGetDenomPriceRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-denom-price [base-denom] [target-denom]",
		Short: "show a latest price of target-denom",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetDenomPriceRecordRequest{
				BaseDenom:   args[0],
				TargetDenom: args[1],
			}

			res, err := queryClient.DenomPriceRecord(context.Background(), params)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
