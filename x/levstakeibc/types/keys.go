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
	EpochTrackerKeyPrefix = "EpochTracker-"
)
// prefix bytes for the interchainquery persistent store
const (
	prefixData  = iota + 1
	prefixQuery = iota + 1
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("interchainquery-port-")
)

var (
	KeyPrefixData  = []byte{prefixData}
	KeyPrefixQuery = []byte{prefixQuery}
)

// EpochTrackerKey returns the store key to retrieve a EpochTracker from the index fields
func EpochTrackerKey(epochIdentifier string) []byte {
	var key []byte

	epochIdentifierBytes := []byte(epochIdentifier)
	key = append(key, epochIdentifierBytes...)
	key = append(key, []byte("/")...)

	return key
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
