package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/lendingpool/types"
)

func getSubInt(d sdk.Dec) sdk.Dec {
	return d.Sub(sdk.NewDecFromInt(d.TruncateInt()))
}

func getIBDenom(base_denom string) string {
	return types.IBPrefix + base_denom
}
