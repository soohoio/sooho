package types

import host "github.com/cosmos/ibc-go/v5/modules/core/24-host"

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId: PortID,
	}
}

func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	return nil
}
