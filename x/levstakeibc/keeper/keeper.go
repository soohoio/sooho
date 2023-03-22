package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

type Keeper struct {
	cdc           codec.BinaryCodec
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}
