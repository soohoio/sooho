package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/soohoio/stayking/v2/x/claim/vesting/exported"
)

// RegisterLegacyAminoCodec registers the vesting interfaces and concrete types on the
// provided LegacyAmino codec. These types are used for Amino JSON serialization
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*exported.VestingAccount)(nil), nil)
	cdc.RegisterConcrete(&StayKingPeriodicVestingAccount{}, "cosmos-sdk/StayKingPeriodicVestingAccount", nil)
}

// RegisterInterface associates protoName with AccountI and VestingAccount
// Interfaces and creates a registry of it's concrete implementations
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"stayking.vesting.v1beta1.VestingAccount",
		(*exported.VestingAccount)(nil),
		&StayKingPeriodicVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&BaseVestingAccount{},
		&StayKingPeriodicVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&BaseVestingAccount{},
		&StayKingPeriodicVestingAccount{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var amino = codec.NewLegacyAmino()

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}
