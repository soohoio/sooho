package types

const (
	EventTypeRegisterZone = "register_zone"

	AttributeKeyAck = "acknowledgement" // IBC Packet Handshake step "Acknowledgement" msg

	AttributeValueCategory     = ModuleName
	AttributeKeyConnectionId   = "connection_id"
	AttributeKeyRecipientChain = "chain_id"
)
