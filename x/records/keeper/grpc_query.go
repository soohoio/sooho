package keeper

import (
	"github.com/soohoio/stayking/x/records/types"
)

var _ types.QueryServer = Keeper{}
