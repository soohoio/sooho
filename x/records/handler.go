package records

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/soohoio/stayking/v2/x/records/keeper"
	"github.com/soohoio/stayking/v2/x/records/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		_ = ctx
		k.Logger(ctx).Info("!!!!!!!!!!!!!! here1 !!!!!!!!!!!!!!!!")

		switch msg := msg.(type) {
		case *types.MsgUpdateDenomPrice:
			k.Logger(ctx).Info("!!!!!!!!!!!!!! here2 !!!!!!!!!!!!!!!!")
			res, err := msgServer.UpdateDenomPrice(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, errorsmod.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
