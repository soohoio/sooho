package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type InterestModelI interface {
	proto.Message

	GetAPR(utilizationRate sdk.Dec) sdk.Dec
	ModelType() string
	ValidateBasic() error
	String() string
}
