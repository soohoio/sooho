#!/bin/bash
# clean up logs one by one before creation (allows auto-updating logs with the command `while true; do make init build=logs ; sleep 5 ; done`)

set -eu
SCRIPT_DIR=$( cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd )

source ${SCRIPT_DIR}/../config.sh

LOGS_DIR=$SCRIPT_DIR/logs
TEMP_LOGS_DIR=$LOGS_DIR/temp

STATE_LOG=state.log
BALANCES_LOG=balances.log

mkdir -p $TEMP_LOGS_DIR

while true; do
    N_VALIDATORS_STAYKING=$($STAYKING_MAIN_CMD q tendermint-validator-set | grep -o address | wc -l | tr -dc '0-9')
    echo "STAYKING @ $($STAYKING_MAIN_CMD q tendermint-validator-set | head -n 1 | tr -dc '0-9') | $N_VALIDATORS_STAYKING VALS" >$TEMP_LOGS_DIR/$STATE_LOG
    echo "STAYKING @ $($STAYKING_MAIN_CMD q tendermint-validator-set | head -n 1 | tr -dc '0-9') | $N_VALIDATORS_STAYKING VALS" >$TEMP_LOGS_DIR/$BALANCES_LOG

    for chain_id in ${HOST_CHAINS[@]}; do
        HOST_MAIN_CMD=$(GET_VAR_VALUE ${chain_id}_MAIN_CMD)
        N_VALIDATORS_HOST=$($HOST_MAIN_CMD q tendermint-validator-set | grep -o address | wc -l | tr -dc '0-9')
        echo "$chain_id   @ $($HOST_MAIN_CMD q tendermint-validator-set | head -n 1 | tr -dc '0-9') | $N_VALIDATORS_HOST VALS" >>$TEMP_LOGS_DIR/$STATE_LOG
        echo "$chain_id   @ $($HOST_MAIN_CMD q tendermint-validator-set | head -n 1 | tr -dc '0-9') | $N_VALIDATORS_HOST VALS" >>$TEMP_LOGS_DIR/$BALANCES_LOG
    done
    sleep 3

    printf '\n%s\n' "LIST-HOST-ZONES STAYKING" >>$TEMP_LOGS_DIR/$STATE_LOG
    $STAYKING_MAIN_CMD q stakeibc list-host-zone >>$TEMP_LOGS_DIR/$STATE_LOG
    printf '\n%s\n' "LIST-DEPOSIT-RECORDS" >>$TEMP_LOGS_DIR/$STATE_LOG
    $STAYKING_MAIN_CMD q records list-deposit-record  >> $TEMP_LOGS_DIR/$STATE_LOG
    printf '\n%s\n' "LIST-EPOCH-UNBONDING-RECORDS" >>$TEMP_LOGS_DIR/$STATE_LOG
    $STAYKING_MAIN_CMD q records list-epoch-unbonding-record  >> $TEMP_LOGS_DIR/$STATE_LOG
    printf '\n%s\n' "LIST-USER-REDEMPTION-RECORDS" >>$TEMP_LOGS_DIR/$STATE_LOG
    $STAYKING_MAIN_CMD q records list-user-redemption-record >> $TEMP_LOGS_DIR/$STATE_LOG
    sleep 3
    printf '\n%s\n' "BALANCES STAYKING" >>$TEMP_LOGS_DIR/$BALANCES_LOG
    $STAYKING_MAIN_CMD q bank balances $(STAYKING_ADDRESS) >>$TEMP_LOGS_DIR/$BALANCES_LOG
    for chain_id in ${HOST_CHAINS[@]}; do

        HOST_MAIN_CMD=$(GET_VAR_VALUE GAIA_MAIN_CMD)
        DELEGATION_ICA_ADDR=$(GET_ICA_ADDR $GAIA_CHAIN_ID delegation)
        REDEMPTION_ICA_ADDR=$(GET_ICA_ADDR $GAIA_CHAIN_ID redemption)
        WITHDRAWAL_ICA_ADDR=$(GET_ICA_ADDR $GAIA_CHAIN_ID withdrawal)
        FEE_ICA_ADDR=$(GET_ICA_ADDR $GAIA_CHAIN_ID fee)

        printf '\n%s\n' "==========================  gaia-localnet  =============================" >>$TEMP_LOGS_DIR/$BALANCES_LOG

        printf '\n%s\n' "BALANCES gaia-localnet (DELEGATION ACCT)" >>$TEMP_LOGS_DIR/$BALANCES_LOG
        $HOST_MAIN_CMD q bank balances $DELEGATION_ICA_ADDR >>$TEMP_LOGS_DIR/$BALANCES_LOG
        printf '\n%s\n' "DELEGATIONS gaia-localnet (DELEGATION ACCT)" >>$TEMP_LOGS_DIR/$BALANCES_LOG
        $HOST_MAIN_CMD q staking delegations $DELEGATION_ICA_ADDR >>$TEMP_LOGS_DIR/$BALANCES_LOG
        printf '\n%s\n' "UNBONDING-DELEGATIONS gaia-localnet (DELEGATION ACCT)" >>$TEMP_LOGS_DIR/$BALANCES_LOG
        $HOST_MAIN_CMD q staking unbonding-delegations $DELEGATION_ICA_ADDR >>$TEMP_LOGS_DIR/$BALANCES_LOG

        printf '\n%s\n' "BALANCES gaia-localnet (REDEMPTION ACCT)" >>$TEMP_LOGS_DIR/$BALANCES_LOG
        $HOST_MAIN_CMD q bank balances $REDEMPTION_ICA_ADDR >>$TEMP_LOGS_DIR/$BALANCES_LOG
        printf '\n%s\n' "BALANCES gaia-localnet (FEE ACCT)" >>$TEMP_LOGS_DIR/$BALANCES_LOG
        $HOST_MAIN_CMD q bank balances $FEE_ICA_ADDR >>$TEMP_LOGS_DIR/$BALANCES_LOG
        printf '\n%s\n' "BALANCES gaia-localnet (WITHDRAWAL ACCT)" >>$TEMP_LOGS_DIR/$BALANCES_LOG
        $HOST_MAIN_CMD q bank balances $WITHDRAWAL_ICA_ADDR >>$TEMP_LOGS_DIR/$BALANCES_LOG
    done

    mv $TEMP_LOGS_DIR/*.log $LOGS_DIR
    sleep 3
done