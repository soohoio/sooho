package app

import (
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govtypesv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ica "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts"
	icacontroller "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/types"
	icahost "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	ibcfeekeeper "github.com/cosmos/ibc-go/v5/modules/apps/29-fee/keeper"
	ibcfeetypes "github.com/cosmos/ibc-go/v5/modules/apps/29-fee/types"
	"github.com/cosmos/ibc-go/v5/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v5/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v5/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v5/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v5/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	porttypes "github.com/cosmos/ibc-go/v5/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v5/modules/core/keeper"
	ibctesting "github.com/cosmos/ibc-go/v5/testing"
	ibctestingtypes "github.com/cosmos/ibc-go/v5/testing/types"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	server "github.com/soohoio/stayking/v2/server"
	"github.com/soohoio/stayking/v2/utils"
	"github.com/soohoio/stayking/v2/x/claim"
	claimkeeper "github.com/soohoio/stayking/v2/x/claim/keeper"
	claimtypes "github.com/soohoio/stayking/v2/x/claim/types"
	claimvesting "github.com/soohoio/stayking/v2/x/claim/vesting"
	claimvestingtypes "github.com/soohoio/stayking/v2/x/claim/vesting/types"
	epochsmodule "github.com/soohoio/stayking/v2/x/epochs"
	epochsmodulekeeper "github.com/soohoio/stayking/v2/x/epochs/keeper"
	epochsmoduletypes "github.com/soohoio/stayking/v2/x/epochs/types"
	icacallbacksmodule "github.com/soohoio/stayking/v2/x/icacallbacks"
	icacallbacksmodulekeeper "github.com/soohoio/stayking/v2/x/icacallbacks/keeper"
	icacallbacksmoduletypes "github.com/soohoio/stayking/v2/x/icacallbacks/types"
	"github.com/soohoio/stayking/v2/x/interchainquery"
	interchainquerykeeper "github.com/soohoio/stayking/v2/x/interchainquery/keeper"
	interchainquerytypes "github.com/soohoio/stayking/v2/x/interchainquery/types"
	"github.com/soohoio/stayking/v2/x/lendingpool"
	lendingpoolkeeper "github.com/soohoio/stayking/v2/x/lendingpool/keeper"
	lendingpooltypes "github.com/soohoio/stayking/v2/x/lendingpool/types"
	"github.com/soohoio/stayking/v2/x/mint"
	mintkeeper "github.com/soohoio/stayking/v2/x/mint/keeper"
	minttypes "github.com/soohoio/stayking/v2/x/mint/types"
	"github.com/soohoio/stayking/v2/x/mockborrow"
	mockborrowkeeper "github.com/soohoio/stayking/v2/x/mockborrow/keeper"
	mockborrowtypes "github.com/soohoio/stayking/v2/x/mockborrow/types"
	recordsmodule "github.com/soohoio/stayking/v2/x/records"
	recordsmodulekeeper "github.com/soohoio/stayking/v2/x/records/keeper"
	recordsmoduletypes "github.com/soohoio/stayking/v2/x/records/types"

	// add levstakeibc module here like stakeibc
	levstakeibcmodule "github.com/soohoio/stayking/v2/x/levstakeibc"
	levstakeibcmodulekeeper "github.com/soohoio/stayking/v2/x/levstakeibc/keeper"
	levstakeibcmoduletypes "github.com/soohoio/stayking/v2/x/levstakeibc/types"
)

const (
	AccountAddressPrefix = "sooho"
	Name                 = "stayking"
	Version              = "2.0.0"
)

func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		distrclient.ProposalHandler,
		upgradeclient.LegacyProposalHandler,
		upgradeclient.LegacyCancelProposalHandler,
		ibcclientclient.UpdateClientProposalHandler,
		ibcclientclient.UpgradeProposalHandler,
		//levstakeibcclient.AddValidatorProposalHandler,
	)

	return govProposalHandlers
}

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(getGovProposalHandlers()),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		claimvesting.AppModuleBasic{},
		//stakeibcmodule.AppModuleBasic{},
		levstakeibcmodule.AppModuleBasic{},
		epochsmodule.AppModuleBasic{},
		interchainquery.AppModuleBasic{},
		ica.AppModuleBasic{},
		recordsmodule.AppModuleBasic{},
		icacallbacksmodule.AppModuleBasic{},
		claim.AppModuleBasic{},
		lendingpool.AppModuleBasic{},
		mockborrow.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter, authtypes.Burner},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		//stakeibcmoduletypes.ModuleName:  {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		levstakeibcmoduletypes.ModuleName: {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		claimtypes.ModuleName:             nil,
		interchainquerytypes.ModuleName:   nil,
		icatypes.ModuleName:               nil,
		lendingpooltypes.ModuleName:       {authtypes.Minter, authtypes.Burner},
		mockborrowtypes.ModuleName:        nil,
	}
)

var (
	_ servertypes.Application = (*StayKingApp)(nil)
	_ ibctesting.TestingApp   = (*StayKingApp)(nil)
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, "."+Name)
}

// StayKingApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type StayKingApp struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	MintKeeper       mintkeeper.Keeper
	DistrKeeper      distrkeeper.Keeper
	GovKeeper        govkeeper.Keeper
	CrisisKeeper     crisiskeeper.Keeper
	UpgradeKeeper    upgradekeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper   evidencekeeper.Keeper
	TransferKeeper   ibctransferkeeper.Keeper
	FeeGrantKeeper   feegrantkeeper.Keeper
	// MonitoringKeeper    monitoringpkeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper
	// ScopedMonitoringKeeper capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper

	//ScopedStakeibcKeeper capabilitykeeper.ScopedKeeper
	//StakeibcKeeper       stakeibcmodulekeeper.Keeper

	ScopedLevstakeibcKeeper capabilitykeeper.ScopedKeeper
	LevstakeibcKeeper       levstakeibcmodulekeeper.Keeper

	EpochsKeeper                epochsmodulekeeper.Keeper
	InterchainqueryKeeper       interchainquerykeeper.Keeper
	ScopedInterchainqueryKeeper capabilitykeeper.ScopedKeeper
	ScopedRecordsKeeper         capabilitykeeper.ScopedKeeper
	RecordsKeeper               recordsmodulekeeper.Keeper
	ScopedIcacallbacksKeeper    capabilitykeeper.ScopedKeeper
	IcacallbacksKeeper          icacallbacksmodulekeeper.Keeper
	ScopedratelimitKeeper       capabilitykeeper.ScopedKeeper
	ClaimKeeper                 claimkeeper.Keeper
	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	LendingPoolKeeper lendingpoolkeeper.Keeper
	MockBorrowKeeper  mockborrowkeeper.Keeper

	mm           *module.Manager
	sm           *module.SimulationManager
	configurator module.Configurator
}

// RUN GOSEC
// New returns a reference to an initialized blockchain app
func NewStayKingApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *StayKingApp {
	appCodec := encodingConfig.Marshaler
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(Name, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey,
		banktypes.StoreKey,
		stakingtypes.StoreKey,
		minttypes.StoreKey,
		distrtypes.StoreKey,
		slashingtypes.StoreKey,
		govtypes.StoreKey,
		paramstypes.StoreKey,
		ibchost.StoreKey,
		upgradetypes.StoreKey,
		feegrant.StoreKey,
		evidencetypes.StoreKey,
		ibctransfertypes.StoreKey,
		capabilitytypes.StoreKey,
		//stakeibcmoduletypes.StoreKey,
		levstakeibcmoduletypes.StoreKey,
		epochsmoduletypes.StoreKey,
		interchainquerytypes.StoreKey,
		icacontrollertypes.StoreKey,
		icahosttypes.StoreKey,
		recordsmoduletypes.StoreKey,
		icacallbacksmoduletypes.StoreKey,
		claimtypes.StoreKey,
		lendingpooltypes.StoreKey,
		mockborrowtypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(
		capabilitytypes.MemStoreKey,
		// stakeibcmoduletypes.MemStoreKey,
		// levstakeibcmoduletypes.MemStoreKey,
		// icacallbacksmoduletypes.MemStoreKey,
		// recordsmoduletypes.MemStoreKey,
		interchainquerytypes.MemStoreKey,
	)

	app := &StayKingApp{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// set the BaseApp's parameter store
	bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable()))

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedICAControllerKeeper := app.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
	scopedICAHostKeeper := app.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey], app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms, AccountAddressPrefix,
	)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec, keys[banktypes.StoreKey], app.AccountKeeper, app.GetSubspace(banktypes.ModuleName), app.BlacklistedModuleAccountAddrs(),
	)
	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec, keys[stakingtypes.StoreKey], app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName),
	)
	epochsKeeper := epochsmodulekeeper.NewKeeper(appCodec, keys[epochsmoduletypes.StoreKey])
	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec, keys[minttypes.StoreKey], app.GetSubspace(minttypes.ModuleName), app.AccountKeeper, app.BankKeeper, app.DistrKeeper, app.EpochsKeeper, authtypes.FeeCollectorName,
	)
	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec, keys[distrtypes.StoreKey], app.GetSubspace(distrtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, authtypes.FeeCollectorName,
	)
	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec, keys[slashingtypes.StoreKey], &stakingKeeper, app.GetSubspace(slashingtypes.ModuleName),
	)
	app.CrisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName), invCheckPeriod, app.BankKeeper, authtypes.FeeCollectorName,
	)

	app.ClaimKeeper = *claimkeeper.NewKeeper(
		appCodec,
		keys[claimtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper, app.StakingKeeper, app.DistrKeeper, epochsKeeper)

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], app.AccountKeeper)
	app.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, app.BaseApp, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks(), app.ClaimKeeper.Hooks()),
	)

	// ... other modules keepers

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, keys[ibchost.StoreKey], app.GetSubspace(ibchost.ModuleName), app.StakingKeeper, app.UpgradeKeeper, scopedIBCKeeper,
	)

	// Create Transfer Keepers
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec, keys[ibctransfertypes.StoreKey], app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper, app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
		app.AccountKeeper, app.BankKeeper, scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(app.TransferKeeper)
	transferIBCModule := transfer.NewIBCModule(app.TransferKeeper)

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, keys[evidencetypes.StoreKey], &app.StakingKeeper, app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper

	// Note: must be above app.StakeibcKeeper
	app.ICAControllerKeeper = icacontrollerkeeper.NewKeeper(
		appCodec, keys[icacontrollertypes.StoreKey], app.GetSubspace(icacontrollertypes.SubModuleName),
		app.IBCKeeper.ChannelKeeper, // may be replaced with middleware such as ics29 fee
		app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
		scopedICAControllerKeeper, app.MsgServiceRouter(),
	)

	scopedIcacallbacksKeeper := app.CapabilityKeeper.ScopeToModule(icacallbacksmoduletypes.ModuleName)
	app.ScopedIcacallbacksKeeper = scopedIcacallbacksKeeper
	app.IcacallbacksKeeper = *icacallbacksmodulekeeper.NewKeeper(
		appCodec,
		keys[icacallbacksmoduletypes.StoreKey],
		keys[icacallbacksmoduletypes.MemStoreKey], // TODO. keys map 에 MemStoreKey 가 없는뎅?
		app.GetSubspace(icacallbacksmoduletypes.ModuleName),
		scopedIcacallbacksKeeper,
		*app.IBCKeeper,
		app.ICAControllerKeeper,
	)

	scopedRecordsKeeper := app.CapabilityKeeper.ScopeToModule(recordsmoduletypes.ModuleName)
	app.ScopedRecordsKeeper = scopedRecordsKeeper

	app.RecordsKeeper = *recordsmodulekeeper.NewKeeper(
		appCodec,
		keys[recordsmoduletypes.StoreKey],
		keys[recordsmoduletypes.MemStoreKey],
		app.GetSubspace(recordsmoduletypes.ModuleName),
		scopedRecordsKeeper,
		app.AccountKeeper,
		app.TransferKeeper,
		*app.IBCKeeper,
		app.IcacallbacksKeeper,
	)
	recordsModule := recordsmodule.NewAppModule(appCodec, app.RecordsKeeper, app.AccountKeeper, app.BankKeeper)

	//scopedStakeibcKeeper := app.CapabilityKeeper.ScopeToModule(stakeibcmoduletypes.ModuleName)
	//app.ScopedStakeibcKeeper = scopedStakeibcKeeper
	//stakeibcKeeper := stakeibcmodulekeeper.NewKeeper(
	//	appCodec,
	//	keys[stakeibcmoduletypes.StoreKey],
	//	keys[stakeibcmoduletypes.MemStoreKey],
	//	app.GetSubspace(stakeibcmoduletypes.ModuleName),
	//	// app.IBCKeeper.ChannelKeeper,
	//	// &app.IBCKeeper.PortKeeper,
	//	app.AccountKeeper,
	//	app.BankKeeper,
	//	app.ICAControllerKeeper,
	//	*app.IBCKeeper,
	//	scopedStakeibcKeeper,
	//	app.InterchainqueryKeeper,
	//	app.RecordsKeeper,
	//	app.StakingKeeper,
	//	app.IcacallbacksKeeper,
	//)
	//app.StakeibcKeeper = *stakeibcKeeper.SetHooks(
	//	stakeibcmoduletypes.NewMultiStakeIBCHooks(app.ClaimKeeper.Hooks()),
	//)

	//stakeibcModule := stakeibcmodule.NewAppModule(appCodec, app.StakeibcKeeper, app.AccountKeeper, app.BankKeeper)
	//stakeibcIBCModule := stakeibcmodule.NewIBCModule(app.StakeibcKeeper)

	app.LendingPoolKeeper = lendingpoolkeeper.NewKeeper(appCodec, keys[lendingpooltypes.StoreKey],
		app.GetSubspace(lendingpooltypes.ModuleName), app.AccountKeeper, app.BankKeeper)

	app.MockBorrowKeeper = mockborrowkeeper.NewKeeper(appCodec, keys[mockborrowtypes.StoreKey],
		app.AccountKeeper, app.BankKeeper, app.LendingPoolKeeper)

	scopedInterchainqueryKeeper := app.CapabilityKeeper.ScopeToModule(interchainquerytypes.ModuleName)
	app.ScopedInterchainqueryKeeper = scopedInterchainqueryKeeper
	app.InterchainqueryKeeper = interchainquerykeeper.NewKeeper(
		appCodec,
		keys[interchainquerytypes.StoreKey],
		keys[interchainquerytypes.MemStoreKey],
		*app.IBCKeeper,
		app.GetSubspace(interchainquerytypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedInterchainqueryKeeper,
		app.RecordsKeeper,
	)

	// levstakeibc module setup
	scopedLevstakeibcKeeper := app.CapabilityKeeper.ScopeToModule(levstakeibcmoduletypes.ModuleName)
	app.LevstakeibcKeeper = levstakeibcmodulekeeper.NewKeeper(
		appCodec,
		keys[levstakeibcmoduletypes.StoreKey],
		keys[levstakeibcmoduletypes.MemStoreKey],
		app.GetSubspace(levstakeibcmoduletypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		scopedLevstakeibcKeeper,
		app.InterchainqueryKeeper,
		app.StakingKeeper,
		*app.IBCKeeper,
		app.ICAControllerKeeper,
		app.IcacallbacksKeeper,
		app.RecordsKeeper,
		app.LendingPoolKeeper,
	)
	//app.LevstakeibcKeeper = levstakeibcKeeper
	levstakeModule := levstakeibcmodule.NewAppModule(appCodec, app.LevstakeibcKeeper, app.AccountKeeper, app.BankKeeper)
	levstakeibcIBCModule := levstakeibcmodule.NewIBCModule(app.LevstakeibcKeeper)

	// Register ICQ callbacks
	err := app.InterchainqueryKeeper.SetCallbackHandler(levstakeibcmoduletypes.ModuleName, app.LevstakeibcKeeper.ICQCallbackHandler())
	if err != nil {
		return nil
	}

	interchainQueryModule := interchainquery.NewAppModule(appCodec, app.InterchainqueryKeeper)
	interchainQueryIBCModule := interchainquery.NewIBCModule(app.InterchainqueryKeeper)

	// Register Gov (must be registerd after stakeibc)
	govRouter := govtypesv1beta1.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypesv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper)).
		AddRoute(levstakeibcmoduletypes.RouterKey, levstakeibcmodule.NewLevStakeibcProposalHandler(app.LevstakeibcKeeper))

	app.GovKeeper = govkeeper.NewKeeper(
		appCodec, keys[govtypes.StoreKey], app.GetSubspace(govtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, govRouter, app.MsgServiceRouter(), govtypes.DefaultConfig(),
	)

	app.EpochsKeeper = *epochsKeeper.SetHooks(
		epochsmoduletypes.NewMultiEpochHooks(
			app.LevstakeibcKeeper.Hooks(),
			app.MintKeeper.Hooks(),
			app.ClaimKeeper.Hooks(),
			app.InterchainqueryKeeper.Hooks(),
		),
	)
	epochsModule := epochsmodule.NewAppModule(appCodec, app.EpochsKeeper)

	icacallbacksModule := icacallbacksmodule.NewAppModule(appCodec, app.IcacallbacksKeeper, app.AccountKeeper, app.BankKeeper)
	// Register ICA calllbacks
	// stakeibc
	err = app.IcacallbacksKeeper.SetICACallbackHandler(levstakeibcmoduletypes.ModuleName, app.LevstakeibcKeeper.ICACallbackHandler())
	if err != nil {
		return nil
	}
	//// records
	err = app.IcacallbacksKeeper.SetICACallbackHandler(recordsmoduletypes.ModuleName, app.RecordsKeeper.ICACallbackHandler())
	if err != nil {
		return nil
	}

	// this line is used by starport scaffolding # stargate/app/keeperDefinition
	ibcFeeKeeper := ibcfeekeeper.NewKeeper(
		appCodec, app.keys[ibcfeetypes.StoreKey], app.GetSubspace(ibcfeetypes.ModuleName),
		app.IBCKeeper.ChannelKeeper, // may be replaced with IBC middleware
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper, app.AccountKeeper, app.BankKeeper,
	)

	// create IBC middleware stacks by combining middleware with base application
	app.ICAHostKeeper = icahostkeeper.NewKeeper(
		appCodec, keys[icahosttypes.StoreKey], app.GetSubspace(icahosttypes.SubModuleName),
		ibcFeeKeeper,
		app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
		app.AccountKeeper, scopedICAHostKeeper, app.MsgServiceRouter())
	icaModule := ica.NewAppModule(&app.ICAControllerKeeper, &app.ICAHostKeeper)

	// Create the middleware stacks

	//icaCallbackForStakeIBCModule := icacallbacksmodule.NewIBCModule(app.IcacallbacksKeeper, stakeibcIBCModule)
	//icaMiddlewareForStakeIBCStack = icacallbacksmodule.NewIBCModule(app.IcacallbacksKeeper, levstakeibcIBCModule)
	//icaMiddlewareForStakeIBCStack := icacontroller.NewIBCMiddleware(icaCallbackForStakeIBCModule, app.ICAControllerKeeper)

	icaCallbackForLevStakeIBCModule := icacallbacksmodule.NewIBCModule(app.IcacallbacksKeeper, levstakeibcIBCModule)
	icaMiddlewareForLevStakeIBCStack := icacontroller.NewIBCMiddleware(icaCallbackForLevStakeIBCModule, app.ICAControllerKeeper)

	icaHostIBCModule := icahost.NewIBCModule(app.ICAHostKeeper)

	// Stack two contains
	// - IBC
	// - records
	// - transfer
	// - base app
	recordsStack := recordsmodule.NewIBCModule(app.RecordsKeeper, transferIBCModule)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := porttypes.NewRouter()
	ibcRouter.
		AddRoute(ibctransfertypes.ModuleName, recordsStack).
		AddRoute(icacontrollertypes.SubModuleName, icaMiddlewareForLevStakeIBCStack).
		AddRoute(icahosttypes.SubModuleName, icaHostIBCModule).
		AddRoute(interchainquerytypes.ModuleName, interchainQueryIBCModule).

		// this line is used by starport scaffolding # ibc/app/router

		//AddRoute(icacontrollertypes.SubModuleName, icaMiddlewareForStakeIBCStack).
		//AddRoute(stakeibcmoduletypes.ModuleName, icaMiddlewareForStakeIBCStack).
		//AddRoute(icacallbacksmoduletypes.ModuleName, icaMiddlewareForStakeIBCStack).
		AddRoute(levstakeibcmoduletypes.ModuleName, icaMiddlewareForLevStakeIBCStack).
		AddRoute(icacallbacksmoduletypes.ModuleName, icaMiddlewareForLevStakeIBCStack)

	//.
	//	AddRoute(icacallbacksmoduletypes.ModuleName, icaMiddlewareForLevStakeIBCStack)

	app.IBCKeeper.SetRouter(ibcRouter)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.

	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		claimvesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, app.BankKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		params.NewAppModule(app.ParamsKeeper),
		claim.NewAppModule(appCodec, app.ClaimKeeper),
		transferModule,
		// monitoringModule,
		//stakeibcModule,
		levstakeModule,
		epochsModule,
		interchainQueryModule,
		icaModule,
		recordsModule,
		icacallbacksModule,
		lendingpool.NewAppModule(appCodec, app.LendingPoolKeeper, app.AccountKeeper, app.BankKeeper),
		mockborrow.NewAppModule(appCodec, app.MockBorrowKeeper, app.AccountKeeper, app.BankKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		vestingtypes.ModuleName,
		claimvestingtypes.ModuleName,
		ibchost.ModuleName,
		ibctransfertypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		// monitoringptypes.ModuleName,
		icatypes.ModuleName,
		//stakeibcmoduletypes.ModuleName,
		levstakeibcmoduletypes.ModuleName,
		epochsmoduletypes.ModuleName,
		interchainquerytypes.ModuleName,
		recordsmoduletypes.ModuleName,
		icacallbacksmoduletypes.ModuleName,
		claimtypes.ModuleName,
		lendingpooltypes.ModuleName,
		mockborrowtypes.ModuleName,
	)

	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		vestingtypes.ModuleName,
		claimvestingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		ibchost.ModuleName,
		ibctransfertypes.ModuleName,
		// monitoringptypes.ModuleName,
		icatypes.ModuleName,
		//stakeibcmoduletypes.ModuleName,
		levstakeibcmoduletypes.ModuleName,
		epochsmoduletypes.ModuleName,
		interchainquerytypes.ModuleName,
		recordsmoduletypes.ModuleName,
		icacallbacksmoduletypes.ModuleName,
		claimtypes.ModuleName,
		lendingpooltypes.ModuleName,
		mockborrowtypes.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		vestingtypes.ModuleName,
		claimvestingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		ibchost.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		ibctransfertypes.ModuleName,
		feegrant.ModuleName,
		// monitoringptypes.ModuleName,
		icatypes.ModuleName,
		//stakeibcmoduletypes.ModuleName,
		levstakeibcmoduletypes.ModuleName,
		epochsmoduletypes.ModuleName,
		interchainquerytypes.ModuleName,
		recordsmoduletypes.ModuleName,
		icacallbacksmoduletypes.ModuleName,
		claimtypes.ModuleName,
		lendingpooltypes.ModuleName,
		mockborrowtypes.ModuleName,
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)
	app.setupUpgradeHandlers()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)

	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   app.AccountKeeper,
			BankKeeper:      app.BankKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			FeegrantKeeper:  app.FeeGrantKeeper,
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)
	if err != nil {
		panic(err)
	}

	app.SetAnteHandler(anteHandler)
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper
	app.ScopedICAControllerKeeper = scopedICAControllerKeeper
	app.ScopedICAHostKeeper = scopedICAHostKeeper
	//app.ScopedStakeibcKeeper = scopedStakeibcKeeper
	app.ScopedLevstakeibcKeeper = scopedLevstakeibcKeeper

	return app
}

// Name returns the name of the App
func (app *StayKingApp) Name() string { return app.BaseApp.Name() }

// GetBaseApp returns the base app of the application
func (app *StayKingApp) GetBaseApp() *baseapp.BaseApp { return app.BaseApp }

// GetStakingKeeper implements the TestingApp interface.
func (app *StayKingApp) GetStakingKeeper() ibctestingtypes.StakingKeeper {
	return app.StakingKeeper
}

// GetIBCKeeper implements the TestingApp interface.
func (app *StayKingApp) GetTransferKeeper() *ibctransferkeeper.Keeper {
	return &app.TransferKeeper
}

// GetIBCKeeper implements the TestingApp interface.
func (app *StayKingApp) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetScopedIBCKeeper implements the TestingApp interface.
func (app *StayKingApp) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

// GetTxConfig implements the TestingApp interface.
func (app *StayKingApp) GetTxConfig() client.TxConfig {
	cfg := MakeEncodingConfig()
	return cfg.TxConfig
}

// BeginBlocker application updates every begin block
func (app *StayKingApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *StayKingApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *StayKingApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *StayKingApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *StayKingApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	// DO NOT REMOVE: StringMapKeys fixes non-deterministic map iteration
	for _, acc := range utils.StringMapKeys(maccPerms) {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *StayKingApp) BlacklistedModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	// DO NOT REMOVE: StringMapKeys fixes non-deterministic map iteration
	for _, acc := range utils.StringMapKeys(maccPerms) {
		// don't blacklist stakeibc module account, so that it can ibc transfer tokens
		if acc == "levstakeibc" {
			continue
		}
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *StayKingApp) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns an app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *StayKingApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns an InterfaceRegistry
func (app *StayKingApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *StayKingApp) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *StayKingApp) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *StayKingApp) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *StayKingApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *StayKingApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if err := server.RegisterSwaggerAPI(apiSvr.ClientCtx, apiSvr.Router, apiConfig.Swagger); err != nil {
		panic(err)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *StayKingApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *StayKingApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(clientCtx, app.BaseApp.GRPCQueryRouter(), app.interfaceRegistry, app.Query)
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypesv1.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	// paramsKeeper.Subspace(monitoringptypes.ModuleName)
	//paramsKeeper.Subspace(stakeibcmoduletypes.ModuleName)
	paramsKeeper.Subspace(levstakeibcmoduletypes.ModuleName)
	paramsKeeper.Subspace(epochsmoduletypes.ModuleName)
	paramsKeeper.Subspace(interchainquerytypes.ModuleName)
	paramsKeeper.Subspace(icacontrollertypes.SubModuleName)
	paramsKeeper.Subspace(icahosttypes.SubModuleName)
	paramsKeeper.Subspace(recordsmoduletypes.ModuleName)
	paramsKeeper.Subspace(icacallbacksmoduletypes.ModuleName)
	// this line is used by starport scaffolding # stargate/app/paramSubspace
	paramsKeeper.Subspace(lendingpooltypes.ModuleName)

	paramsKeeper.Subspace(claimtypes.ModuleName)
	return paramsKeeper
}

// SimulationManager implements the SimulationApp interface
func (app *StayKingApp) SimulationManager() *module.SimulationManager {
	return app.sm
}
