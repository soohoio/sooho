package lendingpool

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v2/x/lendingpool/keeper"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	yearBlocks := k.GetParams(ctx).BlocksPerYear
	pools := k.GetAllPools(ctx)
	for _, p := range pools {
		apr := p.GetInterestModel().GetAPR(p.GetUtilizationRate())
		if apr.Equal(sdk.ZeroDec()) {
			continue
		}
		blockInterest := apr.Quo(sdk.NewDec(int64(yearBlocks)))
		// TODO: what to do with the protocol tax?
		p.RedemptionRate = p.RedemptionRate.Mul(sdk.OneDec().Mul(blockInterest))
		k.SetPool(ctx, p)
	}
}

func Endblocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) {

}
