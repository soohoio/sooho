package keeper

import (
	"github.com/soohoio/stayking/v2/x/records/types"
)

var _ types.QueryServer = Keeper{}
