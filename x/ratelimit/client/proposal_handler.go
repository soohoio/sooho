package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/soohoio/stayking/x/ratelimit/client/cli"
)

var (
	AddRateLimitProposalHandler    = govclient.NewProposalHandler(cli.CmdAddRateLimitProposal)
	UpdateRateLimitProposalHandler = govclient.NewProposalHandler(cli.CmdUpdateRateLimitProposal)
	RemoveRateLimitProposalHandler = govclient.NewProposalHandler(cli.CmdRemoveRateLimitProposal)
	ResetRateLimitProposalHandler  = govclient.NewProposalHandler(cli.CmdResetRateLimitProposal)
)
