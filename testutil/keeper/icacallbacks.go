package keeper

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	staykingApp "github.com/soohoio/stayking/v3/app"
	"github.com/soohoio/stayking/v3/x/icacallbacks/keeper"
)

func IcacallbacksKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	app := staykingApp.InitStayKingTestApp(true)
	icacallbackskeeper := app.IcacallbacksKeeper
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "stayking-1", Time: time.Now().UTC()})

	return &icacallbackskeeper, ctx
}
