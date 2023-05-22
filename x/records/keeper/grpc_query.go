package keeper

import (
	"github.com/soohoio/stayking/v3/x/records/types"
)

var _ types.QueryServer = Keeper{}
