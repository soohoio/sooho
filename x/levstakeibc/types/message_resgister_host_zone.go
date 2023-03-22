package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const TypeMsgRegisterHostZone = "register_host_zone"

var _ sdk.Msg = &MsgRegisterHostZone{}

func NewMsgRegisterHostZone(connectionId string, bech32prefix string, hostDenom string, ibcDenom string, creator string, transferChannelId string, unbondingFrequency uint64) *MsgRegisterHostZone {
	return &MsgRegisterHostZone{
		ConnectionId:       connectionId,
		Bech32Prefix:       bech32prefix,
		HostDenom:          hostDenom,
		IbcDenom:           ibcDenom,
		Creator:            creator,
		TransferChannelId:  transferChannelId,
		UnbondingFrequency: unbondingFrequency,
	}
}

func (msg *MsgRegisterHostZone) ValidateBasic() error {
	return nil
}

func (msg *MsgRegisterHostZone) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}
