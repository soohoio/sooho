package types

import "encoding/binary"

const (
	ModuleName = "levstakeibc"

	StoreKey = ModuleName

	RouterKey = ModuleName

	MemStoreKey = "mem_" + ModuleName

	QuerierRoute = ModuleName

	PortID = ModuleName

	FeeAccount = "sooho1cprgqxgjve4sf5sjeaugympg4hrtyr6hamfjtz"

	LendingPoolFeeAccount = "sooho19zyp74ptdgwfvgnx68fqzu2rh32judvxtpphva"
)

const (
	HostZoneKey           = "HostZone-value-"
	LevStakeInfoKeyPrefix = "LevStake-value-"
	EpochTrackerKeyPrefix = "EpochTracker/value/"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("interchainquery-port-")
)

var (
	KeyPrefixData           = []byte{0x01}
	KeyPrefixQuery          = []byte{0x02}
	KeyPrefixPosition       = []byte{0x03}
	KeyPrefixNextPositionID = []byte{0x04}
)

func GetPositionKey(id uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, id)
	return append(KeyPrefixPosition, b...)
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
