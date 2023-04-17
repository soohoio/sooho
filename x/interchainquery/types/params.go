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
	KeyAllowQueries        = []byte("AllowQueries")
	PriceQueryChannelId    = "channel-1"
	PriceQueryPoolId       = "1"
	PriceQueryRoutesPoolId = "1"
	//@TODO must be usdc denom
	PriceQueryTokenInDenom = "uosmo"
	PriceQueryTokenOut     = "1uosmo"
	PriceQueryPath         = "/osmosis.poolmanager.v1beta1.Query/EstimateSwapExactAmountOut"
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
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}
