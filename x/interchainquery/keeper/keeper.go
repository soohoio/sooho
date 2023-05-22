package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	recordsmodulekeeper "github.com/soohoio/stayking/v3/x/records/keeper"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v5/modules/core/keeper"
	"github.com/soohoio/stayking/v3/utils"
	"github.com/soohoio/stayking/v3/x/interchainquery/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of this module maintains collections of registered zones.
type Keeper struct {
	cdc        codec.Codec
	storeKey   storetypes.StoreKey
	memKey     storetypes.StoreKey
	callbacks  map[string]types.QueryCallbacks
	IBCKeeper  ibckeeper.Keeper
	paramstore paramtypes.Subspace

	ics4Wrapper   types.ICS4Wrapper
	channelKeeper types.ChannelKeeper
	portKeeper    types.PortKeeper
	scopedKeeper  capabilitykeeper.ScopedKeeper
	RecordsKeeper recordsmodulekeeper.Keeper
	AdminKeeper   types.AdminKeeper
}

// NewKeeper returns a new instance of zones Keeper
func NewKeeper(cdc codec.Codec,
	storeKey storetypes.StoreKey,
	memKey storetypes.StoreKey,
	ibckeeper ibckeeper.Keeper,
	ps paramtypes.Subspace,
	ics4Wrapper types.ICS4Wrapper,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	RecordsKeeper recordsmodulekeeper.Keeper) Keeper {

	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		callbacks:  make(map[string]types.QueryCallbacks),
		IBCKeeper:  ibckeeper,
		paramstore: ps,

		ics4Wrapper:   ics4Wrapper,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		scopedKeeper:  scopedKeeper,
		RecordsKeeper: RecordsKeeper,
	}
}

func (k *Keeper) SetCallbackHandler(module string, handler types.QueryCallbacks) error {
	_, found := k.callbacks[module]
	if found {
		return fmt.Errorf("callback handler already set for %s", module)
	}
	k.callbacks[module] = handler.RegisterICQCallbacks()
	return nil
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k *Keeper) MakeRequest(ctx sdk.Context, module string, callbackId string, chainId string, connectionId string, queryType string, request []byte, ttl uint64) error {
	k.Logger(ctx).Info(utils.LogWithHostZone(chainId,
		"Submitting ICQ Request - module=%s, callbackId=%s, connectionId=%s, queryType=%s, ttl=%d", module, callbackId, connectionId, queryType, ttl))

	// Confirm the connectionId and chainId are valid
	if connectionId == "" {
		errMsg := "[ICQ Validation Check] Failed! connection id cannot be empty"
		k.Logger(ctx).Error(errMsg)
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, errMsg)
	}
	if !strings.HasPrefix(connectionId, "connection") {
		errMsg := "[ICQ Validation Check] Failed! connection id must begin with 'connection'"
		k.Logger(ctx).Error(errMsg)
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, errMsg)
	}
	if chainId == "" {
		errMsg := "[ICQ Validation Check] Failed! chain_id cannot be empty"
		k.Logger(ctx).Error(errMsg)
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, errMsg)
	}

	// Confirm the module and callbackId exist
	if module != "" {
		if _, exists := k.callbacks[module]; !exists {
			err := fmt.Errorf("no callback handler registered for module %s", module)
			k.Logger(ctx).Error(err.Error())
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "no callback handler registered for module")
		}
		if exists := k.callbacks[module].HasICQCallback(callbackId); !exists {
			err := fmt.Errorf("no callback %s registered for module %s", callbackId, module)
			k.Logger(ctx).Error(err.Error())
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "no callback handler registered for module")
		}
	}

	// Save the query to the store
	// If the same query is re-requested, it will get replace in the store with an updated TTL
	//  and the RequestSent bool reset to false
	query := k.NewQuery(ctx, module, callbackId, chainId, connectionId, queryType, request, ttl)
	k.SetQuery(ctx, *query)

	return nil
}

// GetPort returns the portID for the transfer module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.PortKey))
}

// SetPort sets the portID for the transfer module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PortKey, []byte(portID))
}

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function
func (k Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	return k.scopedKeeper.AuthenticateCapability(ctx, cap, name)
}

// ClaimCapability wraps the scopedKeeper's ClaimCapability function
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

// BindPort stores the provided portID and binds to it, returning the associated capability
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.portKeeper.BindPort(ctx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// IsBound checks if the interchain query already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.scopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return ok
}
