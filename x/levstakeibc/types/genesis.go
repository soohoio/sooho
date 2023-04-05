package types

import host "github.com/cosmos/ibc-go/v5/modules/core/24-host"

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	queries := []Query{}
	return &GenesisState{
		PortId:  PortID,
		Queries: queries,
	}
}

func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	return nil
}
