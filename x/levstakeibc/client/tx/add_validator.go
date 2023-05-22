package tx

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/soohoio/stayking/v3/x/levstakeibc/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdAddValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-validator [host-zone] [name] [address] [commission] [weight]",
		Short: "Broadcast message add-validator",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argHostZone := args[0]
			argName := args[1]
			argAddress := args[2]

			argCommission, err := cast.ToUint64E(args[3])
			if err != nil {
				return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "commission parse error: (%v)", err.Error())
			}

			argWeight, err := cast.ToUint64E(args[4])
			if err != nil {
				return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "weight parse error: (%v)", err.Error())
			}

			msg := types.NewMsgAddValidator(
				clientCtx.GetFromAddress().String(),
				argHostZone,
				argName,
				argAddress,
				argCommission,
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
