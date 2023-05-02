package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper interface {
	NewAccount(sdk.Context, authtypes.AccountI) authtypes.AccountI
	SetAccount(ctx sdk.Context, acc authtypes.AccountI)
	//GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	//GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI
	GetModuleAddress(moduleName string) sdk.AccAddress
}

type BankKeeper interface {
}

type AdminKeeper interface {
	IsAdmin(ctx sdk.Context, addr sdk.AccAddress) bool
}
