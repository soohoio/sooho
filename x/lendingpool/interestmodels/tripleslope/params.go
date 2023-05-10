package tripleslope

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	// default interest rate:
	// UR 0~60%: 0~20% linearly increasing
	// UR 60%~80%: 20% fixed
	// UR 80%~100%: 20%~140%
	// function: y = m_ix+b_i where r_{i+1} < x < r_i

	DefaultR = []sdk.Dec{
		sdk.NewDecWithPrec(6, 1),
		sdk.NewDecWithPrec(9, 1),
	}
	DefaultM = []sdk.Dec{
		sdk.NewDecWithPrec(333, 3), // 1/3
		sdk.ZeroDec(),
		sdk.NewDecFromInt(sdk.NewInt(6)),
	}
	DefaultB = []sdk.Dec{sdk.ZeroDec(),
		sdk.NewDecWithPrec(2, 1),
		sdk.NewDecWithPrec(-46, 1),
	}

	MaxGap = sdk.NewDecWithPrec(1, 1) // allow 10% gap between interest rate sections
)

func NewParams(r, m, b []sdk.Dec) Params {
	return Params{
		R: r,
		M: m,
		B: b,
	}
}

func DefaultParams() Params {
	return Params{
		R: DefaultR,
		M: DefaultM,
		B: DefaultB,
	}
}

func (p Params) String() string {
	out, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func (p Params) ValidateBasic() error {
	// validatoe number of params
	//if len(p.M) != len(p.R)+1 || len(p.B) != len(p.R)+1 {
	//	return types.ErrInvalidModelParams
	//}
	//
	//// validate slope is *somewhat* continuous
	//for i, r := range p.R {
	//	// left side
	//	left := p.M[i].Mul(r).Add(p.B[i])
	//	// right side
	//	right := p.M[i+1].Mul(r).Add(p.B[i+1])
	//
	//	// allow very little gap
	//	if left.Sub(right).Abs().GT(MaxGap) {
	//		// TODO: print the unmatching conditions for debugging
	//		return types.ErrInvalidModelParams
	//	}
	//}
	return nil
}
