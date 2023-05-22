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

func CmdGetAllPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-all-positions",
		Short: "Get all positions",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())

			params := &types.QueryAllPositionRequest{
				Pagination: pageReq,
			}

			res, err := types.NewQueryClient(clientCtx).AllPosition(context.Background(), params)

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

func CmdGetPositionListBySender() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-user-positions [staker-addr]",
		Short: "Get all positions by sender",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			staker := args[0]

			if staker == "" {
				return errorsmod.Wrap(sdkerrors.ErrInvalidType, "[staker-addr] can not be empty")
			}

			params := &types.QueryGetPositionListBySenderRequest{
				Sender: staker,
			}

			res, err := types.NewQueryClient(clientCtx).PositionListBySender(context.Background(), params)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
