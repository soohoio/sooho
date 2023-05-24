#!/bin/bash

set -eu

# COMMNON
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
echo "$SCRIPT_DIR"
source ${SCRIPT_DIR}/keys.sh

STATE=$SCRIPT_DIR/state
LOGS=$SCRIPT_DIR/logs

KEYS_LOGS=$LOGS/keys.log

STAYKING_HOME=$STATE/stayking1

STAYKING_LOGS=$LOGS/stayking.log

DOCKER_COMPOSE="docker-compose -f $SCRIPT_DIR/docker-compose.yml"

# DENOMS
ATOM_DENOM='uatom'
STAYKING_DENOM='ustay'
STATOM_DENOM="stuatom"
OSMO_DENOM="uosmo"
STOSMO_DENOM="stuosmo"
EVMOS_DENOM="atevmos"
STEVMOS_DENOM="statevmos"

#IBC_STAYKING_DENOM='ibc/0CC4CC37A53BBD3C699114BB24E5993C33FBBA80D16BE7D04E7ECB4CB6DAD11D'

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
IBC_EVMOS_CHANNEL_4_DENOM='ibc/725907476F79A96A2650A4D124501B5D236AB9DDFAF216F929833C6B51E42902'
IBC_EVMOS_CHANNEL_5_DENOM='ibc/B249D1E86F588286FEA286AA8364FFCE69EC65604BD7869D824ADE40F00FA25B'
IBC_EVMOS_CHANNEL_6_DENOM='ibc/35E771B8682D828173F4B795F6C307780F96DC64D6F914FAE4CC9B4666F66364'

# COIN TYPES
# Coin types can be found at https://github.com/satoshilabs/slips/blob/master/slip-0044.md
COSMOS_COIN_TYPE=118
ETH_COIN_TYPE=60
TERRA_COIN_TYPE=330


# STAYKING
STAYKING_CHAIN_ID=stayking-hub
STAYKING_NODE_PREFIX=stayking
STAYKING_NUM_NODES=3
STAYKING_VAL_PREFIX=val
STAYKING_DENOM=$STAYKING_DENOM
STAYKING_RPC_PORT=26657
STAYKING_ADMIN_ACCT=admin
STAYKING_MAIN_CMD="$SCRIPT_DIR/../../build/staykingd --home $STATE/${STAYKING_NODE_PREFIX}1"


#HOST
HOST_CHAINS=(EVMOS)

EVMOS_CHAIN_ID=evmos_9001-2
EVMOS_ENDPOINT=https://evmos-mainnet-archive-tendermint.allthatnode.com:443
EVMOS_ACCOUNT_PREFIX=evmos
EVMOS_VAL_PREFIX=eval
EVMOS_DENOM=aevmos
EVMOS_BINARY=build/evmosd
EVMOS_COIN_TYPE=$ETH_COIN_TYPE

#RELAYER
RELAYER_GAIA_EXEC="$DOCKER_COMPOSE run --rm relayer-gaiaMainnet"
RELAYER_EVMOS_EXEC="$DOCKER_COMPOSE run --rm relayer-evmosMainnet"
RELAYER_CMD="$SCRIPT_DIR/../../build/relayer --home $STATE/relayer"
RELAYER_STAYKING_ACCT=rly1
RELAYER_GAIA_ACCT=rly2
RELAYER_EVMOS_ACCT=rly3

#HERMES
HERMES_GAIA_EXEC="$DOCKER_COMPOSE run --rm hermes-gaiaMainnet"
#hermes_config=$STATE/hermes-${chain_name}
#hermes_logs=${LOGS}/hermes-${chain_name}.log

# STAYKING CHAIN PARAMS
BLOCK_TIME='5s'
STAYKING_DAY_EPOCH_DURATION="86400s"
STAYKING_EPOCH_EPOCH_DURATION="21600s"
STAYKING_UNBONDING_TIME="1814400s"
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
