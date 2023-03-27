package lendingpool

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v2/x/lendingpool/keeper"
)

func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {

}

func Endblocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) {

}
