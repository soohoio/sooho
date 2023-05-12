package types

// DONTCOVER

import errorsmod "cosmossdk.io/errors"

// x/claim module sentinel errors
var (
	ErrTotalWeightParse         = errorsmod.Register(ModuleName, 600, "total weight parse error")
	ErrFailedToGetTotalWeight   = errorsmod.Register(ModuleName, 601, "failed to get total weight")
	ErrAirdropAlreadyExists     = errorsmod.Register(ModuleName, 602, "airdrop with same identifier already exists")
	ErrDistributorAlreadyExists = errorsmod.Register(ModuleName, 603, "airdrop with same distributor already exists")
	ErrInvalidAmount            = errorsmod.Register(ModuleName, 604, "cannot claim negative tokens")
)
