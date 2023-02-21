package types

const (
	// ModuleName defines the module name
	ModuleName = "stakeibc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stakeibc"

	// Version defines the current version the IBC module supports
	Version = "stakeibc-1"

	// PortID is the default port id that module binds to
	PortID = "stakeibc"

	// fee account - F1
	FeeAccount = "sooho1cprgqxgjve4sf5sjeaugympg4hrtyr6hamfjtz"
)

// PortKey defines the key to store the port ID in store
var PortKey = KeyPrefix("stakeibc-port-")

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// EpochTrackerKey returns the store key to retrieve a EpochTracker from the index fields
func EpochTrackerKey(epochIdentifier string) []byte {
	var key []byte

	epochIdentifierBytes := []byte(epochIdentifier)
	key = append(key, epochIdentifierBytes...)
	key = append(key, []byte("/")...)

	return key
}

const (
	HostZoneKey           = "HostZone-value-"
	EpochTrackerKeyPrefix = "EpochTracker/value/"
)
