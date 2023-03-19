package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/x/lendingpool/types"
)

func (k Keeper) GetTaxRate(ctx sdk.Context) sdk.Dec {
	var taxRate sdk.Dec
	k.paramSpace.Get(ctx, types.ParamStoreKeyProtocolTaxRate, &taxRate)
	return taxRate
}

func (k Keeper) SetTaxRate(ctx sdk.Context, taxRate sdk.Dec) {
	k.paramSpace.Set(ctx, types.ParamStoreKeyProtocolTaxRate, taxRate)
}
