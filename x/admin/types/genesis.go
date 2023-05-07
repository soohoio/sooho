package types

func NewGenesisState(admins Admins) GenesisState {
	return GenesisState{Admins: &admins}
}

// using following definition
//var ad = map[string]bool{
//	"sooho13eda6gyezyukjarmm09l0h5x5t06h5ck8yh2js": true,
//	"sooho1ygs3em26qaheucpckxasxuqqej80sqt2p57nyy": true,
//	"sooho143umg272xger2eyurqfpjgt8u533s62mpz5weq": true,
//	"sooho1wal8dgs7whmykpdaz0chan2f54ynythkkcm264": true,  // delete
//	"sooho10d07y265gmmuvt4z0w9aw880jnsr700jefnezl": false, // gov module
//}

func DefaultGenesisState() *GenesisState {
	admins := Admins{
		Admins: []string{
			"sooho13eda6gyezyukjarmm09l0h5x5t06h5ck8yh2js",
			"sooho1ygs3em26qaheucpckxasxuqqej80sqt2p57nyy",
			"sooho143umg272xger2eyurqfpjgt8u533s62mpz5weq",
			"sooho1wal8dgs7whmykpdaz0chan2f54ynythkkcm264",
		},
	}
	return &GenesisState{
		Admins: &admins,
	}
}

func ValidateGenesis(data GenesisState) error {
	if len(data.Admins.Admins) == 0 {
		return ErrNoAdmins
	}
	return nil
}
