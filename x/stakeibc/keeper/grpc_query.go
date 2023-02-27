package keeper

import (
	"github.com/soohoio/stayking/v2/x/stakeibc/types"
)

var _ types.QueryServer = Keeper{}
