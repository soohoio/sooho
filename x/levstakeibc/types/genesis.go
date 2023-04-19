package types

import (
	"fmt"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:           PortID,
		Params:           DefaultParams(),
		HostZoneList:     []HostZone{},
		EpochTrackerList: []EpochTracker{},
		NextPositionId:   1,
	}
}

func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated index in hostZoneList
	hostZoneList := make(map[string]HostZone)
	for _, hostZone := range gs.HostZoneList {
		if _, ok := hostZoneList[hostZone.ChainId]; ok {
			return fmt.Errorf("duplicated index in hostZoneList: %s", hostZone.ChainId)
		}
		hostZoneList[hostZone.ChainId] = hostZone
	}

	// Check for duplicated index in epochTracker
	epochTrackerIndexMap := make(map[string]struct{})

	for _, elem := range gs.EpochTrackerList {
		index := string(EpochTrackerKey(elem.EpochIdentifier))
		if _, ok := epochTrackerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for epochTracker")
		}
		epochTrackerIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
