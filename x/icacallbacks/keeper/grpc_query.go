package keeper

import (
	"github.com/soohoio/stayking/v3/x/icacallbacks/types"
)

var _ types.QueryServer = Keeper{}
