package types

const (
	EventTypeRegisterZone                        = "register_zone"
	EventTypeUpdateZone                          = "update_zone"
	EventTypeDeleteZone                          = "delete_zone"
	EventTypeAddValidator                        = "add_validator"
	EventTypeAdjustPosition                      = "adjust_position"
	EventTypeChangeValidatorWeight               = "change_validator_position"
	EventTypeClearBalance                        = "clear_balance"
	EventTypeDeleteValidator                     = "delete_validator"
	EventTypeExitLeverageStake                   = "exit_leverage_stake"
	EventTypeStakeWithLeverage                   = "stake_with_leverage"
	EventTypeStakeWithoutLeverage                = "stake_without_leverage"
	EventTypeRebalanceValidator                  = "rebalance_validator"
	EventTypeRedeemStake                         = "redeem_stake"
	EventTypeRestoreInterchainAccount            = "restore_interchain_account"
	EventTypeReleaseUnbondedAssetWithLeverage    = "release_unbonded_asset_with_leverage"
	EventTypeReleaseUnbondedAssetWithoutLeverage = "release_unbonded_asset_without_leverage"

	AttributeKeyAck = "acknowledgement" // IBC Packet Handshake step "Acknowledgement" msg

	AttributeValueCategory             = ModuleName
	AttributeKeyConnectionId           = "connection_id"
	AttributeKeyChannelId              = "channel_id"
	AttributeKeyRecipientChain         = "chain_id"
	AttributeKeyAccountName            = "acc_name"
	AttributeKeyAddress                = "address"
	AttributeKeyFromAddress            = "from_address"
	AttributeKeyToAddress              = "to_address"
	AttributeKeyColleteralTokenAmount  = "colleteral_token_amount"
	AttributeKeyDebtTokenAmount        = "debt_token_amount"
	AttributeKeyTransferTokenAmount    = "transfer_token_amount"
	AttributeKeyValidatorWeight        = "validator_weight"
	AttributeKeyPositionId             = "position_id"
	AttributeKeyLoanId                 = "loan_id"
	AttributeKeyUserRedemptionRecordId = "user_redemption_record_id"
	AttributeKeyHostDenom              = "host_denom"
	AttributeKeyIBCDenom               = "ibc_denom"
	AttributeKeyEpochNumber            = "epoch_number"
	AttributeKeyNativeTokenAmount      = "native_token_amount"
	AttributeKeyStTokenAmount          = "st_token_amount"
	AttributeKeyDepositRecordId        = "deposit_record_id"
	AttributeKey
)
