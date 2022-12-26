#!/bin/sh

CHAIN_ID=localstayking
STAYKING_HOME=$HOME/.stayking
CONFIG_FOLDER=$STAYKING_HOME/config
MONIKER=val

MNEMONIC="deer gaze swear marine one perfect hero twice turkey symbol mushroom hub escape accident prevent rifle horse arena secret endless panel equal rely payment"

edit_genesis () {

    GENESIS=$CONFIG_FOLDER/genesis.json

    # Update staking module
    dasel put string -f $GENESIS '.app_state.staking.params.bond_denom' 'ustay'
    dasel put string -f $GENESIS '.app_state.staking.params.unbonding_time' '240s'

    # Update crisis module
    dasel put string -f $GENESIS '.app_state.crisis.constant_fee.denom' 'ustay'

    # Udpate gov module
    dasel put string -f $GENESIS '.app_state.gov.voting_params.voting_period' '60s'
    dasel put string -f $GENESIS '.app_state.gov.deposit_params.min_deposit.[0].denom' 'ustay'

    # Update epochs module
    dasel put string -f $GENESIS '.app_state.epochs.epochs.(.identifier=day).duration' '120s'
    dasel put string -f $GENESIS '.app_state.epochs.epochs.(.identifier=stride_epoch).duration' '120s'

    # Update mint module
    dasel put string -f $GENESIS '.app_state.mint.params.mint_denom' 'ustay'
    dasel put string -f $GENESIS '.app_state.mint.params.epoch_identifier' 'mint'

}

add_genesis_accounts () {

    staykingd add-genesis-account sooho1wal8dgs7whmykpdaz0chan2f54ynythkkcm264 100000000000ustay --home $STAYKING_HOME
    staykingd add-genesis-account sooho1u9klnra0d4zq9ffalpnr3nhz5859yc7cz64jkx 100000000000ustay --home $STAYKING_HOME
    staykingd add-genesis-account sooho1kwax6g0q2nwny5n43fswexgefedge033lrtjvl 100000000000ustay --home $STAYKING_HOME
    staykingd add-genesis-account sooho1dv0ecm36ywdyc6zjftw0q62zy6v3mndr63wwhu 100000000000ustay --home $STAYKING_HOME
    staykingd add-genesis-account sooho1z3dj2tvqpzy2l5shx98f9k5486tleah5fcv79f 100000000000ustay --home $STAYKING_HOME
    staykingd add-genesis-account sooho14khzkfs8luaqymdtplrt5uwzrghrndehpajr9k 100000000000ustay --home $STAYKING_HOME
    staykingd add-genesis-account sooho1qym804u6sa2gvxedfy96c0v9jc0ww7599t600r 100000000000ustay --home $STAYKING_HOME
    staykingd add-genesis-account sooho1y6pcev4mj0w0205qh5dp40f7dvcszuyyandted 100000000000ustay --home $STAYKING_HOME
    staykingd add-genesis-account sooho1tcrlyn05q9j590uauncywf26ptfn8se6q6076h 100000000000ustay --home $STAYKING_HOME
    staykingd add-genesis-account sooho14ugekxs6f4rfleg6wj8k0wegv69khfpxzuyntc 100000000000ustay --home $STAYKING_HOME
    staykingd add-genesis-account sooho12r6ylzpxsxyuzsg2v8g3jd5zfux5nqfecl98pp 100000000000ustay --home $STAYKING_HOME

    echo $MNEMONIC | staykingd keys add $MONIKER --recover --keyring-backend=test --home $STAYKING_HOME

    staykingd gentx $MONIKER 500000000ustay --keyring-backend=test --chain-id=$CHAIN_ID --home $STAYKING_HOME

    staykingd collect-gentxs --home $STAYKING_HOME
}

edit_config () {
    # Remove seeds
    dasel put string -f $CONFIG_FOLDER/config.toml '.p2p.seeds' ''

    # Expose the rpc
    dasel put string -f $CONFIG_FOLDER/config.toml '.rpc.laddr' "tcp://0.0.0.0:26657"
}

if [[ ! -d $CONFIG_FOLDER ]]
then
    echo $MNEMONIC | staykingd init -o --chain-id=$CHAIN_ID --home $STAYKING_HOME --recover $MONIKER
    edit_genesis
    add_genesis_accounts
    edit_config
fi

staykingd start --home $STAYKING_HOME
