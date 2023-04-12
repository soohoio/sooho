package tripleslope

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v2/x/lendingpool/types"
)

const (
	InterestModelTypeTripleSlope = "TripleSlope"
)

var (
	_ types.InterestModelI = &TripleSlope{}
)

func NewTripleSlopeModel(params Params) (TripleSlope, error) {
	return TripleSlope{
		Params: params,
	}, nil
}

// GetAPR calculates the APR given the utilization rate with the parameters.
// it doesn't require there are 3 different slope sections.
// the APR returned is borrower APR, so the lender APR is (utilization_rate * returned APR)
func (ts TripleSlope) GetAPR(utilizationRate sdk.Dec) sdk.Dec {
	for i, r := range ts.Params.R {
		if utilizationRate.LT(r) {
			return ts.Params.M[i].Mul(utilizationRate).Add(ts.Params.B[i])
		}
	}

	// not assuming there are 3 slopes
	return ts.Params.M[len(ts.Params.R)-1].Mul(utilizationRate).Add(ts.Params.B[len(ts.Params.R)-1])
}

func (ts TripleSlope) ModelType() string {
	return InterestModelTypeTripleSlope
}

func (ts TripleSlope) String() string {
	out, err := json.Marshal(ts)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func (ts TripleSlope) ValidateBasic() error {
	return ts.Params.ValidateBasic()
}
