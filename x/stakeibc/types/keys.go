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

	ValidatorKey  = "Validator-value-"
	DelegationKey = "Delegation-value-"

	MinValidatorRequirementsKey = "MinValidatorRequirements-value-"
	ICAAccountKey               = "ICAAccount-value-"

	// fee account - F1
	FeeAccount = "sooho12prkmv4cpegcnzp5yx9505cmu0ynmpz3kaffdc" //cross front loop emerge dial water pass critic odor travel maple pluck bus spider check report blossom round sorry stuff key sauce gown maximum

)

// PortKey defines the key to store the port ID in store
var PortKey = KeyPrefix("stakeibc-port-")

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	HostZoneKey      = "HostZone-value-"
	HostZoneCountKey = "HostZone-count-"
)
