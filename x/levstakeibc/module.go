package levstakeibc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/soohoio/stayking/v2/x/levstakeibc/client/cli"
	"github.com/soohoio/stayking/v2/x/levstakeibc/keeper"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

/**
 * AppModuleBasic implementations
 */

type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

func (a AppModuleBasic) Name() string {
	return types.ModuleName
}

func (AppModuleBasic) RegisterCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

func (a AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

func (a AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

func (a AppModuleBasic) DefaultGenesis(jsonCodec codec.JSONCodec) json.RawMessage {
	return jsonCodec.MustMarshalJSON(types.DefaultGenesis())
}

func (a AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientContext client.Context, mux *runtime.ServeMux) {
	err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientContext))
	if err != nil {
		panic(err)
	}
}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey)
}

/**
 * AppModule implementations
 */

type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
	}
}

func (a AppModule) Name() string {
	return a.AppModuleBasic.Name()
}

func (a AppModule) InitGenesis(context sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState

	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(context, a.keeper, genState)

	return []abci.ValidatorUpdate{}
}

func (a AppModule) ExportGenesis(context sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(context, a.keeper)
	return cdc.MustMarshalJSON(genState)
}

func (a AppModule) RegisterInvariants(registry sdk.InvariantRegistry) {
	keeper.RegisterInvariants(registry, a.keeper)
}

func (a AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.RouterKey, NewHandler(a.keeper))
}

func (a AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

func (a AppModule) LegacyQuerierHandler(amino *codec.LegacyAmino) sdk.Querier {
	return nil
}

func (a AppModule) RegisterServices(configurator module.Configurator) {
	types.RegisterMsgServer(configurator.MsgServer(), keeper.NewMsgServerImpl(a.keeper))
	types.RegisterQueryServer(configurator.QueryServer(), a.keeper)

}

func (a AppModule) ConsensusVersion() uint64 {
	return 1
}
