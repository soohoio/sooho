package client

import (
	"github.com/soohoio/stayking/x/stakeibc/client/cli"
	"github.com/soohoio/stayking/x/stakeibc/client/rest"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	AddValidatorProposalHandler = govclient.NewProposalHandler(cli.CmdAddValidatorProposal, rest.ProposalAddValidatorRESTHandler)
)
