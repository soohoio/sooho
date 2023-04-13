package lendingpool

import (
	"fmt"
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
		loanInterestMultiplier := sdk.OneDec().Add(loanInterest)

		blockInterest := loanInterest.Mul(p.GetUtilizationRate())
		blockInterestMultiplier := sdk.OneDec().Add(blockInterest)
		// protocol tax will remain within the pool until it can be withdrawn by admin.
		// lender_interest_rate = block_interest_rate * (1 - protocol tax rate)
		lenderInterest := blockInterest.Mul(sdk.OneDec().Sub(k.GetParams(ctx).ProtocolTaxRate))
		lenderInterestMultiplier := sdk.OneDec().Add(lenderInterest)
		fmt.Println("abci")
		fmt.Println(blockInterest.String())
		fmt.Println(blockInterestMultiplier.String())
		fmt.Println(lenderInterest.String())
		fmt.Println(lenderInterestMultiplier.String())
		fmt.Println(p.RemainingCoins.String())
		fmt.Println(p.TotalCoins.String())
		p.RedemptionRate = p.RedemptionRate.Mul(lenderInterestMultiplier)

		p.TotalCoins = p.TotalCoins.Mul(blockInterestMultiplier)
		k.SetPool(ctx, p)

		loans := k.GetPoolLoans(ctx, p.Denom)
		// TODO: improve this iteration
		for _, l := range loans {
			if l.Denom == p.Denom {
				l.BorrowedValue = l.BorrowedValue.Mul(loanInterestMultiplier)
				// TODO: process liquidation
				k.SetLoan(ctx, l)
			}
		}
	}
}

func Endblocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) {

}
