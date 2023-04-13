package keeper

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/v2/x/interchainquery/types"
)

// EndBlocker of interchainquery module
func (k Keeper) EndBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)
	_ = k.Logger(ctx)
	events := sdk.Events{}
	// emit events for periodic queries
	k.IterateQueries(ctx, func(_ int64, query types.Query) (stop bool) {
		if !query.RequestSent {
			k.Logger(ctx).Info(fmt.Sprintf("[ICQ REQUEST] Interchainquery event emitted %s", query.Id))
			event := sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
				sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeValueQuery),
				sdk.NewAttribute(types.AttributeKeyQueryId, query.Id),
				sdk.NewAttribute(types.AttributeKeyChainId, query.ChainId),
				sdk.NewAttribute(types.AttributeKeyConnectionId, query.ConnectionId),
				sdk.NewAttribute(types.AttributeKeyType, query.QueryType),
				sdk.NewAttribute(types.AttributeKeyHeight, "0"),
				sdk.NewAttribute(types.AttributeKeyRequest, hex.EncodeToString(query.Request)),
			)
			events = append(events, event)

			// 두 번을 이벤트 emit 하는데 이건 필요 없을듯? query type 이 store/bank/key 로 날리면 bank 의 balance 가 조회될듯..
			event.Type = "query_request"
			events = append(events, event)

			query.RequestSent = true
			k.SetQuery(ctx, query)
		}
		return false
	})

	if len(events) > 0 {
		ctx.EventManager().EmitEvents(events)
	}

}
