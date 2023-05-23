package types

import "encoding/binary"

const (
	ModuleName = "lendingpool"

	StoreKey = ModuleName

	RouterKey = ModuleName

	QuerierRoute = ModuleName

	// IBPrefix is the interest bearing token prefix
	IBPrefix = "ib"

	LendingPoolFeeAccount = "sooho1lsm9tzkh3jnzaglzqzzj4glzwcvy36jnadexgq"
)

var (
	LendingPoolKey       = []byte{0x01}
	NextLendingPoolIDKey = []byte{0x02}
	LoanKey              = []byte{0x03}
	NextLoanIDKey        = []byte{0x04}
)

func GetLendingPoolKey(id uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, id)
	return append(LendingPoolKey, b...)
}

func GetNextLendingPoolKey() []byte {
	return NextLendingPoolIDKey
}

func GetLoanKey(id uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, id)
	return append(LoanKey, b...)
}

func GetNextLoanKey() []byte {
	return NextLoanIDKey
}
