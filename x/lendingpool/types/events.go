package types

const (
	EventTypeMsgCreatePool = "create_pool"
	EventTypeMsgDeposit    = "deposit"
	EventTypeMsgWithdraw   = "withdraw"
	EventTypeMsgLiquidate  = "liquidate"
	AttributeTypeDenom     = "denom"
	AttributeTypePoolId    = "pool_id"
	AttributeTypeAmountIn  = "amount_in"
	AttributeTypeAmountOut = "amount_out"
	AttributeTypeLoanId    = "loan_id"
	AttributeValueCategory = ModuleName
)
