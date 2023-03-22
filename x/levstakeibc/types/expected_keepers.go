package types

type AccountKeeper interface {
	//NewAccount(sdk.Context, authtypes.AccountI) authtypes.AccountI
	//SetAccount(ctx sdk.Context, acc authtypes.AccountI)
	//GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	//GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI
}

type BankKeeper interface {
}