#!/bin/bash

# insert host zone
staykingd tx levstakeibc register-host-zone connection-0 cosmos uatom ibc/0429A217F7AFD21E67CABA80049DD56BB0380B77E9C58C831366D6626D42F399 channel-0 1 --from val --keyring-backend test
# get specific host zone
staykingd q levstakeibc show-host-zone localstayking
# get all host zone
staykingd q levstakeibc list-host-zone

# GAIA > STAYKING IBC Transfer
gaiad tx ibc-transfer transfer transfer channel-0 sooho1ygs3em26qaheucpckxasxuqqej80sqt2p57nyy 1000000000uatom --from admin --keyring-backend test --chain-id gaia-localnet --fees 1000uatom --gas auto --node http://gaia1:26657
osmosisd tx ibc-transfer transfer transfer channel-0 sooho1ygs3em26qaheucpckxasxuqqej80sqt2p57nyy 1000000000uosmo --from admin --keyring-backend test --chain-id osmosis-localnet --fees 1000uosmo --gas auto --node http://osmosis1:26657
evmosd tx ibc-transfer transfer transfer channel-0 sooho1ygs3em26qaheucpckxasxuqqej80sqt2p57nyy 1000000000aevmos --from admin --keyring-backend test --chain-id evmos_9001-2 --fees 10000000000aevmos --gas auto --node http://evmos1:26657

# STAYKING BALANCE CHECK
staykingd query bank balances sooho1ygs3em26qaheucpckxasxuqqej80sqt2p57nyy --node http://stayking1:26657

# USDC UATOM PRICE UPDATE
staykingd tx records update-denom-price usdc uatom 10000 --from val1

# SEARCH DEPOSIT RECORD
staykingd q records list-deposit-record

# GAIA WITHDRWAL ICA BALANCE CHECK
gaiad query bank balances cosmos1kz28l6r539c04rnxyr5eugcsg0lyfeym8lsrz3l7xjmnr5rlvwjq09h0vr --node http://gaia1:26657

# LEVERGAE STAKE (= 1.0 Ratio)
staykingd tx levstakeibc leverage-stake 1000000 uatom 1.0 usdc --from admin
staykingd tx levstakeibc leverage-stake 1000000000 aevmos 1.0 usdc --from admin

# host zone info
staykingd q levstakeibc list-host-zone

# ICQ Query list
staykingd q interchainquery list-pending-queries


# LEVERGAE STAKE (> 1.0 Ratio)
staykingd tx levstakeibc leverage-stake 1000000 uatom 1.5 usdc --from admin

# USDC UATOM UPDATED PRICE LIST
staykingd q records list-denom-price

