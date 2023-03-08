#!/bin/bash

set -eu
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source ${SCRIPT_DIR}/config.sh

initialize=""

echo
  PS3="초기화 모드를 선택하셨습니다 계속 실행하시겠습니까? : "
  COLUMNS=20
  options=(
    "Yes"
    "No"
  )
  select yn in "${options[@]}"; do
      case $yn in
          "Yes") echo "chain init mode";
                # cleanup any stale state
                rm -rf $STATE $LOGS
                mkdir -p $STATE
                mkdir -p $LOGS
                initialize="i"
                break;;
          "No" ) echo "skip....."
                break;;
      esac
  done

# Initialize the state for each chain
if [ $initialize = "i" ]
then
    for chain in STAYKING ${HOST_CHAINS[@]}; do
        bash $SRC/init_chain.sh $chain
    done
fi

# If we're testing an upgrade, setup cosmovisor
if [[ "$UPGRADE_NAME" != "" ]]; then
    printf "\n>>> UPGRADE ENABLED! ($UPGRADE_NAME)\n\n"

    # Update binary #2 with the binary that was just compiled
    mkdir -p $UPGRADES/binaries
    rm -f $UPGRADES/binaries/staykingd2
    cp $SCRIPT_DIR/../build/staykingd $UPGRADES/binaries/staykingd2

    # Build a cosmovisor image with the old binary and replace the stayking docker image with a new one
    #  that has both binaries and is running cosmovisor
    # The reason for having a separate cosmovisor image is so we can cache the building of cosmovisor and the old binary
    echo "Building Cosmovisor..."
    docker build \
        -t soohoio:cosmovisor \
        --build-arg old_commit_hash=$UPGRADE_OLD_COMMIT_HASH \
        --build-arg stayking_admin_address=$STAYKING_ADMIN_ADDRESS \
        -f $UPGRADES/Dockerfile.cosmovisor .

    echo "Re-Building StayKing with Upgrade Support..."
    docker build \
        -t soohoio:stayking \
        --build-arg upgrade_name=$UPGRADE_NAME \
        -f $UPGRADES/Dockerfile.stayking .
    echo "Done"
fi

# Start the chain and create the transfer channels
echo "start_chain.sh executed"
bash $SRC/start_chain.sh
echo "start_relayer executed"
if [[ "$1" == *h*r* || "$1" == *r*h* ]]; then
  echo "start both go and hermes relayers"
  bash $SRC/start_go_relayers.sh
  bash $SRC/start_hermes_relayers.sh
elif [[ "$1" == *r* ]]; then
  echo "start go relayers ..."
  bash $SRC/start_go_relayers.sh
elif [[ "$1" == *h* ]]; then
  echo "start hermes relayers ..."
  bash $SRC/start_hermes_relayers.sh

fi

#Register all host zones
#for i in ${!HOST_CHAINS[@]}; do
#    bash $SRC/register_host.sh ${HOST_CHAINS[$i]} $i
#done

#Create logs
$SRC/create_logs.sh &
