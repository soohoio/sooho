package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	icacallbackstypes "github.com/soohoio/stayking/v2/x/icacallbacks/types"
)

const (
	ICACallbackID_Delegate = "delegate"
	//ICACallbackID_Claim      = "claim"
	//ICACallbackID_Undelegate = "undelegate"
	//ICACallbackID_Reinvest   = "reinvest"
	//ICACallbackID_Redemption = "redemption"
	//ICACallbackID_Rebalance  = "rebalance"
)

// ICACallbacks wrapper struct for stakeibc keeper
type ICACallback func(Keeper, sdk.Context, channeltypes.Packet, *icacallbackstypes.AcknowledgementResponse, []byte) error

type ICACallbacks struct {
	k            Keeper
	icacallbacks map[string]ICACallback
}

var _ icacallbackstypes.ICACallbackHandler = ICACallbacks{}

func (k Keeper) ICACallbackHandler() ICACallbacks {
	return ICACallbacks{k, make(map[string]ICACallback)}
}

func (c ICACallbacks) CallICACallback(ctx sdk.Context, id string, packet channeltypes.Packet, ackResponse *icacallbackstypes.AcknowledgementResponse, args []byte) error {
	return c.icacallbacks[id](c.k, ctx, packet, ackResponse, args)
}

func (c ICACallbacks) HasICACallback(id string) bool {
	_, found := c.icacallbacks[id]
	return found
}

func (c ICACallbacks) AddICACallback(id string, fn interface{}) icacallbackstypes.ICACallbackHandler {
	c.icacallbacks[id] = fn.(ICACallback)
	return c
}

func (c ICACallbacks) RegisterICACallbacks() icacallbackstypes.ICACallbackHandler {
	a := c.
		AddICACallback(ICACallbackID_Delegate, ICACallback(func(keeper Keeper, context sdk.Context, packet channeltypes.Packet, response *icacallbackstypes.AcknowledgementResponse, bytes []byte) error {
			return nil
		}))
	return a.(ICACallbacks)
}
