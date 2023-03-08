#!/bin/bash

set -eu
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

source ${SCRIPT_DIR}/keys.sh

STATE=$SCRIPT_DIR/state
LOGS=$SCRIPT_DIR/logs
UPGRADES=$SCRIPT_DIR/upgrades
SRC=$SCRIPT_DIR/src
PEER_PORT=26656
DOCKER_COMPOSE="docker-compose -f $SCRIPT_DIR/docker-compose.yml"

# Logs
STAYKING_LOGS=$LOGS/stayking.log
TX_LOGS=$SCRIPT_DIR/logs/tx.log
KEYS_LOGS=$SCRIPT_DIR/logs/keys.log

# List of hosts enabled 
#  `start-docker` defaults to just GAIA if HOST_CHAINS is empty
# `start-docker-all` always runs all hosts
HOST_CHAINS=()
if [[ "${ALL_HOST_CHAINS:-false}" == "true" ]]; then
  HOST_CHAINS=(GAIA OSMO)
elif [[ "${#HOST_CHAINS[@]}" == "0" ]]; then
  HOST_CHAINS=(GAIA)
fi

# Sets up upgrade if {UPGRADE_NAME} is non-empty
UPGRADE_NAME=""
UPGRADE_OLD_COMMIT_HASH=""
#UPGRADE_NAME="v2"
#UPGRADE_OLD_COMMIT_HASH="b8ea899f09f1da98611517a72023dfda2ade1173"

# DENOMS
ATOM_DENOM='uatom'
STAYKING_DENOM='ustay'
STATOM_DENOM="stuatom"
OSMO_DENOM="uosmo"
STOSMO_DENOM="stuosmo"

#IBC_STAYKING_DENOM='ibc/0CC4CC37A53BBD3C699114BB24E5993C33FBBA80D16BE7D04E7ECB4CB6DAD11D'
IBC_GAIA_CHANNEL_0_DENOM='ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2'
IBC_OSMO_CHANNEL_0_DENOM='ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518'
IBC_OSMO_CHANNEL_1_DENOM='ibc/0471F1C4E7AFD3F07702BEF6DC365268D64570F7C1FDC98EA6098DD6DE59817B'
# INTEGRATION TEST IBC DENOM
IBC_ATOM_DENOM=$IBC_GAIA_CHANNEL_0_DENOM
IBC_OSMO_DENOM=$IBC_OSMO_CHANNEL_1_DENOM
# CHAIN PARAMS
BLOCK_TIME='5s'
STAYKING_DAY_EPOCH_DURATION="100s"
STAYKING_EPOCH_EPOCH_DURATION="40s"
HOST_DAY_EPOCH_DURATION="60s"
HOST_HOUR_EPOCH_DURATION="60s"
HOST_WEEK_EPOCH_DURATION="60s"
STAYKING_UNBONDING_TIME="120s"
#STAYKING_UNBONDING_TIME="1814400s"
#HOST_UNBONDING_TIME="1814400s"
HOST_UNBONDING_TIME="240s"
MAX_DEPOSIT_PERIOD="120s"
VOTING_PERIOD="120s"

INITIAL_ANNUAL_PROVISIONS="10000000000000.000000000000000000"
VAL_TOKENS=5000000000000
STAKE_TOKENS=5000000000
ADMIN_TOKENS=1000000000

# STAYKING
STAYKING_CHAIN_ID=stayking-localnet
STAYKING_NODE_PREFIX=stayking
STAYKING_NUM_NODES=1
STAYKING_VAL_PREFIX=val
STAYKING_DENOM=$STAYKING_DENOM
STAYKING_RPC_PORT=26657
STAYKING_ADMIN_ACCT=admin
STAYKING_ADMIN_ADDRESS=sooho143umg272xger2eyurqfpjgt8u533s62mpz5weq
STAYKING_FEE_ADDRESS=sooho1ckh2w55t4jkdz6ck74a32mqxeyrlt8ee2ws444 #wire act strong despair apple elite glide industry journey final finish coconut repair judge payment error soul bounce public sport flee library employ position

# Binaries are contigent on whether we're doing an upgrade or not
if [[ "$UPGRADE_NAME" == "" ]]; then 
  STAYKING_CMD="$SCRIPT_DIR/../build/staykingd"
else
  STAYKING_CMD="$UPGRADES/binaries/staykingd1"
fi
STAYKING_MAIN_CMD="$STAYKING_CMD --home $SCRIPT_DIR/state/${STAYKING_NODE_PREFIX}1"

# GAIA 
GAIA_CHAIN_ID=gaia-localnet
GAIA_NODE_PREFIX=gaia
GAIA_NUM_NODES=1
GAIA_CMD="$SCRIPT_DIR/../build/gaiad"
GAIA_VAL_PREFIX=gval
GAIA_REV_ACCT=grev1
GAIA_ADDRESS_PREFIX=cosmos
GAIA_DENOM=$ATOM_DENOM
GAIA_RPC_PORT=26557
GAIA_MAIN_CMD="$GAIA_CMD --home $SCRIPT_DIR/state/${GAIA_NODE_PREFIX}1"
GAIA_RECEIVER_ADDRESS='cosmos1g6qdx6kdhpf000afvvpte7hp0vnpzapuyxp8uf'

# OSMO
OSMO_CHAIN_ID=osmo-localnet
OSMO_NODE_PREFIX=osmo
OSMO_NUM_NODES=1
OSMO_CMD="$SCRIPT_DIR/../build/osmosisd"
OSMO_VAL_PREFIX=oval
OSMO_REV_ACCT=orev1
OSMO_ADDRESS_PREFIX=osmo
OSMO_DENOM=$OSMO_DENOM
OSMO_RPC_PORT=26357
OSMO_MAIN_CMD="$OSMO_CMD --home $SCRIPT_DIR/state/${OSMO_NODE_PREFIX}1"
OSMO_RECEIVER_ADDRESS='osmo1w6wdc2684g9h3xl8nhgwr282tcxx4kl06n4sjl'

# RELAYER
RELAYER_CMD="$SCRIPT_DIR/../build/relayer --home $STATE/relayer"
RELAYER_GAIA_EXEC="$DOCKER_COMPOSE run --rm relayer-gaia"
RELAYER_OSMO_EXEC="$DOCKER_COMPOSE run --rm relayer-osmo"

RELAYER_STAYKING_ACCT=rly1
RELAYER_GAIA_ACCT=rly2
RELAYER_OSMO_ACCT=rly3
HOST_RELAYER_ACCTS=($RELAYER_GAIA_ACCT $RELAYER_OSMO_ACCT)
RELAYER_MNEMONICS=("$RELAYER_GAIA_MNEMONIC" "$RELAYER_OSMO_MNEMONIC")
HERMES_GAIA_EXEC="$DOCKER_COMPOSE run --rm hermes-gaia"


STAYKING_ADDRESS() {
  $STAYKING_MAIN_CMD keys show ${STAYKING_VAL_PREFIX}1 --keyring-backend test -a
}
GAIA_ADDRESS() {
  $GAIA_MAIN_CMD keys show ${GAIA_VAL_PREFIX}1 --keyring-backend test -a 
}
OSMO_ADDRESS() {
  $OSMO_MAIN_CMD keys show ${OSMO_VAL_PREFIX}1 --keyring-backend test -a
}

CSLEEP() {
  for i in $(seq $1); do
    sleep 1
    printf "\r\t$(($1 - $i))s left..."
  done
}

GET_VAR_VALUE() {
  var_name="$1"
  echo "${!var_name}"
}

WAIT_FOR_BLOCK() {
  num_blocks="${2:-1}"
  for i in $(seq $num_blocks); do
    ( tail -f -n0 $1 & ) | grep -q "INF executed block height="
  done
}

WAIT_FOR_STRING() {
  ( tail -f -n0 $1 & ) | grep -q "$2"
}

GET_VAL_ADDR() {
  chain=$1
  val_index=$2

  MAIN_CMD=$(GET_VAR_VALUE ${chain}_MAIN_CMD)
  $MAIN_CMD q staking validators | grep ${chain}_${val_index} -A 5 | grep operator | awk '{print $2}'
}

GET_ICA_ADDR() {
  chain_id="$1"
  ica_type="$2" #delegation, fee, redemption, or withdrawal

  $STAYKING_MAIN_CMD q stakeibc show-host-zone $chain_id | grep ${ica_type}_account -A 1 | grep address | awk '{print $2}'
}

TRIM_TX() {
  grep -E "code:|txhash:" | sed 's/^/  /'
}