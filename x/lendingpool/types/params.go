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
	ParamStoreKeyBlocksPerYear   = []byte("blocksperyear")
)

// default params
var (
	DefaultTaxRate              = sdk.NewDecWithPrec(1, 1) // 10%
	DefaultBlocksPerYear uint64 = 6307200                  // 86400 * 365 / 5
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new params
func NewParams(taxRate sdk.Dec, blocksPerYear uint64) Params {
	return Params{
		ProtocolTaxRate: taxRate,
		BlocksPerYear:   blocksPerYear,
	}
}

// DefaultParams returns default distribution parameters
func DefaultParams() Params {
	return NewParams(
		DefaultTaxRate,
		DefaultBlocksPerYear,
	)
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyProtocolTaxRate, &p.ProtocolTaxRate, validateProtocolTaxRate),
		paramtypes.NewParamSetPair(ParamStoreKeyBlocksPerYear, &p.BlocksPerYear, validateBlocksPerYear),
	}
}

// ValidateBasic performs basic validation on distribution parameters.
func (p Params) ValidateBasic() error {
	err := validateProtocolTaxRate(p.ProtocolTaxRate)
	if err != nil {
		return err
	}
	return validateBlocksPerYear(p.BlocksPerYear)
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

func validateBlocksPerYear(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("blocks per year is 0")
	}
	return nil
}