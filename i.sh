
# Step1. gaia -> stayking atom 보내기
gaiad tx ibc-transfer transfer transfer channel-0 sooho143umg272xger2eyurqfpjgt8u533s62mpz5weq 1000000000uatom --from gval1 --keyring-backend test --chain-id GAIA --fees 1000uatom --gas auto --node http://gaia1:26657

# Step2. stayking 에서 ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2(ATOM 용) denom 밸런스 확인
staykingd query bank balances sooho143umg272xger2eyurqfpjgt8u533s62mpz5weq --node http://stayking1:26657

# Step3. stayking 에서 host chain gaia 에 staking 진행
staykingd tx stakeibc liquid-stake 10000 uatom --from admin --node http://stayking1:26657

# Step4. 처리된 tx log 확인
staykingd query tx {TxHash} --node http://stayking1:26657

# 이때 상태에 따라 Transfer > Delegation 되는 과정을 records 에서 볼 수 있음


