package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
)

type Pools []Pool

func (p Pool) GetInterestModel() InterestModelI {
	model, ok := p.InterestModel.GetCachedValue().(InterestModelI)
	if !ok {
		return nil
	}
	return model
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
