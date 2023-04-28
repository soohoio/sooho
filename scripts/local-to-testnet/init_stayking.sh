#!/bin/bash

set -eu 
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source ${SCRIPT_DIR}/config.sh

IFS=',' read -r -a VAL_MNEMONICS <<< "${VAL_MNEMONICS}"

echo "Initializing Stayking chain..."
for (( i=1; i <= $NUM_NODES; i++ )); do
  NODE_NAME="${NODE_PREFIX}${i}"
  moniker=$(printf "${NODE_PREFIX}_${i}" | awk '{ print toupper($0) }')
  # Create a state directory for the current node and initialize the chain
  mkdir -p $STATE/$NODE_NAME
  CMD="$SCRIPT_DIR/../../build/staykingd --home $STATE/$NODE_NAME"
  $CMD init $moniker --chain-id $STAYKING_CHAIN_ID --overwrite &> /dev/null
  chmod -R 777 $STATE/$NODE_NAME
  # Update node networking configuration
  config_toml="${STATE}/${NODE_NAME}/config/config.toml"
  client_toml="${STATE}/${NODE_NAME}/config/client.toml"
  app_toml="${STATE}/${NODE_NAME}/config/app.toml"
  genesis_json="${STATE}/${NODE_NAME}/config/genesis.json"
  echo $config_toml

  cp $SCRIPT_DIR/templates/interest-model*.json ${STATE}/${NODE_NAME}/config/

  sed -i -E "s|cors_allowed_origins = \[\]|cors_allowed_origins = [\"\*\"]|g" $config_toml
  sed -i -E "s|127.0.0.1|0.0.0.0|g" $config_toml
  sed -i -E "s|timeout_commit = \"5s\"|timeout_commit = \"${BLOCK_TIME}\"|g" $config_toml
  sed -i -E "s|prometheus = false|prometheus = true|g" $config_toml
  sed -i -E "s|max_open_connections = 900|max_open_connections = 3900|g" $config_toml

  sed -i -E "s|max-open-connections = 1000|max-open-connections = 4096|g" $app_toml
  sed -i -E "s|minimum-gas-prices = \".*\"|minimum-gas-prices = \"0${DENOM}\"|g" $app_toml
  sed -i -E '/\[api\]/,/^enable = .*$/ s/^enable = .*$/enable = true/' $app_toml
  sed -i -E '/\[api\]/,/^swagger = .*$/ s/^swagger = .*$/swagger = true/' $app_toml
  sed -i -E 's|unsafe-cors = .*|unsafe-cors = true|g' $app_toml
  sed -i -E "s|snapshot-interval = 0|snapshot-interval = 100|g" $app_toml

  sed -i -E "s|chain-id = \"\"|chain-id = \"${STAYKING_CHAIN_ID}\"|g" $client_toml
  sed -i -E "s|keyring-backend = \"os\"|keyring-backend = \"test\"|g" $client_toml
  sed -i -E "s|node = \".*\"|node = \"tcp://localhost:$RPC_PORT\"|g" $client_toml

  sed -i -E "s|\"stake\"|\"${DENOM}\"|g" $genesis_json

# Get the endpoint and node ID

  node_id=$($CMD tendermint show-node-id)@$NODE_NAME:$PEER_PORT
  echo "Node ID: $node_id"

  # add a validator account
  val_acct="${VAL_PREFIX}${i}"
  val_mnemonic="${VAL_MNEMONICS[((i-1))]}"
  echo "$val_mnemonic" | $CMD keys add $val_acct --recover --keyring-backend=test >> $KEYS_LOGS 2>&1
  VAL_ADDR=$($CMD keys show $val_acct --keyring-backend test -a)
  # Add this account to the current node
  $CMD add-genesis-account ${VAL_ADDR} ${VAL_TOKENS}${DENOM}
# actually set this account as a validator on the current node 
  $CMD gentx $val_acct ${STAKE_TOKENS}${DENOM} --chain-id $STAYKING_CHAIN_ID --keyring-backend test &> /dev/null

# Cleanup from seds
  rm -rf ${client_toml}-E
  rm -rf ${genesis_json}-E
  rm -rf ${app_toml}-E
  if [ $i -eq $MAIN_ID ]; then
      MAIN_NODE_NAME=$NODE_NAME
      MAIN_NODE_CMD=$CMD
      MAIN_NODE_ID=$node_id
      MAIN_CONFIG=$config_toml
      MAIN_GENESIS=$genesis_json
  else
    # also add this account and it's genesis tx to the main node
    $MAIN_NODE_CMD add-genesis-account ${VAL_ADDR} ${VAL_TOKENS}${DENOM}
    cp ${STATE}/${NODE_NAME}/config/gentx/*.json ${STATE}/${MAIN_NODE_NAME}/config/gentx/

    # and add each validator's keys to the first state directory
    echo "$val_mnemonic" | $MAIN_NODE_CMD keys add $val_acct --recover --keyring-backend=test &> /dev/null
  fi
done

# add the stayking admin account
echo "$STAYKING_ADMIN_MNEMONIC" | $MAIN_NODE_CMD keys add $STAYKING_ADMIN_ACCT --recover --keyring-backend=test >> $KEYS_LOGS 2>&1
STAYKING_ADMIN_ADDRESS=$($MAIN_NODE_CMD keys show $STAYKING_ADMIN_ACCT --keyring-backend test -a)
$MAIN_NODE_CMD add-genesis-account ${STAYKING_ADMIN_ADDRESS} ${ADMIN_TOKENS}${DENOM}
$MAIN_NODE_CMD add-genesis-account "sooho19pu8c6herutnjcnqxmp6wdklmtjnrulml3vsq4" ${ADMIN_TOKENS}${DENOM}
$MAIN_NODE_CMD add-genesis-account "sooho12prkmv4cpegcnzp5yx9505cmu0ynmpz3kaffdc" ${ADMIN_TOKENS}${DENOM}
$MAIN_NODE_CMD add-genesis-account "sooho1pw0c95syjpn592ara0jp3shavaxdlhnnll2vs8" ${ADMIN_TOKENS}${DENOM}
$MAIN_NODE_CMD add-genesis-account "sooho10v2nzm6wgasg28qvukh8dp5vfqfhwyaksuefdx" ${ADMIN_TOKENS}${DENOM}
$MAIN_NODE_CMD add-genesis-account "sooho1uyrmx8zw0mxu7sdn58z29wnnqnxtqvvxh9myj5" ${ADMIN_TOKENS}${DENOM}
$MAIN_NODE_CMD add-genesis-account "sooho1z56v8wqvgmhm3hmnffapxujvd4w4rkw606qv29" ${ADMIN_TOKENS}${DENOM}
$MAIN_NODE_CMD add-genesis-account "sooho1gp5fsuud9jlmhw4rltz5zzxt6pndm8xhqndzuj" ${ADMIN_TOKENS}${DENOM}

for i in "${!HOST_RELAYER_ACCTS[@]}"; do
  RELAYER_ACCT="${HOST_RELAYER_ACCTS[i]}"
  echo $RELAYER_ACCT
  RELAYER_MNEMONIC="${RELAYER_MNEMONICS[i]}"
  echo $RELAYER_MNEMONIC
  echo "$RELAYER_MNEMONIC" | $MAIN_NODE_CMD keys add $RELAYER_ACCT --recover --keyring-backend=test >> $KEYS_LOGS 2>&1
  RELAYER_ADDRESS=$($MAIN_NODE_CMD keys show $RELAYER_ACCT --keyring-backend test -a)
  $MAIN_NODE_CMD add-genesis-account ${RELAYER_ADDRESS} ${VAL_TOKENS}${DENOM}
done

# now we process gentx txs on the main node
$MAIN_NODE_CMD collect-gentxs &> /dev/null

# wipe out the persistent peers for the main node (these are incorrectly autogenerated for each validator during collect-gentxs)
sed -i -E "s|persistent_peers = .*|persistent_peers = \"\"|g" $MAIN_CONFIG



# update params
jq '(.app_state.epochs.epochs[] | select(.identifier=="day") ).duration = $epochLen' --arg epochLen $STAYKING_DAY_EPOCH_DURATION $MAIN_GENESIS > json.tmp && mv json.tmp $MAIN_GENESIS
jq '(.app_state.epochs.epochs[] | select(.identifier=="stayking_epoch") ).duration = $epochLen' --arg epochLen $STAYKING_EPOCH_EPOCH_DURATION $MAIN_GENESIS > json.tmp && mv json.tmp $MAIN_GENESIS
jq '.app_state.staking.params.unbonding_time = $newVal' --arg newVal "$STAYKING_UNBONDING_TIME" $MAIN_GENESIS > json.tmp && mv json.tmp $MAIN_GENESIS
jq '.app_state.gov.deposit_params.max_deposit_period = $newVal' --arg newVal "$MAX_DEPOSIT_PERIOD" $MAIN_GENESIS > json.tmp && mv json.tmp $MAIN_GENESIS
jq '.app_state.gov.voting_params.voting_period = $newVal' --arg newVal "$VOTING_PERIOD" $MAIN_GENESIS > json.tmp && mv json.tmp $MAIN_GENESIS

for (( i=2; i <= $NUM_NODES; i++ )); do
    node_name="${NODE_PREFIX}${i}"
    config_toml="${STATE}/${node_name}/config/config.toml"
    genesis_json="${STATE}/${node_name}/config/genesis.json"

    # add the main node as a persistent peer
    sed -i -E "s|persistent_peers = .*|persistent_peers = \"${MAIN_NODE_ID}\"|g" $config_toml
    # copy the main node's genesis to the peer nodes to ensure they all have the same genesis
    cp $MAIN_GENESIS $genesis_json

    rm -rf ${config_toml}-E
done

# Cleanup from seds
rm -rf ${config_toml}-E
rm -rf ${genesis_json}-E
