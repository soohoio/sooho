#!/bin/bash

set -eu
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
echo $SCRIPT_DIR
docker stop $(docker ps -aq)
docker rm $(docker ps -aq)
docker image rm $(docker image ls -qa)