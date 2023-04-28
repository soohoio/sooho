#!/bin/bash

set -eu
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

source ${SCRIPT_DIR}/config.sh



for chain_id in ${HOST_CHAINS[@]}; do
  relayer_exec=$(GET_VAR_VALUE RELAYER_${chain_id}_EXEC)
  chain_name=$(printf "$chain_id" | awk '{ print tolower($0) }')
  chain_name_testnet=${chain_name}Testnet
  echo $chain_name
  echo $chain_id
  account_name=$(GET_VAR_VALUE RELAYER_${chain_id}_ACCT)
  mnemonic=$(GET_VAR_VALUE     RELAYER_${chain_id}_MNEMONIC)
  coin_type=$(GET_VAR_VALUE    ${chain_id}_COIN_TYPE)
  relayer_logs=${LOGS}/relayer-${chain_name_testnet}.log
  relayer_config=$STATE/relayer-${chain_name_testnet}/config
  echo $coin_type


  mkdir -p $relayer_config
  cp ${SCRIPT_DIR}/templates/relayer_${chain_name}_config.yaml $relayer_config/config.yaml
  echo $relayer_exec

  if [[ $# -ne 0 && $1 = "i" ]]; then

    RELAYER_CONFIG_FILE="$STATE/relayer-$chain_name_testnet/config/config.yaml"

    echo "Adding Relayer keys..."
   $relayer_exec rly keys restore stayking $RELAYER_STAYKING_ACCT "$RELAYER_STAYKING_MNEMONIC" >> $relayer_logs 2>&1
   $relayer_exec rly keys restore $chain_name_testnet $account_name "$RELAYER_STAYKING_MNEMONIC" --coin-type $coin_type >> $relayer_logs 2>&1

    echo "Done"

  fi

  printf "STAYKING <> $chain_name_testnet - Creating client, connection, and transfer channel..." | tee -a $relayer_logs
  $relayer_exec rly tx link stayking-${chain_name_testnet} --override >> $relayer_logs 2>&1
  echo "Done"

  printf "STAYKING <> $chain_name_testnet Testnet"

  $DOCKER_COMPOSE up -d relayer-${chain_name_testnet}
  $DOCKER_COMPOSE logs -f relayer-${chain_name_testnet} | sed -r -u "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" >> $relayer_logs 2>&1 &

done

