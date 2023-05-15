package types

const (
	EventTypeMsgCreatePool          = "create_pool"
	EventTypeMsgDeposit             = "deposit"
	EventTypeMsgWithdraw            = "withdraw"
	EventTypeMsgLiquidate           = "liquidate"
	AttributeTypeDenom              = "denom"
	AttributeTypePoolId             = "pool_id"
	AttributeTypeAmountIn           = "amount_in"
	AttributeTypeAmountOut          = "amount_out"
	AttributeTypeRepayWithChange    = "repay_with_change"
	AttributeTypeRepayWithoutChange = "repay_without_change"
	AttributeTypeLoanId             = "loan_id"
	AttributeTypeBorrowerAddress    = "borrwoer_address"
	AttributeTypeTotalValue         = "total_value"
	AttributeTypeBorrowedValue      = "borrowed_value"
	AttributeTypeRepayValue         = "replay_value"
	AttributeTypeChangeValue        = "change_value"
	AttributeValueCategory          = ModuleName
)
