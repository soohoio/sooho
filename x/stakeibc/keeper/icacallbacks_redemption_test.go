package keeper_test

import (
	"fmt"
	icacallbacktypes "github.com/soohoio/stayking/x/icacallbacks/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	_ "github.com/stretchr/testify/suite"

	recordtypes "github.com/soohoio/stayking/x/records/types"
	stakeibckeeper "github.com/soohoio/stayking/x/stakeibc/keeper"
	"github.com/soohoio/stayking/x/stakeibc/types"
	stakeibc "github.com/soohoio/stayking/x/stakeibc/types"
)

type RedemptionCallbackState struct {
	epochUnbondingNumbers   []uint64
	userRedemptionRecordIds []string
	epochNumber             uint64
}

type RedemptionCallbackArgs struct {
	packet      channeltypes.Packet
	ackResponse *icacallbacktypes.AcknowledgementResponse
	args        []byte
}

type RedemptionCallbackTestCase struct {
	initialState RedemptionCallbackState
	validArgs    RedemptionCallbackArgs
}

func (s *KeeperTestSuite) SetupRedemptionCallback() RedemptionCallbackTestCase {
	epochNumber := uint64(1)

	// individual userRedemptionRecords should be claimable, as long as the host zone unbonding allows for claims
	recordId1 := recordtypes.UserRedemptionRecordKeyFormatter(HostChainId, epochNumber, "sender")
	userRedemptionRecord1 := recordtypes.UserRedemptionRecord{
		Id: recordId1,
	}
	recordId2 := recordtypes.UserRedemptionRecordKeyFormatter(HostChainId, epochNumber, "other_sender")
	userRedemptionRecord2 := recordtypes.UserRedemptionRecord{
		Id: recordId2,
	}

	// the hostZoneUnbonding should have HostZoneUnbonding_EXIT_TRANSFER_QUEUE - meaning unbonding has completed, but the tokens
	// have not yet been transferred to the redemption account
	hostZoneUnbonding := recordtypes.HostZoneUnbonding{
		HostZoneId:            HostChainId,
		Status:                recordtypes.HostZoneUnbonding_EXIT_TRANSFER_QUEUE,
		UserRedemptionRecords: []string{recordId1, recordId2},
	}

	epochUnbondingRecord := recordtypes.EpochUnbondingRecord{
		EpochNumber:        epochNumber,
		HostZoneUnbondings: []*recordtypes.HostZoneUnbonding{&hostZoneUnbonding},
	}
	hostZone := stakeibc.HostZone{
		ChainId:        HostChainId,
		HostDenom:      Atom,
		IbcDenom:       IbcAtom,
		RedemptionRate: sdk.NewDec(1.0),
	}
	s.App.StakeibcKeeper.SetHostZone(s.Ctx, hostZone)
	s.App.RecordsKeeper.SetEpochUnbondingRecord(s.Ctx, epochUnbondingRecord)
	s.App.RecordsKeeper.SetUserRedemptionRecord(s.Ctx, userRedemptionRecord1)
	s.App.RecordsKeeper.SetUserRedemptionRecord(s.Ctx, userRedemptionRecord2)

	packet := channeltypes.Packet{}
	ackResponse := icacallbacktypes.AcknowledgementResponse{Status: icacallbacktypes.AckResponseStatus_SUCCESS}
	callbackArgs := types.RedemptionCallback{
		HostZoneId:              HostChainId,
		EpochUnbondingRecordIds: []uint64{epochNumber},
	}
	callbackArgsBz, err := s.App.StakeibcKeeper.MarshalRedemptionCallbackArgs(s.Ctx, callbackArgs)
	s.Require().NoError(err)

	return RedemptionCallbackTestCase{
		initialState: RedemptionCallbackState{
			epochUnbondingNumbers:   []uint64{epochNumber},
			userRedemptionRecordIds: []string{userRedemptionRecord1.Id, userRedemptionRecord2.Id},
			epochNumber:             epochNumber,
		},
		validArgs: RedemptionCallbackArgs{
			packet:      packet,
			ackResponse: &ackResponse,
			args:        callbackArgsBz,
		},
	}
}

func (s *KeeperTestSuite) TestRedemptionCallback_Successful() {
	tc := s.SetupRedemptionCallback()
	initialState := tc.initialState
	validArgs := tc.validArgs

	err := stakeibckeeper.RedemptionCallback(s.App.StakeibcKeeper, s.Ctx, validArgs.packet, validArgs.ackResponse, validArgs.args)
	s.Require().NoError(err, "redemption callback succeeded")

	for _, epochNumber := range initialState.epochUnbondingNumbers {
		// fetch the epoch unbonding record
		epochUnbondingRecord, found := s.App.RecordsKeeper.GetEpochUnbondingRecord(s.Ctx, epochNumber)
		s.Require().True(found, "epoch unbonding record found")
		for _, hzu := range epochUnbondingRecord.HostZoneUnbondings {
			// check that the status is CLAIMABLE
			if hzu.HostZoneId == HostChainId {
				s.Require().Equal(recordtypes.HostZoneUnbonding_CLAIMABLE, hzu.Status, "host zone unbonding status is CLAIMABLE")
			}
		}
	}
}

func (s *KeeperTestSuite) checkRedemptionStateIfCallbackFailed(tc RedemptionCallbackTestCase) {
	initialState := tc.initialState
	for _, epochNumber := range initialState.epochUnbondingNumbers {
		// fetch the epoch unbonding record
		epochUnbondingRecord, found := s.App.RecordsKeeper.GetEpochUnbondingRecord(s.Ctx, epochNumber)
		s.Require().True(found, "epoch unbonding record found")
		for _, hzu := range epochUnbondingRecord.HostZoneUnbondings {
			// check that the status is NOT CLAIMABLE
			s.Require().Equal(recordtypes.HostZoneUnbonding_EXIT_TRANSFER_QUEUE, hzu.Status, "host zone unbonding status is NOT CLAIMABLE (EXIT_TRANSFER_QUEUE)")
		}
	}
}

func (s *KeeperTestSuite) TestRedemptionCallback_RedemptionCallbackTimeout() {
	tc := s.SetupRedemptionCallback()
	invalidArgs := tc.validArgs
	// a nil ack means the request timed out
	invalidArgs.ackResponse = nil
	err := stakeibckeeper.RedemptionCallback(s.App.StakeibcKeeper, s.Ctx, invalidArgs.packet, invalidArgs.ackResponse, invalidArgs.args)
	s.Require().NoError(err)
	s.checkRedemptionStateIfCallbackFailed(tc)
}

func (s *KeeperTestSuite) TestRedemptionCallback_RedemptionCallbackErrorOnHost() {
	tc := s.SetupRedemptionCallback()

	invalidArgs := tc.validArgs
	invalidArgs.ackResponse.Status = icacallbacktypes.AckResponseStatus_FAILURE

	err := stakeibckeeper.RedemptionCallback(s.App.StakeibcKeeper, s.Ctx, invalidArgs.packet, invalidArgs.ackResponse, invalidArgs.args)
	s.Require().NoError(err)
	s.checkRedemptionStateIfCallbackFailed(tc)
}

func (s *KeeperTestSuite) TestRedemptionCallback_WrongCallbackArgs() {
	tc := s.SetupRedemptionCallback()
	invalidArgs := tc.validArgs

	err := stakeibckeeper.RedemptionCallback(s.App.StakeibcKeeper, s.Ctx, invalidArgs.packet, invalidArgs.ackResponse, []byte("random bytes"))
	s.Require().EqualError(err, "Unable to unmarshal redemption callback args | unexpected EOF: unable to unmarshal data structure")
	s.checkRedemptionStateIfCallbackFailed(tc)
}

func (s *KeeperTestSuite) TestRedemptionCallback_EpochUnbondingRecordNotFound() {
	tc := s.SetupRedemptionCallback()
	invalidArgs := tc.validArgs
	callbackArgs := types.RedemptionCallback{
		HostZoneId:              HostChainId,
		EpochUnbondingRecordIds: []uint64{tc.initialState.epochNumber + 1},
	}
	args, err := s.App.StakeibcKeeper.MarshalRedemptionCallbackArgs(s.Ctx, callbackArgs)
	s.Require().NoError(err)
	invalidArgs.args = args
	err = stakeibckeeper.RedemptionCallback(s.App.StakeibcKeeper, s.Ctx, invalidArgs.packet, invalidArgs.ackResponse, invalidArgs.args)
	expectedErr := fmt.Sprintf("Error fetching host zone unbonding record for epoch: %d, host zone: GAIA: host zone not found", tc.initialState.epochNumber+1)
	s.Require().EqualError(err, expectedErr)
	s.checkRedemptionStateIfCallbackFailed(tc)
}

func (s *KeeperTestSuite) TestRedemptionCallback_HostZoneUnbondingNotFound() {
	tc := s.SetupRedemptionCallback()
	valid := tc.validArgs
	// remove the hzu from the epoch unbonding record
	epochUnbondingRecord, found := s.App.RecordsKeeper.GetEpochUnbondingRecord(s.Ctx, tc.initialState.epochNumber)
	s.Require().True(found)
	epochUnbondingRecord.HostZoneUnbondings = []*recordtypes.HostZoneUnbonding{}
	s.App.RecordsKeeper.SetEpochUnbondingRecord(s.Ctx, epochUnbondingRecord)
	err := stakeibckeeper.RedemptionCallback(s.App.StakeibcKeeper, s.Ctx, valid.packet, valid.ackResponse, valid.args)
	s.Require().EqualError(err, fmt.Sprintf("Error fetching host zone unbonding record for epoch: %d, host zone: GAIA: host zone not found", tc.initialState.epochNumber))
	s.checkRedemptionStateIfCallbackFailed(tc)
}
