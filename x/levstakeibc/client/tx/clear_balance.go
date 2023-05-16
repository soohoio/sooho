package tx

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"

	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

var _ = strconv.Itoa(0)

func CmdClearBalance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear-balance [chain-id] [amount] [channel-id]",
		Short: "Broadcast message clear-balance",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChainId := args[0]
			argAmount, found := sdk.NewIntFromString(args[1])
			if !found {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidType, "can not convert string to int")
			}
			argChannelId := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgClearBalance(
				clientCtx.GetFromAddress().String(),
				argChainId,
				argAmount,
				argChannelId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}