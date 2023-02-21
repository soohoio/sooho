package keeper_test

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	icacallbacktypes "github.com/soohoio/stayking/x/icacallbacks/types"
	_ "github.com/stretchr/testify/suite"

	stakeibckeeper "github.com/soohoio/stayking/x/stakeibc/keeper"
	"github.com/soohoio/stayking/x/stakeibc/types"
	stakeibctypes "github.com/soohoio/stayking/x/stakeibc/types"
)

type RebalanceCallbackState struct {
	hostZone          stakeibctypes.HostZone
	initialValidators []*stakeibctypes.Validator
}

type RebalanceCallbackArgs struct {
	packet      channeltypes.Packet
	ackResponse *icacallbacktypes.AcknowledgementResponse
	args        []byte
}

type RebalanceCallbackTestCase struct {
	initialState RebalanceCallbackState
	validArgs    RebalanceCallbackArgs
}

func (s *KeeperTestSuite) SetupRebalanceCallback() RebalanceCallbackTestCase {
	rebalanceValidatorsTestCase := s.SetupRebalanceValidators()

	packet := channeltypes.Packet{}
	ackResponse := icacallbacktypes.AcknowledgementResponse{Status: icacallbacktypes.AckResponseStatus_SUCCESS}
	callbackArgs := types.RebalanceCallback{
		HostZoneId: HostChainId,
		Rebalancings: []*types.Rebalancing{
			{
				SrcValidator: "stride_VAL3",
				DstValidator: "stride_VAL1",
				Amt:          sdkmath.NewInt(104),
			},
			{
				SrcValidator: "stride_VAL4",
				DstValidator: "stride_VAL1",
				Amt:          sdkmath.NewInt(13),
			},
		},
	}
	args, err := s.App.StakeibcKeeper.MarshalRebalanceCallbackArgs(s.Ctx, callbackArgs)
	s.Require().NoError(err)

	return RebalanceCallbackTestCase{
		initialState: RebalanceCallbackState{
			hostZone:          rebalanceValidatorsTestCase.hostZone,
			initialValidators: rebalanceValidatorsTestCase.initialValidators,
		},
		validArgs: RebalanceCallbackArgs{
			packet:      packet,
			ackResponse: &ackResponse,
			args:        args,
		},
	}
}

func (s *KeeperTestSuite) TestRebalanceCallback_Successful() {
	tc := s.SetupRebalanceCallback()
	_ = tc.initialState
	validArgs := tc.validArgs

	err := stakeibckeeper.RebalanceCallback(s.App.StakeibcKeeper, s.Ctx, validArgs.packet, validArgs.ackResponse, validArgs.args)
	s.Require().NoError(err, "rebalance callback succeeded")

	hz, found := s.App.StakeibcKeeper.GetHostZone(s.Ctx, "GAIA")
	s.Require().True(found, "host zone found")

	validators := hz.GetValidators()
	s.Require().Len(validators, 5, "host zone has 5 validators")

	s.Require().Equal(sdk.NewInt(217), validators[0].DelegationAmt, "validator 1 stake")
	s.Require().Equal(sdk.NewInt(500), validators[1].DelegationAmt, "validator 2 stake")
	s.Require().Equal(sdk.NewInt(96), validators[2].DelegationAmt, "validator 3 stake")
	s.Require().Equal(sdk.NewInt(387), validators[3].DelegationAmt, "validator 4 stake")
	s.Require().Equal(sdk.NewInt(400), validators[4].DelegationAmt, "validator 5 stake")
}

func (s *KeeperTestSuite) checkDelegationStateIfCallbackFailed() {
	hz, found := s.App.StakeibcKeeper.GetHostZone(s.Ctx, "GAIA")
	s.Require().True(found, "host zone found")

	validators := hz.GetValidators()
	s.Require().Len(validators, 5, "host zone has 5 validators")

	s.Require().Equal(sdk.NewInt(100), validators[0].DelegationAmt, "validator 1 stake")
	s.Require().Equal(sdk.NewInt(500), validators[1].DelegationAmt, "validator 2 stake")
	s.Require().Equal(sdk.NewInt(200), validators[2].DelegationAmt, "validator 3 stake")
	s.Require().Equal(sdk.NewInt(400), validators[3].DelegationAmt, "validator 4 stake")
	s.Require().Equal(sdk.NewInt(400), validators[4].DelegationAmt, "validator 5 stake")
}

func (s *KeeperTestSuite) TestRebalanceCallback_Timeout() {
	tc := s.SetupRebalanceCallback()
	invalidArgs := tc.validArgs
	// a nil ack means the request timed out
	invalidArgs.ackResponse = nil

	err := stakeibckeeper.RebalanceCallback(s.App.StakeibcKeeper, s.Ctx, invalidArgs.packet, invalidArgs.ackResponse, invalidArgs.args)
	s.Require().NoError(err)
	s.checkDelegationStateIfCallbackFailed()
}

func (s *KeeperTestSuite) TestRebalanceCallback_ErrorOnHost() {
	tc := s.SetupRebalanceCallback()
	invalidArgs := tc.validArgs
	invalidArgs.ackResponse.Status = icacallbacktypes.AckResponseStatus_FAILURE

	err := stakeibckeeper.RebalanceCallback(s.App.StakeibcKeeper, s.Ctx, invalidArgs.packet, invalidArgs.ackResponse, invalidArgs.args)
	s.Require().NoError(err)
	s.checkDelegationStateIfCallbackFailed()
}

func (s *KeeperTestSuite) TestRebalanceCallback_WrongCallbackArgs() {
	tc := s.SetupRebalanceCallback()
	invalidArgs := tc.validArgs

	err := stakeibckeeper.RebalanceCallback(s.App.StakeibcKeeper, s.Ctx, invalidArgs.packet, invalidArgs.ackResponse, []byte("random bytes"))
	s.Require().EqualError(err, "Unable to unmarshal rebalance callback args | unexpected EOF: unable to unmarshal data structure")
	s.checkDelegationStateIfCallbackFailed()
}

func (s *KeeperTestSuite) TestRebalanceCallback_WrongValidator() {
	tc := s.SetupRebalanceCallback()

	callbackArgs := types.RebalanceCallback{
		HostZoneId: HostChainId,
		Rebalancings: []*types.Rebalancing{
			{
				SrcValidator: "stride_VAL3",
				DstValidator: "stride_VAL1",
				Amt:          sdk.NewInt(104),
			},
			{
				SrcValidator: "stride_VAL4_WRONG",
				DstValidator: "stride_VAL1",
				Amt:          sdk.NewInt(13),
			},
		},
	}
	invalidArgsOne, err := s.App.StakeibcKeeper.MarshalRebalanceCallbackArgs(s.Ctx, callbackArgs)
	s.Require().NoError(err)

	callbackArgs = types.RebalanceCallback{
		HostZoneId: HostChainId,
		Rebalancings: []*types.Rebalancing{
			{
				SrcValidator: "stride_VAL3",
				DstValidator: "stride_VAL1_WRONG",
				Amt:          sdk.NewInt(104),
			},
			{
				SrcValidator: "stride_VAL4",
				DstValidator: "stride_VAL1",
				Amt:          sdk.NewInt(13),
			},
		},
	}
	invalidArgsTwo, err := s.App.StakeibcKeeper.MarshalRebalanceCallbackArgs(s.Ctx, callbackArgs)
	s.Require().NoError(err)

	err = stakeibckeeper.RebalanceCallback(s.App.StakeibcKeeper, s.Ctx, tc.validArgs.packet, tc.validArgs.ackResponse, invalidArgsOne)
	s.Require().EqualError(err, "validator not found stride_VAL4_WRONG: invalid request")
	s.checkDelegationStateIfCallbackFailed()

	err = stakeibckeeper.RebalanceCallback(s.App.StakeibcKeeper, s.Ctx, tc.validArgs.packet, tc.validArgs.ackResponse, invalidArgsTwo)
	s.Require().EqualError(err, "validator not found stride_VAL1_WRONG: invalid request")
	s.checkDelegationStateIfCallbackFailed()
}
