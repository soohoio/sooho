package tripleslope

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v2/x/lendingpool/types"
)

var _ types.InterestModelI = &TripleSlope{}

func NewTripleSlopeModel(params Params) (TripleSlope, error) {
	return TripleSlope{
		Params: params,
	}, nil
}

func (ts TripleSlope) GetAPR(utilizationRate sdk.Dec) sdk.Dec {
	for i, r := range ts.Params.R {
		if utilizationRate.LT(r) {
			return ts.Params.M[i].Mul(utilizationRate).Add(ts.Params.B[i])
		}
	}

	// assume there is only 3 slopes, since "triple" slope
	return ts.Params.M[2].Mul(utilizationRate).Add(ts.Params.B[2])
}

func (ts TripleSlope) String() string {
	out, err := json.Marshal(ts)
	if err != nil {
		panic(err)
	}
	return string(out)
}
