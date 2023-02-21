package interchainquery

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/soohoio/stayking/x/interchainquery/keeper"
	"github.com/soohoio/stayking/x/interchainquery/types"
)

// NewHandler returns a handler for interchainquery module messages
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(_ sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
		return nil, errorsmod.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	}
}
