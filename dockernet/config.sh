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
  HOST_CHAINS=(GAIA OSMOSIS EVMOS)
elif [[ "${#HOST_CHAINS[@]}" == "0" ]]; then
  HOST_CHAINS=(OSMOSIS)
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
EVMOS_DENOM="aevmos"
STEVMOS_DENOM="staevmos"


#IBC_STAYKING_DENOM='ibc/0CC4CC37A53BBD3C699114BB24E5993C33FBBA80D16BE7D04E7ECB4CB6DAD11D'
IBC_GAIA_CHANNEL_0_DENOM='ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2'
IBC_GAIA_CHANNEL_1_DENOM='ibc/C4CFF46FD6DE35CA4CF4CE031E643C8FDC9BA4B99AE598E9B0ED98FE3A2319F9'
IBC_GAIA_CHANNEL_2_DENOM='ibc/9117A26BA81E29FA4F78F57DC2BD90CD3D26848101BA880445F119B22A1E254E'
IBC_GAIA_CHANNEL_3_DENOM='ibc/A4DB47A9D3CF9A068D454513891B526702455D3EF08FB9EB558C561F9DC2B701'

IBC_OSMOSIS_CHANNEL_0_DENOM='ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518'
IBC_OSMOSIS_CHANNEL_1_DENOM='ibc/0471F1C4E7AFD3F07702BEF6DC365268D64570F7C1FDC98EA6098DD6DE59817B'
IBC_OSMOSIS_CHANNEL_2_DENOM='ibc/13B2C536BB057AC79D5616B8EA1B9540EC1F2170718CAFF6F0083C966FFFED0B'
IBC_OSMOSIS_CHANNEL_3_DENOM='ibc/47BD209179859CDE4A2806763D7189B6E6FE13A17880FE2B42DE1E6C1E329E23'

IBC_EVMOS_CHANNEL_0_DENOM='ibc/8EAC8061F4499F03D2D1419A3E73D346289AE9DB89CAB1486B72539572B1915E'
IBC_EVMOS_CHANNEL_1_DENOM='ibc/6993F2B27985C9363D3B94D702111940055833A2BA86DA93F33A67D03E4D1B7D'
IBC_EVMOS_CHANNEL_2_DENOM='ibc/0E8BF52B5A990E16C4AF2E5ED426503F3F0B12067FB2B4B660015A64CCE38EA0'
IBC_EVMOS_CHANNEL_3_DENOM='ibc/5590FF5DA750B007818BB275A9CDC8B6704414F8411E2EF8CC6C43A913B6CE88'
# INTEGRATION TEST IBC DENOM
IBC_ATOM_DENOM=$IBC_GAIA_CHANNEL_0_DENOM
IBC_OSMOSIS_DENOM=$IBC_OSMOSIS_CHANNEL_1_DENOM
IBC_EVMOS_DENOM=$IBC_EVMOS_CHANNEL_2_DENOM
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

# COIN TYPES
# Coin types can be found at https://github.com/satoshilabs/slips/blob/master/slip-0044.md
COSMOS_COIN_TYPE=118
ETH_COIN_TYPE=60
TERRA_COIN_TYPE=330

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
GAIA_COIN_TYPE=$COSMOS_COIN_TYPE

# OSMOSIS 
OSMOSIS_CHAIN_ID=osmosis-localnet
OSMOSIS_NODE_PREFIX=osmosis
OSMOSIS_NUM_NODES=1
OSMOSIS_CMD="$SCRIPT_DIR/../build/osmosisd"
OSMOSIS_VAL_PREFIX=oval
OSMOSIS_REV_ACCT=orev1
OSMOSIS_ADDRESS_PREFIX=osmo
OSMOSIS_DENOM=$OSMO_DENOM
OSMOSIS_RPC_PORT=26457
OSMOSIS_MAIN_CMD="$OSMOSIS_CMD --home $SCRIPT_DIR/state/${OSMOSIS_NODE_PREFIX}1"
OSMOSIS_RECEIVER_ADDRESS='osmo1w6wdc2684g9h3xl8nhgwr282tcxx4kl06n4sjl'
OSMOSIS_COIN_TYPE=$COSMOS_COIN_TYPE

# EVMOS
EVMOS_CHAIN_ID=evmos_9001-2
EVMOS_NODE_PREFIX=evmos
EVMOS_NUM_NODES=1
EVMOS_CMD="$SCRIPT_DIR/../build/evmosd"
EVMOS_VAL_PREFIX=eval
EVMOS_REV_ACCT=erev1
EVMOS_ADDRESS_PREFIX=evmos
EVMOS_DENOM=$EVMOS_DENOM
EVMOS_RPC_PORT=26357
EVMOS_MAIN_CMD="$EVMOS_CMD --home $SCRIPT_DIR/state/${EVMOS_NODE_PREFIX}1"
EVMOS_RECEIVER_ADDRESS='evmos1st0l6uvwwqyxj3sce9de0kzcnq3dmu9dk594x9'
EVMOS_COIN_TYPE=$ETH_COIN_TYPE


# RELAYER
RELAYER_CMD="$SCRIPT_DIR/../build/relayer --home $STATE/relayer"
RELAYER_GAIA_EXEC="$DOCKER_COMPOSE run --rm relayer-gaia"
RELAYER_OSMOSIS_EXEC="$DOCKER_COMPOSE run --rm relayer-osmosis"
RELAYER_EVMOS_EXEC="$DOCKER_COMPOSE run --rm relayer-evmos"

RELAYER_STAYKING_ACCT=rly1
RELAYER_GAIA_ACCT=rly2
RELAYER_OSMOSIS_ACCT=rly3
RELAYER_EVMOS_ACCT=rly4
HOST_RELAYER_ACCTS=($RELAYER_GAIA_ACCT $RELAYER_OSMOSIS_ACCT $RELAYER_EVMOS_ACCT)
RELAYER_MNEMONICS=("$RELAYER_GAIA_MNEMONIC" "$RELAYER_OSMOSIS_MNEMONIC" "$RELAYER_EVMOS_MNEMONIC")
HERMES_GAIA_EXEC="$DOCKER_COMPOSE run --rm hermes-gaia"
HERMES_OSMOSIS_EXEC="$DOCKER_COMPOSE run --rm hermes-osmosis"
HERMES_EVMOS_EXEC="$DOCKER_COMPOSE run --rm hermes-evmos"


STAYKING_ADDRESS() {
  $STAYKING_MAIN_CMD keys show ${STAYKING_VAL_PREFIX}1 --keyring-backend test -a
}
GAIA_ADDRESS() {
  $GAIA_MAIN_CMD keys show ${GAIA_VAL_PREFIX}1 --keyring-backend test -a 
}
OSMOSIS_ADDRESS() {
  $OSMOSIS_MAIN_CMD keys show ${OSMOSIS_VAL_PREFIX}1 --keyring-backend test -a
}
EVMOS_ADDRESS() {
  $EVMOS_MAIN_CMD keys show ${EVMOS_VAL_PREFIX}1 --keyring-backend test -a
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
