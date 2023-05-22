package keeper

import (
	"github.com/soohoio/stayking/v3/x/stakeibc/types"
)

var _ types.QueryServer = Keeper{}
