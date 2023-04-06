#!/usr/bin/env bash

set -eo pipefail

mkdir -p ./tmp-swagger-gen
cd proto
proto_dirs=$(find ./stayking -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 2 -name '*.proto'); do
    if grep go_package "$file" &>/dev/null; then
      buf generate --template buf.gen.gogo.yaml "$file"
    fi
  done
done

cd ..
swagger-combine ./client/docs/config.json -o ./client/docs/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# clean swagger files
rm -rf ./tmp-swagger-gen
