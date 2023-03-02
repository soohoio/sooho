#!/bin/bash

set -eu 
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source ${SCRIPT_DIR}/config.sh

BUILDDIR=${SCRIPT_DIR}/../build
mkdir -p $BUILDDIR

build_local_and_docker() {
   module="$1"
   folder="$2"
   title=$(printf "$module" | awk '{ print toupper($0) }')

    printf '%s' "Building $title Locally...  "
    cwd=$PWD
    cd $folder
    GOBIN=$BUILDDIR go install -mod=readonly -trimpath ./... 2>&1 | grep -v -E "deprecated|keychain" | true
    echo $GOBIN
    local_build_succeeded=${PIPESTATUS[0]}
    cd $cwd

    if [[ "$local_build_succeeded" == "0" ]]; then
        echo "Done"
    else
        echo "Failed"
        return $local_build_succeeded
    fi

    echo "Building $title Docker...  "
    if [[ "$module" == "stayking" ]]; then
        image=Dockerfile
    else
        image=dockernet/dockerfiles/Dockerfile.$module
    fi

    DOCKER_BUILDKIT=1 docker build --tag soohoio:$module -f $image . | true
    docker_build_succeeded=${PIPESTATUS[0]}

    if [[ "$docker_build_succeeded" == "0" ]]; then
      echo "Done"
    else
      echo "Failed"
    fi
    return $docker_build_succeeded
}

build_local_hermes_and_docker () {
  printf '%s' "Building Hermes Locally... ";
  cd deps/hermes;
  cargo build --release --target-dir $BUILDDIR/hermes;
  local_build_succeeded=${PIPESTATUS[0]}
  if [[ "$local_build_succeeded" == "0" ]]; then
    echo "Done"
  else
    echo "Failed"
    return $local_build_succeeded
  fi
  cd ../..

  echo "Building Hermes Docker... ";
  DOCKER_BUILDKIT=1 docker build --tag soohoio:hermes -f dockernet/dockerfiles/Dockerfile.hermes . | true
  docker_build_succeeded=${PIPESTATUS[0]}

   if [[ "$docker_build_succeeded" == "0" ]]; then
      echo "Done"
   else
      echo "Failed"
   fi
   return $docker_build_succeeded
}

ADMINS_FILE=${SCRIPT_DIR}/../utils/admins.go
ADMINS_FILE_BACKUP=${SCRIPT_DIR}/../utils/admins.go.main

replace_admin_address() {
   cp $ADMINS_FILE $ADMINS_FILE_BACKUP
   sed -i -E "s|sooho1k8c2m5cn322akk5wy8lpt87dd2f4yh9azg7jlh|$STAYKING_ADMIN_ADDRESS|g" $ADMINS_FILE
}

revert_admin_address() {
   mv $ADMINS_FILE_BACKUP $ADMINS_FILE
   rm -f ${ADMINS_FILE}-E
}

# build docker images and local binaries
while getopts sgrh flag; do
   case "${flag}" in
      # For stayking, we need to update the admin address to one that we have the seed phrase for
      s) replace_admin_address
         if (build_local_and_docker stayking .) ; then
            revert_admin_address
         else
            revert_admin_address
            exit 1
         fi
         ;;
      g) build_local_and_docker gaia deps/gaia ;;
      r) build_local_and_docker relayer deps/relayer ;;
      h) build_local_hermes_and_docker ;;

   esac
done