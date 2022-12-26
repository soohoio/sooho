#!/bin/bash

set -eu 
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

STRIDE_CHAIN_ID="$1"
STATE=$SCRIPT_DIR/../state
LOGS=$SCRIPT_DIR/../logs
KEYS_LOGS=$LOGS/keys.log

# CHAIN PARAMS
BLOCK_TIME='5s'
STRIDE_DAY_EPOCH_DURATION="120s"
STRIDE_EPOCH_EPOCH_DURATION="120s"
MAX_DEPOSIT_PERIOD="30s"
VOTING_PERIOD="30s"
INITIAL_ANNUAL_PROVISIONS="10000000000000.000000000000000000"
VAL_TOKENS=5000000000000
STAKE_TOKENS=5000000000
ADMIN_TOKENS=1000000000
PEER_PORT=26656
RPC_PORT=26657
VAL_PREFIX=val
DENOM=ustay

NODE_NAME="stride1"
STRIDE_VAL_MNEMONIC="close soup mirror crew erode defy knock trigger gather eyebrow tent farm gym gloom base lemon sleep weekend rich forget diagram hurt prize fly"
HERMES_MNEMONIC="alter old invest friend relief slot swear pioneer syrup economy vendor tray focus hedgehog artist legend antenna hair almost donkey spice protect sustain increase"
RELAYER_MNEMONIC="pride narrow breeze fitness sign bounce dose smart squirrel spell length federal replace coral lunar thunder vital push nuclear crouch fun accident hood need"

STRIDE_ADMIN_ACCT=admin
STRIDE_VAL_ACCT=val1
HERMES_ACCT=hrly1
RELAYER_ACCT=rly1

CMD="$SCRIPT_DIR/../../build/strided --home $SCRIPT_DIR/../state/stride1"

echo "Initializing Stride chain..."

# Create a state directory for the current node and initialize the chain
mkdir -p $STATE/$NODE_NAME
$CMD init STRIDE --chain-id $STRIDE_CHAIN_ID --overwrite &> /dev/null

# Update node networking configuration 
config_toml="${STATE}/${NODE_NAME}/config/config.toml"
client_toml="${STATE}/${NODE_NAME}/config/client.toml"
app_toml="${STATE}/${NODE_NAME}/config/app.toml"
genesis_json="${STATE}/${NODE_NAME}/config/genesis.json"

sed -i -E "s|cors_allowed_origins = \[\]|cors_allowed_origins = [\"\*\"]|g" $config_toml
sed -i -E "s|127.0.0.1|0.0.0.0|g" $config_toml
sed -i -E "s|timeout_commit = \"5s\"|timeout_commit = \"${BLOCK_TIME}\"|g" $config_toml
sed -i -E "s|prometheus = false|prometheus = true|g" $config_toml

sed -i -E "s|minimum-gas-prices = \".*\"|minimum-gas-prices = \"0${DENOM}\"|g" $app_toml
sed -i -E '/\[api\]/,/^enable = .*$/ s/^enable = .*$/enable = true/' $app_toml
sed -i -E 's|unsafe-cors = .*|unsafe-cors = true|g' $app_toml

sed -i -E "s|chain-id = \"\"|chain-id = \"${STRIDE_CHAIN_ID}\"|g" $client_toml
sed -i -E "s|keyring-backend = \"os\"|keyring-backend = \"test\"|g" $client_toml
sed -i -E "s|node = \".*\"|node = \"tcp://localhost:$RPC_PORT\"|g" $client_toml

sed -i -E "s|\"stake\"|\"${DENOM}\"|g" $genesis_json

# Get the endpoint and node ID
node_id=$($CMD tendermint show-node-id)@$NODE_NAME:$PEER_PORT
echo "Node ID: $node_id"

# add a validator account
echo "$STRIDE_VAL_MNEMONIC" | $CMD keys add $STRIDE_VAL_ACCT --recover --keyring-backend=test >> $KEYS_LOGS 2>&1
VAL_ADDR=$($CMD keys show $STRIDE_VAL_ACCT --keyring-backend test -a)
# Add this account to the current node
$CMD add-genesis-account ${VAL_ADDR} ${VAL_TOKENS}${DENOM}
# actually set this account as a validator on the current node 
$CMD gentx $STRIDE_VAL_ACCT ${STAKE_TOKENS}${DENOM} --chain-id $STRIDE_CHAIN_ID --keyring-backend test &> /dev/null

# Cleanup from seds
rm -rf ${client_toml}-E
rm -rf ${genesis_json}-E
rm -rf ${app_toml}-E

# add Hermes and relayer accounts on Stride
echo "$HERMES_MNEMONIC" | $CMD keys add $HERMES_ACCT --recover --keyring-backend=test >> $KEYS_LOGS 2>&1
echo "$RELAYER_MNEMONIC" | $CMD keys add $RELAYER_ACCT --recover --keyring-backend=test >> $KEYS_LOGS 2>&1
HERMES_ADDRESS=$($CMD keys show $HERMES_ACCT --keyring-backend test -a)
RELAYER_ADDRESS=$($CMD keys show $RELAYER_ACCT --keyring-backend test -a)

# give relayer accounts token balances
$CMD add-genesis-account ${HERMES_ADDRESS} ${VAL_TOKENS}${DENOM}
$CMD add-genesis-account ${RELAYER_ADDRESS} ${VAL_TOKENS}${DENOM}

# add the stayking admin account
echo "$STRIDE_ADMIN_MNEMONIC" | $CMD keys add $STRIDE_ADMIN_ACCT --recover --keyring-backend=test >> $KEYS_LOGS 2>&1
STRIDE_ADMIN_ADDRESS=$($CMD keys show $STRIDE_ADMIN_ACCT --keyring-backend test -a)
$CMD add-genesis-account ${STRIDE_ADMIN_ADDRESS} ${ADMIN_TOKENS}${DENOM}

# now we process gentx txs on the main node
$CMD collect-gentxs &> /dev/null

# wipe out the persistent peers for the main node (these are incorrectly autogenerated for each validator during collect-gentxs)
sed -i -E "s|persistent_peers = .*|persistent_peers = \"\"|g" $config_toml

# update params
jq '(.app_state.epochs.epochs[] | select(.identifier=="day") ).duration = $epochLen' --arg epochLen $STRIDE_DAY_EPOCH_DURATION $genesis_json > json.tmp && mv json.tmp $genesis_json
jq '(.app_state.epochs.epochs[] | select(.identifier=="stayking_epoch") ).duration = $epochLen' --arg epochLen $STRIDE_EPOCH_EPOCH_DURATION $genesis_json > json.tmp && mv json.tmp $genesis_json
jq '.app_state.gov.deposit_params.max_deposit_period = $newVal' --arg newVal "$MAX_DEPOSIT_PERIOD" $genesis_json > json.tmp && mv json.tmp $genesis_json
jq '.app_state.gov.voting_params.voting_period = $newVal' --arg newVal "$VOTING_PERIOD" $genesis_json > json.tmp && mv json.tmp $genesis_json

# Cleanup from seds
rm -rf ${config_toml}-E
rm -rf ${genesis_json}-E