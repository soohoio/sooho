package types

import (
	"fmt"

	"sigs.k8s.io/yaml"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
	ParamStoreKeyProtocolTaxRate = []byte("protocoltaxrate")
)

// default params
var (
	DefaultTaxRate = sdk.NewDecWithPrec(1, 1) // 10%
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new params
func NewParams(taxRate sdk.Dec) Params {
	return Params{
		ProtocolTaxRate: taxRate,
	}
}

// DefaultParams returns default distribution parameters
func DefaultParams() Params {
	return Params{
		ProtocolTaxRate: DefaultTaxRate,
	}
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyProtocolTaxRate, &p.ProtocolTaxRate, validateProtocolTaxRate),
	}
}

// ValidateBasic performs basic validation on distribution parameters.
func (p Params) ValidateBasic() error {
	if p.ProtocolTaxRate.IsNegative() || p.ProtocolTaxRate.GT(sdk.OneDec()) {
		return fmt.Errorf(
			"community tax should be non-negative and less than one: %s", p.ProtocolTaxRate,
		)
	}
	return nil
}

func validateProtocolTaxRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("protocol tax rate must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("protocol tax rate must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("protocol tax rate is bigger than one: %s", v)
	}

	return nil
}
