package types

type Pools []Pool

func (p Pool) GetInterestModel() InterestModelI {
	model, ok := p.InterestModel.GetCachedValue().(InterestModelI)
	if !ok {
		return nil
	}
	return model
}
