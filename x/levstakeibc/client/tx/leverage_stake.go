package tx

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	"github.com/spf13/cobra"
)

func CmdLeverageStake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "leverage-stake [equity] [hostDenom] [leverageRatio] [receiver]",
		Short: "Broadcast tx message leverage-stake",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argEquity, found := sdk.NewIntFromString(args[0])

			if !found {
				errorsmod.Wrap(sdkerrors.ErrInvalidType, "can not convert string value to int")
			}

			hostDenom := args[1]

			leverageRatio, err := sdk.NewDecFromStr(args[2])

			if err != nil {
				errorsmod.Wrap(sdkerrors.ErrInvalidType, "can not convert string value to sdk.Dec")
			}

			msg := types.NewMsgLeverageStake(
				clientCtx.GetFromAddress().String(),
				argEquity,
				hostDenom,
				leverageRatio,
				args[3],
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
