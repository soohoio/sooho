package client

import (
	"github.com/soohoio/stayking/v2/x/stakeibc/client/cli"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	AddValidatorProposalHandler = govclient.NewProposalHandler(cli.CmdAddValidatorProposal)
)
