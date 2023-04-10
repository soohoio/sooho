package lendingpool

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v2/x/lendingpool/keeper"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	yearBlocks := k.GetParams(ctx).BlocksPerYear
	pools := k.GetAllPools(ctx)
	loans := k.GetAllLoans(ctx)
	for _, p := range pools {
		apr := p.GetInterestModel().GetAPR(p.GetUtilizationRate())
		if apr.Equal(sdk.ZeroDec()) {
			continue
		}
		blockInterest := apr.Quo(sdk.NewDec(int64(yearBlocks))).Mul(p.GetUtilizationRate())
		blockInterestMultiplier := sdk.OneDec().Add(blockInterest)
		// protocol tax will remain within the pool until it can be withdrawn by admin.
		// lender_interest_rate = block_interest_rate * (1 - protocol tax rate)
		lenderInterest := blockInterest.Mul(sdk.OneDec().Sub(k.GetParams(ctx).ProtocolTaxRate))
		lenderInterestMultiplier := sdk.OneDec().Add(lenderInterest)
		p.RedemptionRate = p.RedemptionRate.Mul(lenderInterestMultiplier)

		decCoins := sdk.NewDecCoinsFromCoins(p.TotalCoins...)
		p.TotalCoins, _ = decCoins.MulDec(blockInterestMultiplier).TruncateDecimal()
		k.SetPool(ctx, p)

		// TODO: improve this iteration
		for _, l := range loans {
			if l.Denom == p.Denom {
				l.BorrowedValue = l.BorrowedValue.Mul(blockInterestMultiplier)
				// TODO: process liquidation
				k.SetLoan(ctx, l)
			}
		}
	}
}

func Endblocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) {

}
