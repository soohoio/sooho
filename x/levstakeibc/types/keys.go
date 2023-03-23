package types

const (
	ModuleName = "levstakeibc"

	StoreKey = ModuleName

	RouterKey = ModuleName

	MemStoreKey = "mem_" + ModuleName

	QuerierRoute = ModuleName

	PortID = ModuleName
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	HostZoneKey = "HostZone-"
)
