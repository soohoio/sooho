package types

func NewGenesisState(queries []Query) *GenesisState {
	return &GenesisState{
		PortId: PortID,
		// this line is used by starport scaffolding # genesis/types/default
		Params:  DefaultParams(),
		Queries: queries}
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	queries := []Query{}
	return NewGenesisState(queries)
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// TODO: validate genesis state.
	return gs.Params.Validate()
	return nil
}
