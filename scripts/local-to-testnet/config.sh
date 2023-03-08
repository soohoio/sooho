#!/bin/bash

set -eu

source ${SCRIPT_DIR}/keys.sh

# COMMNON
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

STATE=$SCRIPT_DIR/state
LOGS=$SCRIPT_DIR/logs

chain_name=gaiaTestnet

KEYS_LOGS=$LOGS/keys.log

STAYKING_HOME=$STATE/stayking1

relayer_config=$STATE/relayer-${chain_name}/config
STAYKING_LOGS=$LOGS/stayking.log
relayer_logs=${LOGS}/relayer-${chain_name}.log

DOCKER_COMPOSE="docker-compose -f $SCRIPT_DIR/docker-compose.yml"

# DENOMS
ATOM_DENOM='uatom'
STAYKING_DENOM='ustay'
STATOM_DENOM="stuatom"

#IBC_STAYKING_DENOM='ibc/0CC4CC37A53BBD3C699114BB24E5993C33FBBA80D16BE7D04E7ECB4CB6DAD11D'
IBC_HOST_CHANNEL_0_DENOM='ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2'


# STAYKING
STAYKING_CHAIN_ID=stayking-testnet-001
STAYKING_NODE_PREFIX=stayking
STAYKING_NUM_NODES=3
STAYKING_VAL_PREFIX=val
STAYKING_DENOM=$STAYKING_DENOM
STAYKING_RPC_PORT=26657
STAYKING_ADMIN_ACCT=admin
STAYKING_MAIN_CMD="$SCRIPT_DIR/../../build/staykingd --home $STATE/${STAYKING_NODE_PREFIX}1"


#HOST
HOST_CHAINS=(GAIA)

HOST_CHAIN_ID=theta-testnet-001
HOST_NUM_NODES=1
HOST_ENDPOINT=https://cosmos-testnet-archive.allthatnode.com:26657/AowVlngs1uvTAB6cbCEF2y3Xwy0Qk7qL
HOST_ACCOUNT_PREFIX=cosmos
HOST_VAL_PREFIX=gval
HOST_DENOM=uatom
HOST_BINARY=build/gaiad

#RELAYER
RELAYER_GAIA_TESTNET_EXEC="$DOCKER_COMPOSE run --rm relayer-gaiaTestnet"
RELAYER_EXEC=$RELAYER_GAIA_TESTNET_EXEC
RELAYER_CMD="$SCRIPT_DIR/../../build/relayer --home $STATE/relayer"

#HERMES
HERMES_GAIA_EXEC="$DOCKER_COMPOSE run --rm hermes-gaiaTestnet"
hermes_config=$STATE/hermes-${chain_name}
hermes_logs=${LOGS}/hermes-${chain_name}.log

# STAYKING CHAIN PARAMS
BLOCK_TIME='5s'
STAYKING_DAY_EPOCH_DURATION="120s"
STAYKING_EPOCH_EPOCH_DURATION="120s"
STAYKING_UNBONDING_TIME="172800s"
MAX_DEPOSIT_PERIOD="120s"
VOTING_PERIOD="120s"
INITIAL_ANNUAL_PROVISIONS="10000000000000.000000000000000000"
VAL_TOKENS=5000000000000
STAKE_TOKENS=5000000000
ADMIN_TOKENS=1000000000
PEER_PORT=26656
RPC_PORT=26657
VAL_PREFIX=val
DENOM=ustay

NODE_NAME="stayking1"
NODE_PREFIX="stayking"


MAIN_ID=1 # Node responsible for genesis and persistent_peers
MAIN_NODE_NAME=""
MAIN_NODE_CMD=""
MAIN_NODE_ID=""
MAIN_CONFIG=""
MAIN_GENESIS=""
NUM_NODES=3

STAYKING_ADDRESS() {
  $STAYKING_MAIN_CMD keys show ${STAYKING_VAL_PREFIX}1 --keyring-backend test -a
}

GAIA_ADDRESS() {
  $GAIA_MAIN_CMD keys show ${HOST_VAL_PREFIX}1 --keyring-backend test -a
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