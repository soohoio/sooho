package types

const (
	ModuleName = "levstakeibc"

	StoreKey = ModuleName

	RouterKey = ModuleName

	MemStoreKey = "mem_" + ModuleName

	QuerierRoute = ModuleName

	PortID = ModuleName
)

const (
	HostZoneKey           = "HostZone-"
	EpochTrackerKeyPrefix = "EpochTracker/"
)

// EpochTrackerKey returns the store key to retrieve a EpochTracker from the index fields
func EpochTrackerKey(epochIdentifier string) []byte {
	var key []byte

	epochIdentifierBytes := []byte(epochIdentifier)
	key = append(key, epochIdentifierBytes...)
	key = append(key, []byte("/")...)

	return key
}
