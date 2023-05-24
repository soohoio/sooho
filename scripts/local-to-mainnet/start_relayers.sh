#!/bin/bash

set -eu
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

source ${SCRIPT_DIR}/config.sh

for chain_id in ${HOST_CHAINS[@]}; do
  relayer_exec=$(GET_VAR_VALUE RELAYER_${chain_id}_EXEC)
  chain_name=$(printf "$chain_id" | awk '{ print tolower($0) }')
  chain_name_mainnet=${chain_name}Mainnet
  echo $chain_name
  echo $chain_id
  account_name=$(GET_VAR_VALUE RELAYER_${chain_id}_ACCT)
  mnemonic=$(GET_VAR_VALUE     RELAYER_${chain_id}_MNEMONIC)
  coin_type=$(GET_VAR_VALUE    ${chain_id}_COIN_TYPE)
  relayer_logs=${LOGS}/relayer-${chain_name_mainnet}.log
  relayer_config=$STATE/relayer-${chain_name_mainnet}/config
  echo $coin_type


  mkdir -p $relayer_config
  cp ${SCRIPT_DIR}/templates/relayer_${chain_name}_config.yaml $relayer_config/config.yaml
  echo $relayer_exec

  if [[ $# -ne 0 && $1 = "i" ]]; then

    RELAYER_CONFIG_FILE="$STATE/relayer-$chain_name_mainnet/config/config.yaml"

    echo "Adding Relayer keys..."
   $relayer_exec rly keys restore stayking $RELAYER_STAYKING_ACCT "$RELAYER_STAYKING_MNEMONIC" >> $relayer_logs 2>&1
   $relayer_exec rly keys restore $chain_name_mainnet $account_name "$RELAYER_STAYKING_MNEMONIC" --coin-type $coin_type >> $relayer_logs 2>&1

    echo "Done"

  fi


  printf "STAYKING <> $chain_name_mainnet - Creating client, connection, and transfer channel..." | tee -a $relayer_logs
  #$relayer_exec rly tx link stayking-${chain_name_mainnet} --override >> $relayer_logs 2>&1
  #$relayer_exec rly tx client evmosMainnet stayking stayking-evmosMainnet --debug >> $relayer_logs 2>&1
  #$relayer_exec rly tx connection stayking-evmosMainnet --timeout 30s --override --debug >> $relayer_logs 2>&1
  $relayer_exec rly tx channel stayking-evmosMainnet --src-port transfer --dst-port transfer --order unordered --version ics20-1 --timeout 30s --max-retries 5 --debug >> $relayer_logs 2>&1
  echo "Done"

  printf "STAYKING <> $chain_name_mainnet Mainnet"

  $DOCKER_COMPOSE up -d relayer-${chain_name_mainnet}
  $DOCKER_COMPOSE logs -f relayer-${chain_name_mainnet} | sed -r -u "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" >> $relayer_logs 2>&1 &

done
