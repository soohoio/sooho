package tripleslope

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/gogo/protobuf/proto"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/soohoio/stayking/v3/x/lendingpool/types"
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
	cryptocodec.RegisterCrypto(amino)
}
