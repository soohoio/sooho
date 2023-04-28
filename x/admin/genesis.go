package admin

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v2/x/admin/keeper"
	"github.com/soohoio/stayking/v2/x/admin/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	k.SetAdmins(ctx, *data.Admins)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	admins := k.GetAdmins(ctx)
	return types.GenesisState{
		Admins: &admins,
	}
}
