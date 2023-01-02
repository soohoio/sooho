
# Step1. gaia -> stayking atom 보내기
gaiad tx ibc-transfer transfer transfer channel-0 sooho143umg272xger2eyurqfpjgt8u533s62mpz5weq 1000000000uatom --from gval1 --keyring-backend test --chain-id GAIA --fees 1000uatom --gas auto --node http://gaia1:26657

# Step2. stayking 에서 ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2(ATOM 용) denom 밸런스 확인
staykingd query bank balances sooho143umg272xger2eyurqfpjgt8u533s62mpz5weq --node http://stayking1:26657

# Step3. stayking 에서 host chain gaia 에 staking 진행
staykingd tx stakeibc liquid-stake 10000 uatom --from admin --node http://stayking1:26657

# Step4. 처리된 tx log 확인
staykingd query tx {TxHash} --node http://stayking1:26657

# 이때 상태에 따라 Transfer > Delegation 되는 과정을 records 에서 볼 수 있음


# Unbonding

#Step 1. Redeem stake
staykingd tx stakeibc redeem-stake 1000 GAIA cosmos1uk4ze0x4nvh4fk0xm4jdud58eqn4yxhrgl2scj --chain-id STAYKING --from admin --keyring-backend test --node http://stayking1:26657

#Step 2. check tx log
staykingd q tx {TxHash} --node http://stayking1:26657

#Step 3. check records ( records statusr가 CLAIMABLE이 될 때까지 대기)
staykingd q records list-epoch-unbonding-record --node http://stayking1:26657
staykingd q records list-user-redemption-record --node http://stayking1:26657

#Step 4.Claimable한 undelegate token에 대해서 claim 요청
staykingd tx stakeibc claim-undelegated-tokens GAIA {epochNumber} sooho143umg272xger2eyurqfpjgt8u533s62mpz5weq --from admin --node http://stayking1:26657

#Step5. unstaking요청한 만큼 uatom 들어온 것 확인
gaiad q bank balances cosmos1uk4ze0x4nvh4fk0xm4jdud58eqn4yxhrgl2scj --node http://gaia1:26657

#Step 6. stuatom이 unstaking 요청한 양만큼 사라진 것을 확인
staykingd q bank balances sooho143umg272xger2eyurqfpjgt8u533s62mpz5weq --node http://stayking1:26657