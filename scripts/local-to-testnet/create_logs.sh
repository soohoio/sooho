#!/bin/bash
# clean up logs one by one before creation (allows auto-updating logs with the command `while true; do make init build=logs ; sleep 5 ; done`)

set -eu
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)

source ${SCRIPT_DIR}/config.sh

LOGS_DIR=$SCRIPT_DIR/logs
TEMP_LOGS_DIR=$LOGS_DIR/temp

STATE_LOG=state.log
BALANCES_LOG=balances.log

mkdir -p $TEMP_LOGS_DIR

while true; do
    N_VALIDATORS_STAYKING=$($STAYKING_MAIN_CMD q tendermint-validator-set | grep -o address | wc -l | tr -dc '0-9')
    echo "STAYKING @ $($STAYKING_MAIN_CMD q tendermint-validator-set | head -n 1 | tr -dc '0-9') | $N_VALIDATORS_STAYKING VALS" >$TEMP_LOGS_DIR/$STATE_LOG
    echo "STAYKING @ $($STAYKING_MAIN_CMD q tendermint-validator-set | head -n 1 | tr -dc '0-9') | $N_VALIDATORS_STAYKING VALS" >$TEMP_LOGS_DIR/$BALANCES_LOG

    printf '\n%s\n' "LIST-HOST-ZONES STAYKING" >>$TEMP_LOGS_DIR/$STATE_LOG
    $STAYKING_MAIN_CMD q levstakeibc list-host-zone >>$TEMP_LOGS_DIR/$STATE_LOG
    printf '\n%s\n' "LIST-DEPOSIT-RECORDS" >>$TEMP_LOGS_DIR/$STATE_LOG
    $STAYKING_MAIN_CMD q records list-deposit-record  >> $TEMP_LOGS_DIR/$STATE_LOG
    printf '\n%s\n' "LIST-EPOCH-UNBONDING-RECORDS" >>$TEMP_LOGS_DIR/$STATE_LOG
    $STAYKING_MAIN_CMD q records list-epoch-unbonding-record  >> $TEMP_LOGS_DIR/$STATE_LOG
    printf '\n%s\n' "LIST-USER-REDEMPTION-RECORDS" >>$TEMP_LOGS_DIR/$STATE_LOG
    $STAYKING_MAIN_CMD q records list-user-redemption-record >> $TEMP_LOGS_DIR/$STATE_LOG

    printf '\n%s\n' "BALANCES STAYKING" >>$TEMP_LOGS_DIR/$BALANCES_LOG
    $STAYKING_MAIN_CMD q bank balances $(STAYKING_ADMIN_ADDRESS) >>$TEMP_LOGS_DIR/$BALANCES_LOG

    mv $TEMP_LOGS_DIR/*.log $LOGS_DIR
    sleep 3
done