package keeper

import (
	"github.com/soohoio/stayking/v2/x/icacallbacks/types"
)

var _ types.QueryServer = Keeper{}
