   0 gaiad
   1 gaiad query bank balances cosmos16lmf7t0jhaatan6vnxlgv47h2wf0k5lnhvye5h
   2 netstat
   3 gaiad query bank balances cosmos16lmf7t0jhaatan6vnxlgv47h2wf0k5lnhvye5h --node http://gaia1:26657
   4 gaiad tx ibc-transfer transfer transfer $CHANNEL $CRESCENT_ADDRESS $AMOUNT --from $WALLET --keyring-backend test --chain-id $SENDER_CHAIN --fees $FEE --gas auto --node $GAIA_NODE -y
   5 gaiad tx ibc-transfer transfer transfer channel-0 stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 $AMOUNT --from $WALLET --keyring-backend test --chain-id $SENDER_CHAIN --fees 1000uatom --gas auto --node http://gaia1:26657 -y
   6 gaiad keys list
   7 gaiad tx ibc-transfer transfer transfer channel-0 stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 $AMOUNT --from gval1 --keyring-backend test --chain-id GAIA --fees 1000uatom --gas auto --node http://gaia1:26657 -y
   8 gaiad tx ibc-transfer transfer transfer channel-0 stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 1000000000 --from gval1 --keyring-backend test --chain-id GAIA --fees 1000uatom --gas auto --node http://gaia1:26657
   9 gaiad tx ibc-transfer transfer transfer channel-0 sooho143umg272xger2eyurqfpjgt8u533s62mpz5weq 1000000000uatom --from gval1 --keyring-backend test --chain-id GAIA --fees 1000uatom --gas auto --node http://gaia1:26657
  10 gaiad query ibc-transfer
  11 gaiad query tx D6313F46620AD600B3C7F23D6E400BB63AF72A147A4094C53252400A0558948D
  12 gaiad query tx D6313F46620AD600B3C7F23D6E400BB63AF72A147A4094C53252400A0558948D --node http://gaia1:26657
  13 history




   0 ll
     1 docker ps
     2 strided query bank balances stride16lmf7t0jhaatan6vnxlgv47h2wf0k5ln58y9qm --node http://stride1:26657
     3 strided tx stakeibc
     4 strided tx stakeibc liquid-stake
     5 strided tx stakeibc liquid-stake 1 atom --from
     6 strided keys
     7 strided keys list
     8 strided tx stakeibc liquid-stake 1 atom --from admin
     9 strided query
    10 strided query stakeibc
    11 strided query stakeibc module-address
    12 strided query
    13 strided query records
    14 strided query records list-deposit-record
    15 strided query records
    16 strided query bank
    17 strided query bank denom-metadata
    18 strided query bank total
    19 strided keys list
    20 staykingd query bank balances sooho143umg272xger2eyurqfpjgt8u533s62mpz5weq --node http://stayking1:26657
    21 strided query bank total stride16lmf7t0jhaatan6vnxlgv47h2wf0k5ln58y9qm --node http://stride1:26657
    22 strided tx stakeibc liquid-stake 1000 uatom --from admin
    23 strided query bank total stride16lmf7t0jhaatan6vnxlgv47h2wf0k5ln58y9qm --node http://stride1:26657
    24 strided query records list-deposit-record
    25 strided query bank total stride16lmf7t0jhaatan6vnxlgv47h2wf0k5ln58y9qm --node http://stride1:26657
    26 strided query
    27 strided query staking
    28 strided query staking pool
    29 strided query interchainquery
    30 strided query interchainquery list-pending-queries
    31 strided query interchainquery
    32 strided query
    33 strided query icacallbacks
    34 strided query
    35 hisotttdf
    36 history