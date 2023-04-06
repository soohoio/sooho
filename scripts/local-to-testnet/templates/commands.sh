#### SETUP HOT WALLET (Only needs to be run once)
echo "$HOT_WALLET_1_MNEMONIC" | HOST_BINARY keys add hot --recover --keyring-backend test 


#### START RELAYERS
# Create connections and channels
DOCKER_COMPOSE run --rm relayer rly transact link stayking-host

# (OR) If the go relayer isn't working, use hermes (you'll have to add the connections to the relayer config though in `STATE/relaye/config/config.yaml`)
# DOCKER_COMPOSE run --rm hermes hermes create connection --a-chain HOST_CHAIN_ID --b-chain STAYKING_CHAIN_ID
# DOCKER_COMPOSE run --rm hermes hermes create channel --a-chain STAYKING_CHAIN_ID --a-connection connection-0 --a-port transfer --b-port transfer

# Ensure Relayer Config is updated (`STATE/relayer/config/config.yaml`)
#    paths:
#     stayking-host:
#       src:
#         chain-id: stayking-1
#         client-id: 07-tendermint-0
#         connection-id: connection-0
#       dst:
#         chain-id: cosmoshub-4
#         client-id: {CLIENT-ID}
#         connection-id: {CONNECTION-ID}

# Get channel ID created on the host
build/staykingd --home STAYKING_HOME q ibc channel channels
transfer_channel=$(build/staykingd --home STAYKING_HOME q ibc channel channels | grep channel-0 -A 4 | grep counterparty -A 1 | grep channel | awk '{print $2}') && echo $transfer_channel

# Start Hermes Relayer
DOCKER_COMPOSE up -d hermes
DOCKER_COMPOSE logs -f hermes | sed -r -u "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" >> LOGS/hermes.log 2>&1 &

# Configure the Go Relayer to only run ICQ
sed -i -E "s|rule: \"\"|rule: allowlist|g" STATE/relayer/config/config.yaml

# Start Go Relayer (for ICQ)
DOCKER_COMPOSE up -d relayer
DOCKER_COMPOSE logs -f relayer | sed -r -u "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" >> LOGS/relayer.log 2>&1 &


#### REGISTER HOST
# IBC Transfer from HOST to stayking (from relayer account)
HOST_BINARY tx ibc-transfer transfer transfer $transfer_channel sooho1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 4000000HOST_DENOM --from hot --chain-id HOST_CHAIN_ID -y --keyring-backend test --node http://HOST_ENDPOINT:26657

# Confirm funds were recieved on stayking and get IBC denom
build/staykingd --home STAYKING_HOME q bank balances sooho1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8

# Register host zone
IBC_DENOM=$(build/staykingd --home STAYKING_HOME q bank balances sooho1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 | grep ibc | awk '{print $2}' | tr -d '"') && echo $IBC_DENOM
build/staykingd --home STAYKING_HOME tx stakeibc register-host-zone \
    connection-0 HOST_DENOM HOST_ACCOUNT_PREFIX $IBC_DENOM channel-0 1 \
    --from admin --gas 1000000 -y

# Add validator
build/staykingd --home STAYKING_HOME tx stakeibc add-validator HOST_CHAIN_ID HOST_VAL_NAME_1 HOST_VAL_ADDRESS_1 10 5 --chain-id STAYKING_CHAIN_ID --keyring-backend test --from admin -y

# Confirm ICA channels were registered
build/staykingd --home STAYKING_HOME q stakeibc list-host-zone

#### FLOW
## Go Through Flow
# Liquid stake (then wait and LS again)
build/staykingd --home STAYKING_HOME tx stakeibc liquid-stake 1000000 HOST_DENOM --keyring-backend test --from admin -y --chain-id STAYKING_CHAIN_ID -y

# Confirm stTokens, StakedBal, and Redemption Rate
build/staykingd --home STAYKING_HOME q bank balances sooho1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
build/staykingd --home STAYKING_HOME q stakeibc list-host-zone

# Redeem
build/staykingd --home STAYKING_HOME tx stakeibc redeem-stake 1000 HOST_CHAIN_ID HOT_WALLET_ADDRESS --from admin --keyring-backend test --chain-id STAYKING_CHAIN_ID -y

# Confirm stTokens and StakedBal
build/staykingd --home STAYKING_HOME q bank balances sooho1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
build/staykingd --home STAYKING_HOME q stakeibc list-host-zone

# Add another validator
build/staykingd --home STAYKING_HOME tx stakeibc add-validator HOST_CHAIN_ID HOST_VAL_NAME_2 HOST_VAL_ADDRESS_2 10 5 --chain-id STAYKING_CHAIN_ID --keyring-backend test --from admin -y

# Liquid stake and confirm the stake was split 50/50 between the validators
build/staykingd --home STAYKING_HOME tx stakeibc liquid-stake 1000000 HOST_DENOM --keyring-backend test --from admin -y --chain-id STAYKING_CHAIN_ID -y

# Change validator weights
build/staykingd --home STAYKING_HOME tx stakeibc change-validator-weight HOST_CHAIN_ID HOST_VAL_ADDRESS_1 1 --from admin -y
build/staykingd --home STAYKING_HOME tx stakeibc change-validator-weight HOST_CHAIN_ID HOST_VAL_ADDRESS_2 49 --from admin -y

# LS and confirm delegation aligned with new weights
build/staykingd --home STAYKING_HOME tx stakeibc liquid-stake 1000000 HOST_DENOM --keyring-backend test --from admin -y --chain-id STAYKING_CHAIN_ID -y

# Call rebalance to and confirm new delegations
build/staykingd --home STAYKING_HOME tx stakeibc rebalance-validators HOST_CHAIN_ID 5 --from admin

# Clear balances
fee_address=$(build/staykingd --home STAYKING_HOME q stakeibc show-host-zone osmosis-1 | grep feeAccount -A 1 | grep address | awk '{print $2}') && echo $fee_address
balance=$(build/osmosisd --home STAYKING_HOME q bank balances $fee_address | grep amount | awk '{print $3}' | tr -d '"') && echo $balance
build/staykingd --home STAYKING_HOME tx stakeibc clear-balance HOST_CHAIN_ID $balance $transfer_channel --from admin

# Update delegations (just submit this query and confirm the ICQ callback displays in the stayking logs)
# Must be submitted in ICQ window
build/staykingd --home STAYKING_HOME tx stakeibc update-delegation HOST_CHAIN_ID HOST_VAL_ADDRESS_1 --from admin -y

#### MISC 
# If a channel closes, restore it with:
build/staykingd --home STAYKING_HOME tx stakeibc restore-interchain-account HOST_CHAIN_ID {DELEGATION | WITHDRAWAL | FEE | REDEMPTION} --from admin