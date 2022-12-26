package types

import (
	"time"
)

var (
	DefaultClaimDenom                      = "ustay"
	DefaultAirdropDuration                 = time.Hour * 24 * 30 * 12 * 3 // 3 years
	DefaultVestingDurationForDelegateStake = time.Hour * 24 * 30 * 3      // 3 months
	DefaultVestingDurationForLiquidStake   = time.Hour * 24 * 30 * 3      // 3 months
	DefaultVestingInitialPeriod            = time.Hour * 24 * 30 * 3      // 3 months
	DefaultAirdropIdentifier               = "stayking"
)
