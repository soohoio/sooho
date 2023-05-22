package tx

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/soohoio/stayking/v3/x/levstakeibc/types"
	"github.com/spf13/cobra"
	"strconv"
)

var _ = strconv.Itoa(0)

func CmdRedeemStake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redeem-stake [stTokenAmount] [chainId] [receiver]",
		Short: "Broadcast message redeem-stake",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			argStTokenAmount, found := sdk.NewIntFromString(args[0])

			if !found {
				return errorsmod.Wrap(sdkerrors.ErrInvalidType, "can not convert string to int")
			}

			chainId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)

			if err != nil {
				return err
			}

			argReceiver := args[2]

			msg := types.NewMsgRedeemStake(
				clientCtx.GetFromAddress().String(), // Creator
				argStTokenAmount,
				chainId,
				argReceiver,
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
