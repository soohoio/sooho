package types

import "encoding/binary"

const (
	ModuleName = "lendingpool"

	StoreKey = ModuleName

	RouterKey = ModuleName

	QuerierRoute = ModuleName

	// IBPrefix is the interest bearing token prefix
	IBPrefix = "ib_"
)

var (
	LendingPoolKey       = []byte{0x01}
	NextLendingPoolIDKey = []byte{0x02}
)

func GetLendingPoolKey(id uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, id)
	return append(LendingPoolKey, b...)
}

func GetNextLendingPoolKey() []byte {
	return NextLendingPoolIDKey
}
