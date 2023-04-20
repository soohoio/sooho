package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

const (
	// DefaultHostEnabled is the default value for the host param (set to true)
	DefaultHostEnabled            = true
	DefaultPriceQueryChannelId    = "channel-1"
	DefaultPriceQueryPoolId       = "1"
	DefaultPriceQueryRoutesPoolId = "1"
	//@TODO must be usdc denom
	DefaultPriceQueryTokenInDenom = "uosmo"
	DefaultPriceQueryTokenOut     = "1uosmo"
	DefaultPriceQueryPath         = "/osmosis.poolmanager.v1beta1.Query/EstimateSwapExactAmountOut"
)

var (
	// KeyHostEnabled is the store key for HostEnabled Params
	KeyHostEnabled = []byte("HostEnabled")
	// KeyAllowQueries is the store key for the AllowQueries Params
	KeyAllowQueries           = []byte("AllowQueries")
	KeyPriceQueryChannelId    = []byte("PriceQueryChannelId")
	KeyPriceQueryPoolId       = []byte("PriceQueryPoolId")
	KeyPriceQueryRoutesPoolId = []byte("PriceQueryRoutesPoolId")
	//@TODO must be usdc denom
	KeyPriceQueryTokenInDenom = []byte("PriceQueryTokenInDenom")
	KeyPriceQueryTokenOut     = []byte("PriceQueryTokenOut")
	KeyPriceQueryPath         = []byte("PriceQueryPath")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	enableHost bool,
	allowQueries []string,
	priceQueryChannelId string,
	priceQueryPoolId string,
	priceQueryRoutesPoolId string,
	priceQueryTokenInDenom string,
	priceQueryTokenOut string,
	priceQueryPath string,
) Params {
	return Params{
		HostEnabled:            enableHost,
		AllowQueries:           allowQueries,
		PriceQueryChannelId:    priceQueryChannelId,
		PriceQueryPoolId:       priceQueryPoolId,
		PriceQueryRoutesPoolId: priceQueryRoutesPoolId,
		PriceQueryTokenInDenom: priceQueryTokenInDenom,
		PriceQueryTokenOut:     priceQueryTokenOut,
		PriceQueryPath:         priceQueryPath,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultHostEnabled,
		nil,
		DefaultPriceQueryChannelId,
		DefaultPriceQueryPoolId,
		DefaultPriceQueryRoutesPoolId,
		DefaultPriceQueryTokenInDenom,
		DefaultPriceQueryTokenOut, DefaultPriceQueryPath)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {

	return paramtypes.ParamSetPairs{paramtypes.NewParamSetPair(KeyHostEnabled, &p.HostEnabled, validateParam),
		paramtypes.NewParamSetPair(KeyAllowQueries, &p.AllowQueries, validateParam),
		paramtypes.NewParamSetPair(KeyPriceQueryChannelId, &p.PriceQueryChannelId, validateParam),
		paramtypes.NewParamSetPair(KeyPriceQueryPoolId, &p.PriceQueryPoolId, validateParam),
		paramtypes.NewParamSetPair(KeyPriceQueryRoutesPoolId, &p.PriceQueryRoutesPoolId, validateParam),
		paramtypes.NewParamSetPair(KeyPriceQueryTokenInDenom, &p.PriceQueryTokenInDenom, validateParam),
		paramtypes.NewParamSetPair(KeyPriceQueryTokenOut, &p.PriceQueryTokenOut, validateParam),
		paramtypes.NewParamSetPair(KeyPriceQueryPath, &p.PriceQueryPath, validateParam),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

func validateParam(i interface{}) error {
	return nil
}
