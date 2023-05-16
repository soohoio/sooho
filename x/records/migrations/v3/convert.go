package v3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	oldrecordstypes "github.com/soohoio/stayking/v2/x/records/migrations/v3/types"
	recordstypes "github.com/soohoio/stayking/v2/x/records/types"
)

func convertToNewDepositRecord(oldDepositRecord oldrecordstypes.DepositRecord) recordstypes.DepositRecord {
	return recordstypes.DepositRecord{
		Id:                 oldDepositRecord.Id,
		Amount:             sdk.Int(oldDepositRecord.Amount),
		Denom:              oldDepositRecord.Denom,
		HostZoneId:         oldDepositRecord.HostZoneId,
		Status:             recordstypes.DepositRecord_Status(oldDepositRecord.Status),
		DepositEpochNumber: oldDepositRecord.DepositEpochNumber,
		Source:             recordstypes.DepositRecord_Source(oldDepositRecord.Source),
	}
}

func convertToNewUserRedemptionRecord(oldRedemptionRecord oldrecordstypes.UserRedemptionRecord) recordstypes.UserRedemptionRecord {
	return recordstypes.UserRedemptionRecord{
		Id:             oldRedemptionRecord.Id,
		Sender:         oldRedemptionRecord.Sender,
		Receiver:       oldRedemptionRecord.Receiver,
		Amount:         sdk.Int(oldRedemptionRecord.Amount),
		Denom:          oldRedemptionRecord.Denom,
		HostZoneId:     oldRedemptionRecord.HostZoneId,
		EpochNumber:    oldRedemptionRecord.EpochNumber,
		ClaimIsPending: oldRedemptionRecord.ClaimIsPending,
		PositionId:     uint64(0),
	}
}
