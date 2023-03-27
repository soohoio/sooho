package keeper

import "github.com/soohoio/stayking/v2/x/levstakeibc/types"

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
