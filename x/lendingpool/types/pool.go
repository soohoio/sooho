package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Pools []Pool

func (p Pool) GetInterestModel() InterestModelI {
	model, ok := p.InterestModel.GetCachedValue().(InterestModelI)
	if !ok {
		return nil
	}
	return model
}

func (p Pool) GetUtilizationRate() sdk.Dec {
	dividend := sdk.NewDecFromInt(p.TotalCoins.Sub(p.Coins...).AmountOf(p.Denom))
	divisor := sdk.NewDecFromInt(p.TotalCoins.AmountOf(p.Denom))
	if divisor.Equal(sdk.ZeroDec()) {
		return sdk.ZeroDec()
	}
	return dividend.Quo(divisor)
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Pool) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var model InterestModelI
	return unpacker.UnpackAny(p.InterestModel, &model)
}

func (ps Pools) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	for _, x := range ps {
		err := x.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}
