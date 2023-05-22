package tx

import (
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/soohoio/stayking/v3/x/levstakeibc/types"
)

var _ = strconv.Itoa(0)

func CmdChangeValidatorWeight() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "change-validator-weight [host-zone] [address] [weight]",
		Short: "Broadcast message change-validator-weight",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argHostZone := args[0]
			argAddress := args[1]
			argWeight, err := cast.ToUint64E(args[2])
			if err != nil {
				return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "weight parse error: (%v)", err.Error())
			}

			msg := types.NewMsgChangeValidatorWeight(
				clientCtx.GetFromAddress().String(),
				argHostZone,
				argAddress,
				argWeight,
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
