#!/bin/bash

set -eu 
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source ${SCRIPT_DIR}/config.sh

# build 실행
bash ${SCRIPT_DIR}/build.sh

# 체인 초기화 실행
if [[ $# -ne 0 && $1 = "i" ]]; then
  echo
  PS3="초기화 모드를 선택하셨습니다 계속 실행하시겠습니까?"
  COLUMNS=20
  options=(
    "Yes"
    "No"
  )
  select yn in "${options[@]}"; do
      case $yn in
          "Yes")    echo "chain nit mode";         break;;
          "No" )    exit;;
      esac
  done
# 기존 STATE LOG 폴더 지우기
  rm -rf $STATE $LOGS
  mkdir -p $STATE
  mkdir -p $LOGS
  # Start stayking chain
  echo "StayKing init mode..."
  bash ${SCRIPT_DIR}/init_stayking.sh $STAYKING_CHAIN_ID
fi

for chain_id in STAYKING; do
    num_nodes=$(GET_VAR_VALUE ${chain_id}_NUM_NODES)
    node_prefix=$(GET_VAR_VALUE ${chain_id}_NODE_PREFIX)

    log_file=$LOGS/${node_prefix}.log

    echo "Starting $chain_id chain"
    nodes_names=$(i=1; while [ $i -le $num_nodes ]; do printf "%s " ${node_prefix}${i}; i=$(($i + 1)); done;)

    $DOCKER_COMPOSE up -d $nodes_names
    $DOCKER_COMPOSE logs -f ${node_prefix}1 | sed -r -u "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" > $log_file 2>&1 &
done

for chain_id in STAYKING; do
    printf "Waiting for $chain_id to start..."

    node_prefix=$(GET_VAR_VALUE ${chain_id}_NODE_PREFIX)
    log_file=$LOGS/${node_prefix}.log

    ( tail -f -n0 $log_file & ) | grep -q "finalizing commit of block"
    echo "Done"
done

sleep 5

if [[ $# -ne 0 && $1 = "i" ]]; then
  echo "add relayer keys and start relayers !"
  bash $SCRIPT_DIR/start_relayers.sh $1
else
  echo "start relayers !"
  bash $SCRIPT_DIR/start_relayers.sh
fi

echo "register host !"
bash $SCRIPT_DIR/register_host.sh

echo "create logs !"
$SCRIPT_DIR/create_logs.sh &

echo "Done"
