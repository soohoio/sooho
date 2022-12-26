package keeper

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	strideapp "github.com/soohoio/stayking/app"
	"github.com/soohoio/stayking/x/icacallbacks/keeper"
)

func IcacallbacksKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	app := strideapp.InitStayKingTestApp(true)
	icacallbackskeeper := app.IcacallbacksKeeper
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "stayking-1", Time: time.Now().UTC()})

	return &icacallbackskeeper, ctx
}
