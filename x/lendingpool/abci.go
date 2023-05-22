package lendingpool

import (
	"fmt"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v3/x/lendingpool/keeper"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	yearBlocks := k.GetParams(ctx).BlocksPerYear
	pools := k.GetAllPools(ctx)
	for _, p := range pools {
		apr := p.GetInterestModel().GetAPR(p.GetUtilizationRate())
		if apr.Equal(sdk.ZeroDec()) {
			continue
		}
		// Default Setting 값이 0 이 아니기 때문에 division by zero 에 대한 문제 없음
		loanInterest := apr.Quo(sdk.NewDec(int64(yearBlocks)))

		blockInterest := loanInterest.Mul(p.GetUtilizationRate())
		blockInterestMultiplier := sdk.OneDec().Add(blockInterest)
		// protocol tax will remain within the pool until it can be withdrawn by admin.
		// lender_interest_rate = block_interest_rate * (1 - protocol tax rate)
		lenderInterest := blockInterest.Mul(sdk.OneDec().Sub(k.GetParams(ctx).ProtocolTaxRate))
		lenderInterestMultiplier := sdk.OneDec().Add(lenderInterest)
		lenderInterestMultiplierWithoutTax := sdk.OneDec().Add(blockInterest)
		p.RedemptionRate = p.RedemptionRate.Mul(lenderInterestMultiplier)
		p.RedemptionRateWithoutTax = p.RedemptionRateWithoutTax.Mul(lenderInterestMultiplierWithoutTax)
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
					// totalBorrow 가 0 인 경우는 연출되지 않는다.
					entryInterest = interest.Mul(l.BorrowedValue).Quo(totalBorrow)
					interest = interest.Sub(entryInterest)
				}

				l.BorrowedValue = l.BorrowedValue.Add(entryInterest)
				k.SetLoan(ctx, l)
				if l.GetDebtRatio().GT(p.MaxDebtRatio) && l.Active {
					err := k.Liquidate(ctx, l.Id)
					if err != nil {
						k.Logger(ctx).Error(fmt.Sprintf("[CUSTOM DEBUG] Liquidation error reason : %v", err.Error()))
						k.Logger(ctx).Error(fmt.Sprintf("[CUSTOM DEBUG] Liquidation executed from loan id %v, debtRatio %v", l.Id, l.GetDebtRatio()))
						continue
					}
					l.Active = false
					k.SetLoan(ctx, l)
				}
			}
		}
	}
}

func Endblocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) {

}
