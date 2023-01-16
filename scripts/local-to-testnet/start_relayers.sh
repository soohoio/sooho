#!/bin/bash

set -eu
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

source ${SCRIPT_DIR}/config.sh

RELAYER_CONFIG_FILE="$STATE/relayer-$chain_name/config/config.yaml"

mkdir -p $relayer_config
cp ${SCRIPT_DIR}/templates/relayer_config.yaml $relayer_config/config.yaml

echo "Adding Relayer keys..."
RELAYER_GAIA_TESTNET_EXEC="$DOCKER_COMPOSE run --rm relayer-gaiaTestnet"
RELAYER_EXEC=$RELAYER_GAIA_TESTNET_EXEC
RELAYER_CMD="$SCRIPT_DIR/../../build/relayer --home $STATE/relayer"

$RELAYER_EXEC rly keys restore stayking rly1 "$RELAYER_STAYKING_MNEMONIC" >> $relayer_logs 2>&1
$RELAYER_EXEC rly keys restore gaiaTestnet rly2 "$RELAYER_STAYKING_MNEMONIC" >> $relayer_logs 2>&1
echo "Done"

printf "STAYKING <> $chain_name - Creating client, connection, and transfer channel..." | tee -a $relayer_logs

$RELAYER_EXEC rly tx link stayking-${chain_name} --client-tp 24h >> $relayer_logs 2>&1
echo "Done"
printf "STAYKING <> GAIA Testnet"

$DOCKER_COMPOSE up -d relayer-${chain_name}
$DOCKER_COMPOSE logs -f relayer-${chain_name} | sed -r -u "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" >> $relayer_logs 2>&1 &
