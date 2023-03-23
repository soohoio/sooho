package tripleslope

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewParams(r, m, b []sdk.Dec) Params {
	return Params{
		R: r,
		M: m,
		B: b,
	}
}

func (p Params) String() string {
	out, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(out)
}
