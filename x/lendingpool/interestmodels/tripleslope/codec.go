package tripleslope

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/gogo/protobuf/proto"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/soohoio/stayking/v2/x/lendingpool/types"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"stayking.lendingpool.v1.InterestModelI",
		(*types.InterestModelI)(nil),
		&TripleSlope{},
	)
	registry.RegisterImplementations(
		(*proto.Message)(nil),
		&Params{},
	)
}

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	amino.Seal()
}
