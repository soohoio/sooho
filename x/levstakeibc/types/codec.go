package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gogo/protobuf/proto"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterHostZone{}, "levstakeibc/RegisterHostZone", nil)
	cdc.RegisterConcrete(&MsgSendQueryAllBalances{}, "levstakeibc/SendQueryAllBalances", nil)
	cdc.RegisterConcrete(&SendEstimateSwapExactAmountOutRequest{}, "levstakeibc/SendEstimateSwapExactAmountOutRequest", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgRegisterHostZone{},&MsgSendQueryAllBalances{})
	registry.RegisterImplementations((*proto.Message)(nil),
		&banktypes.QueryAllBalancesRequest{},
		&banktypes.QueryAllBalancesResponse{},
	)
	registry.RegisterImplementations((*proto.Message)(nil),
		&EstimateSwapExactAmountOutRequest{},
		&EstimateSwapExactAmountOutResponse{},
	)
	//registry.RegisterImplementations()

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
}
