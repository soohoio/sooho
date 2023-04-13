package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gogo/protobuf/proto"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSubmitQueryResponse{}, "/stayking.interchainquery.v1.MsgSubmitQueryResponse", nil)
	cdc.RegisterConcrete(&MsgSendQueryBalance{}, "interchainquery/SendQueryBalance", nil)
	cdc.RegisterConcrete(&MsgSendEstimateSwapExactAmountOutResponse{}, "interchainquery/SendEstimateSwapExactAmountOutResponse", nil)

	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitQueryResponse{},
	)
	registry.RegisterImplementations((*proto.Message)(nil),
		&banktypes.QueryBalanceRequest{},
		&banktypes.QueryBalanceResponse{},
	)
	registry.RegisterImplementations((*proto.Message)(nil),
		&EstimateSwapExactAmountOutRequest{},
		&EstimateSwapExactAmountOutResponse{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)
	amino.Seal()
}
