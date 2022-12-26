package keeper

import (
	"github.com/soohoio/stayking/x/icacallbacks/types"
)

var _ types.QueryServer = Keeper{}
