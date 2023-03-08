#!/bin/bash

set -eu
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

source ${SCRIPT_DIR}/config.sh
echo $SCRIPT_DIR
for chain_id in ${HOST_CHAINS[@]}; do
    hermes_exec=$(GET_VAR_VALUE HERMES_${chain_id}_EXEC)
    chain_name=$(printf "$chain_id" | awk '{ print tolower($0) }')
    echo $chain_name
    host_chain_id=$(GET_VAR_VALUE HOST_CHAIN_ID)
    mkdir -p $hermes_config
    cp ${SCRIPT_DIR}/templates/hermes_config.toml $hermes_config/config.toml
    echo "$RELAYER_GAIA_TESTNET_MNEMONIC" > ${SCRIPT_DIR}/templates/mnemonic_file_hub
    cp ${SCRIPT_DIR}/templates/mnemonic_file_hub $hermes_config/mnemonic_file_hub
    rm ${SCRIPT_DIR}/templates/mnemonic_file_hub

    printf "STAYKING <> $chain_id - Adding relayer keys..."
    $hermes_exec hermes keys add --key-name rly1 --chain $STAYKING_CHAIN_ID --mnemonic-file /home/hermes/.hermes/mnemonic_file_hub
    $hermes_exec hermes keys add --key-name rly2 --chain $host_chain_id --mnemonic-file /home/hermes/.hermes/mnemonic_file_hub
    rm $hermes_config/mnemonic_file_hub
    echo "Done"


    printf "STAYKING <> $chain_id - Creating client, connection, and transfer channel..." | tee -a $hermes_logs
#    $hermes_exec hermes create channel --a-chain $STAYKING_CHAIN_ID --b-chain $host_chain_id --a-port transfer --b-port transfer --new-client-connection --yes >> $hermes_logs 2>&1
    echo "Done"

    printf "STAYKING <> GAIA"

    $DOCKER_COMPOSE up -d hermes-gaiaTestnet
    $DOCKER_COMPOSE logs -f hermes-gaiaTestnet | sed -r -u "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" >> $hermes_logs 2>&1 &

done