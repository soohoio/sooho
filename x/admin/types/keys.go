package types

const (
	ModuleName = "admin"

	StoreKey = ModuleName

	RouterKey = ModuleName

	QuerierRoute = ModuleName
)

var (
	AdminKey = []byte{0x01}
)

func GetAdminKey() []byte {
	return AdminKey
}
