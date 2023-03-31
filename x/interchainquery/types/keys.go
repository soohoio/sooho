package types

import icqtypes "github.com/strangelove-ventures/async-icq/v5/types"

const (
	// ModuleName defines the module name
	ModuleName = "interchainquery"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_interchainquery"

	// Version defines the current version the IBC module supports
	Version = icqtypes.Version


	PortID = ModuleName
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

// keys for proof queries to various stores, note: there's an implicit assumption here that
// the stores on the counterparty chain are prefixed with the standard cosmos-sdk module names
// this might not be true for all IBC chains, and is something we should verify before onboarding a
// new chain

const (
	STAKING_STORE_QUERY_WITH_PROOF = "store/staking/key"
	BANK_STORE_QUERY_WITH_PROOF    = "store/bank/key"
)

var (
	KeyPrefixData  = []byte{prefixData}
	KeyPrefixQuery = []byte{prefixQuery}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
