#!/bin/bash

set -eu
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

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
  HOST_CHAINS=(GAIA)
elif [[ "${#HOST_CHAINS[@]}" == "0" ]]; then 
  HOST_CHAINS=(GAIA)
fi

# Sets up upgrade if {UPGRADE_NAME} is non-empty
UPGRADE_NAME=""
UPGRADE_OLD_COMMIT_HASH=""

# DENOMS
ATOM_DENOM='uatom'
STAYKING_DENOM='ustay'
STATOM_DENOM="stuatom"

IBC_STAYKING_DENOM='ibc/0CC4CC37A53BBD3C699114BB24E5993C33FBBA80D16BE7D04E7ECB4CB6DAD11D'

IBC_GAIA_CHANNEL_0_DENOM='ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2'
IBC_GAIA_CHANNEL_1_DENOM='ibc/C4CFF46FD6DE35CA4CF4CE031E643C8FDC9BA4B99AE598E9B0ED98FE3A2319F9'
IBC_GAIA_CHANNEL_2_DENOM='ibc/9117A26BA81E29FA4F78F57DC2BD90CD3D26848101BA880445F119B22A1E254E'
IBC_GAIA_CHANNEL_3_DENOM='ibc/A4DB47A9D3CF9A068D454513891B526702455D3EF08FB9EB558C561F9DC2B701'

# INTEGRATION TEST IBC DENOM
IBC_ATOM_DENOM=$IBC_GAIA_CHANNEL_0_DENOM

# CHAIN PARAMS
BLOCK_TIME='1s'
STAYKING_DAY_EPOCH_DURATION="100s"
STAYKING_EPOCH_EPOCH_DURATION="40s"
HOST_DAY_EPOCH_DURATION="60s"
HOST_HOUR_EPOCH_DURATION="60s"
HOST_WEEK_EPOCH_DURATION="60s"
UNBONDING_TIME="120s"
MAX_DEPOSIT_PERIOD="30s"
VOTING_PERIOD="30s"

INITIAL_ANNUAL_PROVISIONS="10000000000000.000000000000000000"
VAL_TOKENS=5000000000000
STAKE_TOKENS=5000000000
ADMIN_TOKENS=1000000000

# CHAIN MNEMONICS
VAL_MNEMONIC_1="close soup mirror crew erode defy knock trigger gather eyebrow tent farm gym gloom base lemon sleep weekend rich forget diagram hurt prize fly"
VAL_MNEMONIC_2="turkey miss hurry unable embark hospital kangaroo nuclear outside term toy fall buffalo book opinion such moral meadow wing olive camp sad metal banner"
VAL_MNEMONIC_3="tenant neck ask season exist hill churn rice convince shock modify evidence armor track army street stay light program harvest now settle feed wheat"
VAL_MNEMONIC_4="tail forward era width glory magnet knock shiver cup broken turkey upgrade cigar story agent lake transfer misery sustain fragile parrot also air document"
VAL_MNEMONIC_5="crime lumber parrot enforce chimney turtle wing iron scissors jealous indicate peace empty game host protect juice submit motor cause second picture nuclear area"
VAL_MNEMONICS=("$VAL_MNEMONIC_1","$VAL_MNEMONIC_2","$VAL_MNEMONIC_3","$VAL_MNEMONIC_4","$VAL_MNEMONIC_5")
REV_MNEMONIC="tonight bonus finish chaos orchard plastic view nurse salad regret pause awake link bacon process core talent whale million hope luggage sauce card weasel"

# STAYKING
STAYKING_CHAIN_ID=STAYKING
STAYKING_NODE_PREFIX=stayking
STAYKING_NUM_NODES=3
STAYKING_VAL_PREFIX=val
STAYKING_DENOM=$STAYKING_DENOM
STAYKING_RPC_PORT=26657
STAYKING_ADMIN_ACCT=admin
STAYKING_ADMIN_ADDRESS=sooho143umg272xger2eyurqfpjgt8u533s62mpz5weq
STAYKING_ADMIN_MNEMONIC="stock viable sponsor length claw raccoon swift hollow inspire addict order turtle forest pony talk surface harbor bargain glide proof trigger history valid penalty"
STAYKING_FEE_ADDRESS=sooho1ckh2w55t4jkdz6ck74a32mqxeyrlt8ee2ws444 #wire act strong despair apple elite glide industry journey final finish coconut repair judge payment error soul bounce public sport flee library employ position

# Binaries are contigent on whether we're doing an upgrade or not
if [[ "$UPGRADE_NAME" == "" ]]; then 
  STAYKING_CMD="$SCRIPT_DIR/../build/staykingd"
else
  STAYKING_CMD="$UPGRADES/binaries/staykingd1"
fi
STAYKING_MAIN_CMD="$STAYKING_CMD --home $SCRIPT_DIR/state/${STAYKING_NODE_PREFIX}1"

# GAIA 
GAIA_CHAIN_ID=GAIA
GAIA_NODE_PREFIX=gaia
GAIA_NUM_NODES=1
GAIA_CMD="$SCRIPT_DIR/../build/gaiad"
GAIA_VAL_PREFIX=gval
GAIA_REV_ACCT=grev1
GAIA_ADDRESS_PREFIX=cosmos
GAIA_DENOM=$ATOM_DENOM
GAIA_RPC_PORT=26557
#GAIA_RPC_PORT=26560
GAIA_MAIN_CMD="$GAIA_CMD --home $SCRIPT_DIR/state/${GAIA_NODE_PREFIX}1"
GAIA_RECEIVER_ADDRESS='cosmos1g6qdx6kdhpf000afvvpte7hp0vnpzapuyxp8uf'

# HERMES
HERMES_CMD="$SCRIPT_DIR/../build/hermes/release/hermes --config $STATE/hermes/config.toml"
HERMES_EXEC="$DOCKER_COMPOSE run --rm hermes hermes"

HERMES_STAYKING_ACCT=hrly1
HERMES_GAIA_ACCT=hrly2

HERMES_STAYKING_MNEMONIC="alter old invest friend relief slot swear pioneer syrup economy vendor tray focus hedgehog artist legend antenna hair almost donkey spice protect sustain increase"
HERMES_GAIA_MNEMONIC="resemble accident lake amateur physical jewel taxi nut demand magnet person blanket trip entire awkward fiber usual current index limb lady lady depart train"

# RELAYER
RELAYER_CMD="$SCRIPT_DIR/../build/relayer --home $STATE/relayer"
RELAYER_GAIA_EXEC="$DOCKER_COMPOSE run --rm relayer-gaia"

RELAYER_STAYKING_ACCT=rly1
RELAYER_GAIA_ACCT=rly2
HOST_RELAYER_ACCTS=($RELAYER_GAIA_ACCT)

RELAYER_GAIA_MNEMONIC="fiction perfect rapid steel bundle giant blade grain eagle wing cannon fever must humble dance kitchen lazy episode museum faith off notable rate flavor"
RELAYER_MNEMONICS=("$RELAYER_GAIA_MNEMONIC")

STAYKING_ADDRESS() {
  $STAYKING_MAIN_CMD keys show ${STAYKING_VAL_PREFIX}1 --keyring-backend test -a
}
GAIA_ADDRESS() { 
  $GAIA_MAIN_CMD keys show ${GAIA_VAL_PREFIX}1 --keyring-backend test -a 
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