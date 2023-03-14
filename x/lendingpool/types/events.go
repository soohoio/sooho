package types

const (
	EventTypeMsgCreatePool = "create_pool"
	EventTypeMsgDeposit    = "deposit"
	EventTypeMsgWithdraw   = "withdraw"

	AttributeTypeDenom     = "denom"
	AttributeTypePoolId    = "pool_id"
	AttributeTypeAmountIn  = "amount_in"
	AttributeTypeAmountOut = "amount_out"

	AttributeValueCategory = ModuleName
)
