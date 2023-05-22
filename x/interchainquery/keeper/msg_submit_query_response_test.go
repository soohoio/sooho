package keeper_test

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/proto/tendermint/crypto"

	"github.com/soohoio/stayking/v3/x/interchainquery/types"
)

const (
	HostChainId = "GAIA"
)

type MsgSubmitQueryResponseTestCase struct {
	validMsg types.MsgSubmitQueryResponse
	goCtx    context.Context
	query    types.Query
}

func (s *KeeperTestSuite) SetupMsgSubmitQueryResponse() MsgSubmitQueryResponseTestCase {
	// set up IBC
	s.CreateTransferChannel(HostChainId)

	// define the query
	goCtx := sdk.WrapSDKContext(s.Ctx)
	h, err := s.App.StakeibcKeeper.GetLightClientHeightSafely(s.Ctx, s.TransferPath.EndpointA.ConnectionID)
	s.Require().NoError(err)
	height := int64(h - 1) // start at the (LC height) - 1  height, which is the height the query executes at!
	result := []byte("result-example")
	proofOps := crypto.ProofOps{}
	fromAddress := s.TestAccs[0].String()
	expectedId := "9792c1d779a3846a8de7ae82f31a74d308b279a521fa9e0d5c4f08917117bf3e"

	_, addr, _ := bech32.DecodeAndConvert(s.TestAccs[0].String())
	data := banktypes.CreateAccountBalancesPrefix(addr)
	// save the query to StayKing state, so it can be retrieved in the response
	query := types.Query{
		Id:           expectedId,
		CallbackId:   "withdrawalbalance",
		ChainId:      HostChainId,
		ConnectionId: s.TransferPath.EndpointA.ConnectionID,
		QueryType:    types.BANK_STORE_QUERY_WITH_PROOF,
		Request:      append(data, []byte(HostChainId)...),
		Ttl:          uint64(12545592938) * uint64(1000000000), // set ttl to August 2050, mult by nano conversion factor
	}

	return MsgSubmitQueryResponseTestCase{
		validMsg: types.MsgSubmitQueryResponse{
			ChainId:     HostChainId,
			QueryId:     expectedId,
			Result:      result,
			ProofOps:    &proofOps,
			Height:      height,
			FromAddress: fromAddress,
		},
		goCtx: goCtx,
		query: query,
	}
}

func (s *KeeperTestSuite) TestMsgSubmitQueryResponse_WrongProof() {
	tc := s.SetupMsgSubmitQueryResponse()

	s.App.InterchainqueryKeeper.SetQuery(s.Ctx, tc.query)

	resp, err := s.GetMsgServer().SubmitQueryResponse(tc.goCtx, &tc.validMsg)
	s.Require().ErrorContains(err, "Unable to verify membership proof: proof cannot be empty")
	s.Require().Nil(resp)
}

func (s *KeeperTestSuite) TestMsgSubmitQueryResponse_UnknownId() {
	tc := s.SetupMsgSubmitQueryResponse()

	tc.query.Id = tc.query.Id + "INVALID_SUFFIX" // create an invalid query id
	s.App.InterchainqueryKeeper.SetQuery(s.Ctx, tc.query)

	resp, err := s.GetMsgServer().SubmitQueryResponse(tc.goCtx, &tc.validMsg)
	s.Require().NoError(err)
	s.Require().NotNil(resp)
	s.Require().Equal(&types.MsgSubmitQueryResponseResponse{}, resp)

	// check that the query is STILL in the store, as it should NOT be deleted because the query was not found
	_, found := s.App.InterchainqueryKeeper.GetQuery(s.Ctx, tc.query.Id)
	s.Require().True(found)
}

func (s *KeeperTestSuite) TestMsgSubmitQueryResponse_ExceededTtl() {
	tc := s.SetupMsgSubmitQueryResponse()

	// Remove key from the query type so to bypass the VerifyKeyProof function
	tc.query.QueryType = strings.ReplaceAll(tc.query.QueryType, "key", "")

	// set ttl to be expired
	tc.query.Ttl = uint64(1)
	s.App.InterchainqueryKeeper.SetQuery(s.Ctx, tc.query)

	resp, err := s.GetMsgServer().SubmitQueryResponse(tc.goCtx, &tc.validMsg)
	s.Require().NoError(err)
	s.Require().NotNil(resp)

	// check that the query was deleted (since the query timed out)
	_, found := s.App.InterchainqueryKeeper.GetQuery(s.Ctx, tc.query.Id)
	s.Require().False(found)
}

func (s *KeeperTestSuite) TestMsgSubmitQueryResponse_FindAndInvokeCallback_WrongHostZone() {
	tc := s.SetupMsgSubmitQueryResponse()

	s.App.InterchainqueryKeeper.SetQuery(s.Ctx, tc.query)

	// rather than testing by executing the callback in its entirety,
	//   check by invoking it without a registered host zone and catching the appropriate error
	err := s.App.InterchainqueryKeeper.InvokeCallback(s.Ctx, &tc.validMsg, tc.query)
	s.Require().ErrorContains(err, "no registered zone for queried chain ID", "callback was invoked")
}
