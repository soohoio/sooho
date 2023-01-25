#!/bin/bash

set -eu
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
echo $SCRIPT_DIR
pkill -f "docker-compose .*stayking.* logs" | true
pkill -f "/bin/bash.*create_logs.sh" | true
pkill -f "tail .*.log" | true
docker stop $(docker ps -aq)
docker rm $(docker ps -aq)
docker image rm $(docker image ls -qa)
