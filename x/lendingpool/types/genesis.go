package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewGenesisState creates a new genesis state for the governance module
func NewGenesisState(params Params, pools []Pool, nextPoolId uint64) *GenesisState {
	return &GenesisState{
		Params:     params,
		Pools:      pools,
		NextPoolId: nextPoolId,
	}
}

// DefaultGenesisState creates a new default lending pool genesis state.
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		DefaultParams(),
		[]Pool{},
		1,
	)
}

func ValidateGenesis(data GenesisState) error {
	if data.NextPoolId == 0 {
		return ErrInvalidPoolID
	}

	if data.Params.ProtocolTaxRate.LT(sdk.ZeroDec()) ||
		data.Params.ProtocolTaxRate.LT(sdk.ZeroDec()) {
		return ErrInvalidProtocolTaxRate
	}

	return nil
}
