package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/soohoio/stayking/v2/utils"
	"github.com/soohoio/stayking/v2/x/levstakeibc/types"
	recordstypes "github.com/soohoio/stayking/v2/x/records/types"
)

func (k Keeper) SetWithdrawalAddress(ctx sdk.Context) {
	k.Logger(ctx).Info("Setting Withdrawal Addresses...")

	for _, hostZone := range k.GetAllHostZone(ctx) {
		err := k.SetWithdrawalAddressOnHost(ctx, hostZone)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Unable to set withdrawal address on %s, err: %s", hostZone.ChainId, err))
		}
	}
}

// UpdateRedemptionRates redemption rate = (Unbonded Balance + Staked Balance + Module Account Balance) / (stToken Supply)
func (k Keeper) UpdateRedemptionRates(ctx sdk.Context, depositRecords []recordstypes.DepositRecord) {
	k.Logger(ctx).Info("Updating Redemption Rates...")

	// Update the redemption rate for each host zone
	for _, hostZone := range k.GetAllHostZone(ctx) {

		// Gather redemption rate components
		stSupply := k.bankKeeper.GetSupply(ctx, types.StAssetDenomFromHostZoneDenom(hostZone.HostDenom)).Amount
		if stSupply.IsZero() {
			k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "No st%s in circulation - redemption rate is unchanged", hostZone.HostDenom))
			continue
		}
		undelegatedBalance, err := k.GetUndelegatedBalance(hostZone, depositRecords)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Could not get undelegated balance for host zone %s: %s", hostZone.ChainId, err.Error()))
			return
		}
		stakedBalance := hostZone.StakedBal
		moduleAcctBalance, err := k.GetModuleAccountBalance(hostZone, depositRecords)
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("Could not get module account balance for host zone %s: %s", hostZone.ChainId, err.Error()))
			return
		}

		k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId,
			"Redemption Rate Components - Undelegated Balance: %v, Staked Balance: %v, Module Account Balance: %v, stToken Supply: %v",
			undelegatedBalance, stakedBalance, moduleAcctBalance, stSupply))

		// Calculate the redemption rate
		redemptionRate := (sdk.NewDecFromInt(undelegatedBalance).Add(sdk.NewDecFromInt(stakedBalance)).Add(sdk.NewDecFromInt(moduleAcctBalance))).Quo(sdk.NewDecFromInt(stSupply))
		k.Logger(ctx).Info(utils.LogWithHostZone(hostZone.ChainId, "New Redemption Rate: %v (vs Prev Rate: %v)", redemptionRate, hostZone.RedemptionRate))

		// Update the host zone
		hostZone.LastRedemptionRate = hostZone.RedemptionRate
		hostZone.RedemptionRate = redemptionRate
		k.SetHostZone(ctx, hostZone)
	}
}

func (k Keeper) GetUndelegatedBalance(hostZone types.HostZone, depositRecords []recordstypes.DepositRecord) (sdk.Int, error) {
	// filter to only the deposit records for the host zone with status DELEGATION_QUEUE
	UndelegatedDepositRecords := utils.FilterDepositRecords(depositRecords, func(record recordstypes.DepositRecord) (condition bool) {
		return (record.Status == recordstypes.DepositRecord_DELEGATION_QUEUE || record.Status == recordstypes.DepositRecord_DELEGATION_IN_PROGRESS) && record.HostZoneId == hostZone.ChainId
	})

	// sum the amounts of the deposit records
	totalAmount := sdk.ZeroInt()
	for _, depositRecord := range UndelegatedDepositRecords {
		totalAmount = totalAmount.Add(depositRecord.Amount)
	}

	return totalAmount, nil
}

func (k Keeper) GetModuleAccountBalance(hostZone types.HostZone, depositRecords []recordstypes.DepositRecord) (sdk.Int, error) {
	// filter to only the deposit records for the host zone with status DELEGATION
	ModuleAccountRecords := utils.FilterDepositRecords(depositRecords, func(record recordstypes.DepositRecord) (condition bool) {
		return (record.Status == recordstypes.DepositRecord_TRANSFER_QUEUE || record.Status == recordstypes.DepositRecord_TRANSFER_IN_PROGRESS) && record.HostZoneId == hostZone.ChainId
	})
	// sum the amounts of the deposit records
	totalAmount := sdk.ZeroInt()
	for _, depositRecord := range ModuleAccountRecords {
		totalAmount = totalAmount.Add(depositRecord.Amount)
	}

	return totalAmount, nil
}
