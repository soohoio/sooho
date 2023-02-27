package v2

import (
	"github.com/cosmos/cosmos-sdk/types"
	oldrecordstypes "github.com/soohoio/stayking/v2/x/records/migrations/v2/types"
	recordstypes "github.com/soohoio/stayking/v2/x/records/types"
)

func convertToNewDepositRecord(oldDepositRecord oldrecordstypes.DepositRecord) recordstypes.DepositRecord {
	return recordstypes.DepositRecord{
		Id:                 oldDepositRecord.Id,
		Amount:             types.Int(oldDepositRecord.Amount),
		Denom:              oldDepositRecord.Denom,
		HostZoneId:         oldDepositRecord.HostZoneId,
		Status:             recordstypes.DepositRecord_Status(oldDepositRecord.Status),
		DepositEpochNumber: oldDepositRecord.DepositEpochNumber,
		Source:             recordstypes.DepositRecord_Source(oldDepositRecord.Source),
	}
}

func convertToNewHostZoneUnbonding(oldHostZoneUnbondings oldrecordstypes.HostZoneUnbonding) recordstypes.HostZoneUnbonding {
	return recordstypes.HostZoneUnbonding{
		StTokenAmount:         types.Int(oldHostZoneUnbondings.StTokenAmount),
		NativeTokenAmount:     types.Int(oldHostZoneUnbondings.NativeTokenAmount),
		Denom:                 oldHostZoneUnbondings.Denom,
		HostZoneId:            oldHostZoneUnbondings.HostZoneId,
		UnbondingTime:         oldHostZoneUnbondings.UnbondingTime,
		Status:                recordstypes.HostZoneUnbonding_Status(oldHostZoneUnbondings.Status),
		UserRedemptionRecords: oldHostZoneUnbondings.UserRedemptionRecords,
	}
}

func convertToNewEpochUnbondingRecord(oldEpochUnbondingRecord oldrecordstypes.EpochUnbondingRecord) recordstypes.EpochUnbondingRecord {
	var epochUnbondingRecord recordstypes.EpochUnbondingRecord
	for _, oldHostZoneUnbonding := range oldEpochUnbondingRecord.HostZoneUnbondings {
		newHostZoneUnbonding := convertToNewHostZoneUnbonding(*oldHostZoneUnbonding)
		epochUnbondingRecord.HostZoneUnbondings = append(epochUnbondingRecord.HostZoneUnbondings, &newHostZoneUnbonding)
	}
	return epochUnbondingRecord
}

func convertToNewUserRedemptionRecord(oldRedemptionRecord oldrecordstypes.UserRedemptionRecord) recordstypes.UserRedemptionRecord {
	return recordstypes.UserRedemptionRecord{
		Id:             oldRedemptionRecord.Id,
		Sender:         oldRedemptionRecord.Sender,
		Receiver:       oldRedemptionRecord.Receiver,
		Amount:         types.Int(oldRedemptionRecord.Amount),
		Denom:          oldRedemptionRecord.Denom,
		HostZoneId:     oldRedemptionRecord.HostZoneId,
		EpochNumber:    oldRedemptionRecord.EpochNumber,
		ClaimIsPending: oldRedemptionRecord.ClaimIsPending,
	}
}
