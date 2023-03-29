package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v5/modules/core/keeper"
	ibctmtypes "github.com/cosmos/ibc-go/v5/modules/light-clients/07-tendermint/types"
	icacallbackskeeper "github.com/soohoio/stayking/v2/x/icacallbacks/keeper"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	recordsmodulekeeper "github.com/soohoio/stayking/v2/x/records/keeper"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	cdc                 codec.BinaryCodec
	storeKey            storetypes.StoreKey
	memKey              storetypes.StoreKey
	paramstore          paramtypes.Subspace
	accountKeeper       types.AccountKeeper
	bankKeeper          bankkeeper.Keeper
	scopedKeeper        capabilitykeeper.ScopedKeeper
	StakingKeeper       stakingkeeper.Keeper
	IBCKeeper           ibckeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICACallbacksKeeper  icacallbackskeeper.Keeper
	RecordsKeeper       recordsmodulekeeper.Keeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	StakingKeeper stakingkeeper.Keeper,
	ibcKeeper ibckeeper.Keeper,
	icaControllerKeeper icacontrollerkeeper.Keeper,
	icaCallbacksKeeper icacallbackskeeper.Keeper,
	recordsKeeper recordsmodulekeeper.Keeper,

) Keeper {
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:                 cdc,
		storeKey:            storeKey,
		memKey:              memKey,
		paramstore:          ps,
		accountKeeper:       accountKeeper,
		bankKeeper:          bankKeeper,
		scopedKeeper:        scopedKeeper,
		StakingKeeper:       StakingKeeper,
		IBCKeeper:           ibcKeeper,
		ICAControllerKeeper: icaControllerKeeper,
		ICACallbacksKeeper:  icaCallbacksKeeper,
		RecordsKeeper:       recordsKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) ClaimCapability(ctx sdk.Context, channelCap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, channelCap, name)
}

func (k Keeper) GetConnectionId(ctx sdk.Context, portId string) (string, error) {
	interchainAccounts := k.ICAControllerKeeper.GetAllInterchainAccounts(ctx)

	for _, interchainAccount := range interchainAccounts {
		if interchainAccount.PortId == portId {
			return interchainAccount.ConnectionId, nil
		}
	}

	return "", fmt.Errorf("PortId %s has no associated connectionId", portId)
}

func (k Keeper) GetChainID(ctx sdk.Context, connectionId string) (string, error) {
	conn, found := k.IBCKeeper.ConnectionKeeper.GetConnection(ctx, connectionId)
	if !found {
		errMsg := fmt.Sprintf("invalid connection id, %s not found", connectionId)
		k.Logger(ctx).Error(errMsg)
		return "", fmt.Errorf(errMsg)
	}
	clientState, found := k.IBCKeeper.ClientKeeper.GetClientState(ctx, conn.ClientId)
	if !found {
		errMsg := fmt.Sprintf("client id %s not found for connection %s", conn.ClientId, connectionId)
		k.Logger(ctx).Error(errMsg)
		return "", fmt.Errorf(errMsg)
	}
	client, ok := clientState.(*ibctmtypes.ClientState)
	if !ok {
		errMsg := fmt.Sprintf("invalid client state for client %s on connection %s", conn.ClientId, connectionId)
		k.Logger(ctx).Error(errMsg)
		return "", fmt.Errorf(errMsg)
	}

	return client.ChainId, nil
}
