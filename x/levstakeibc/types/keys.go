package types

import "encoding/binary"

const (
	ModuleName = "levstakeibc"

	StoreKey = ModuleName

	RouterKey = ModuleName

	MemStoreKey = "mem_" + ModuleName

	QuerierRoute = ModuleName

	PortID = ModuleName
)

const (
	HostZoneKey           = "HostZone-value-"
	LevStakeInfoKeyPrefix = "LevStake-value-"
	EpochTrackerKeyPrefix = "EpochTracker/value/"
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

	PositionKey       = []byte{0x01}
	NextPositionIDKey = []byte{0x02}
)

func GetPositionKey(id uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, id)
	return append(PositionKey, b...)
}

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
