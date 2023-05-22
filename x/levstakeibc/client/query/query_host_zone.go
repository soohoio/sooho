package query

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"github.com/spf13/cobra"
)

func CmdGetAllHostZone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-host-zone",
		Short: "Get all registered host zones",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())

			if err != nil {
				return err
			}

			params := &types.QueryAllHostZoneRequest{
				Pagination: pageReq,
			}

			res, err := types.NewQueryClient(clientCtx).AllHostZone(context.Background(), params)

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetHostZone() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-host-zone [chain-id]",
		Short: "show a host zone",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			chainId := args[0]

			if chainId == "" {
				return errorsmod.Wrap(sdkerrors.ErrInvalidType, "[chain-id] can not be empty")
			}

			params := &types.QueryGetHostZoneRequest{
				ChainId: chainId,
			}

			res, err := types.NewQueryClient(clientCtx).HostZone(context.Background(), params)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
