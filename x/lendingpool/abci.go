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
		loanInterest := apr.Quo(sdk.NewDec(int64(yearBlocks)))

		blockInterest := loanInterest.Mul(p.GetUtilizationRate())
		blockInterestMultiplier := sdk.OneDec().Add(blockInterest)
		// protocol tax will remain within the pool until it can be withdrawn by admin.
		// lender_interest_rate = block_interest_rate * (1 - protocol tax rate)
		lenderInterest := blockInterest.Mul(sdk.OneDec().Sub(k.GetParams(ctx).ProtocolTaxRate))
		lenderInterestMultiplier := sdk.OneDec().Add(lenderInterest)

		p.RedemptionRate = p.RedemptionRate.Mul(lenderInterestMultiplier)

		prevTotalCoins := p.TotalCoins

		p.TotalCoins = p.TotalCoins.Mul(blockInterestMultiplier)
		k.SetPool(ctx, p)

		// some variables to calculate borrow interests
		interest := p.TotalCoins.Sub(prevTotalCoins)
		totalBorrow := p.TotalCoins.Sub(p.RemainingCoins)
		loans := k.GetPoolLoans(ctx, p.Denom)
		// TODO: improve this iteration
		for i, l := range loans {
			if l.Denom == p.Denom {
				var entryInterest sdk.Dec
				// if last of the loans, take the remaining interest
				if i == len(loans)-1 {
					entryInterest = interest
				} else {
					entryInterest = interest.Mul(l.BorrowedValue).Quo(totalBorrow)
					interest = interest.Sub(entryInterest)
				}

				l.BorrowedValue = l.BorrowedValue.Add(entryInterest)
				if l.GetDebtRatio().GT(p.LiquidationThreshold) {
					k.Liquidate(ctx, l.Id)
					l.Active = false
				}
				k.SetLoan(ctx, l)
			}
		}
	}
}

func Endblocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) {

}
