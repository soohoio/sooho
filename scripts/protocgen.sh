#!/usr/bin/env bash

set -eo pipefail

# get protoc executions
go get github.com/regen-network/cosmos-proto/protoc-gen-gocosmos 2>/dev/null

# get cosmos sdk from github
# go get github.com/cosmos/cosmos-sdk@v0.45.11 2>/dev/null

echo "Generating gogo proto code"
cd proto
proto_dirs=$(find ./stride -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep go_package "$file" &>/dev/null; then
      buf generate --template buf.gen.gogo.yaml "$file"
    fi
  done
done

cd ..

# move proto files to the right places
#
# Note: Proto files are suffixed with the current binary version.
cp -r github.com/Stride-Labs/stride/v4/* ./
rm -rf github.com