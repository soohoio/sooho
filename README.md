![Logo](https://user-images.githubusercontent.com/27390552/223629157-bb68ad1a-3222-4450-859d-0e797a86d9f3.png)
# Cross Zone Leverage Staking 


[Twitter](https://twitter.com/staykingg) | [Service](https://stayking.xyz/)


## What is StayKing?

We are currently building AppChain DeFi products in the Cosmos ecosystem.

- First, we provide "Auto-Compounding" feature based on highly effective assets in the Cosmos 50+ Zones
- Second, we build our own lending pool and provide leverage staking through token swap
- Third, we provide "Liquid Unstaking" feature to secure quick liquidity for users after staking.

It is the mission of the StayKing team to build Web 3.0 that makes all these features easy and friendly for users.


## How Does It Work?

- To be written


## Getting Started

### Prerequisite

#### Installing StayKing

To install the latest version of StayKing blockchain node binary, execute the following command on your machine:

```
git clone https://github.com/soohoio/stayking
```

#### Set up local development mode
Install the required git submodule dependencies (gaia, relayer, etc)
```
git submodule update --init
```

### Developing on StayKing

Since all execution environments are based on Docker, docker and docker-compose must be installed beforehand.
If it is not installed, please install them first.

There are three modes of execution for developing as a developer.

#### 1. StayKing Local Node
The first is to launch only StayKing AppChain in the local environment. you can check that source files for navigating to the "testutil/localstayking" directory and 
you execute the command with Makefile for running a node

```
make localnet-start
```

If it failed to start, you execute `make localnet-clean` command first. The most likely cause of the error is that the chain update has changed the contents of the genesis file in the past.

If successful, you'll see the genesis block successfully created as shown below

```
localnet-staykingd-1  | 8:27AM INF Found height=0 index=0 module=consensus wal=/home/stayking/.stayking/data/cs.wal/wal
localnet-staykingd-1  | 8:27AM INF Catchup by replaying consensus messages height=1 module=consensus
localnet-staykingd-1  | 8:27AM INF Replay: Done module=consensus
localnet-staykingd-1  | 8:27AM INF service start impl=Evidence module=evidence msg={}
localnet-staykingd-1  | 8:27AM INF service start impl=StateSync module=statesync msg={}
localnet-staykingd-1  | 8:27AM INF service start impl=PEX module=pex msg={}
localnet-staykingd-1  | 8:27AM INF service start book=/home/stayking/.stayking/config/addrbook.json impl=AddrBook module=p2p msg={}
localnet-staykingd-1  | 8:27AM INF Saving AddrBook to file book=/home/stayking/.stayking/config/addrbook.json module=p2p size=0
localnet-staykingd-1  | 8:27AM INF Ensure peers module=pex numDialing=0 numInPeers=0 numOutPeers=0 numToDial=10
localnet-staykingd-1  | 8:27AM INF No addresses to dial. Falling back to seeds module=pex
localnet-staykingd-1  | 8:27AM INF serve module=api-server msg={}
localnet-staykingd-1  | 8:27AM INF Timed out dur=4983.850731 height=1 module=consensus round=0 step=1
localnet-staykingd-1  | 8:27AM INF received proposal module=consensus proposal={"Type":32,"block_id":{"hash":"57D3CEC28A875E2B05F4E6D1DFB2B00AE3EA764116231DC8C25665776DB148C9","parts":{"hash":"64A79CA2353AE8E1B7257F9200CE4B9BE17CBB9ED16659AEA8EA3EEAC6E5145B","total":1}},"height":1,"pol_round":-1,"round":0,"signature":"VLQpbIztL4wSvdZ8xVQWa8ZawGSEVSpBUktaq14RDj6XcuTqwBuDsirMevtDIcxijsTHI0DEJhQGi3g32vu6Ag==","timestamp":"2023-03-08T08:27:10.776724102Z"}
localnet-staykingd-1  | 8:27AM INF received complete proposal block hash=57D3CEC28A875E2B05F4E6D1DFB2B00AE3EA764116231DC8C25665776DB148C9 height=1 module=consensus
localnet-staykingd-1  | 8:27AM INF finalizing commit of block hash={} height=1 module=consensus num_txs=0 root=E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855
localnet-staykingd-1  | 8:27AM INF initial day epoch module=x/epochs
localnet-staykingd-1  | 8:27AM INF Initiating all host zone unbondings for epoch 1... module=x/stakeibc
localnet-staykingd-1  | 8:27AM INF Sweeping All Unbonded Tokens... module=x/stakeibc
```

As you may have noticed, the execution environment for this node is run as a shell script through the docker-compose.yml setting.

Since our environment mostly consists of sending messages between chains in IBC packets, it is unlikely that the environment provided will be used often, but it could be used to quickly verify that the StayKing chain's binary is working properly.


#### 2. StayKing LocalNet with HostChain and Go-Relayer

The second is to launch StayKing AppChain with other host chains (Gaia, Evmos) and Relayer for the LocalNet environment. 
you can also check it for navigating to the "dockernet" directory and to configure this environment, you execute the command with Makefile for running a node.

To test the chain with several mnemonics for operating at least 2 nodes and 1 relayer and you need to write `dockernet/keys.sh` file with `dockernet/keys.sample.sh` as a reference.

```
set -eu

# CHAIN MNEMONICS
VAL_MNEMONIC_1="fill this out"
VAL_MNEMONIC_2="fill this out"
VAL_MNEMONIC_3="fill this out"
VAL_MNEMONIC_4="fill this out"
VAL_MNEMONIC_5="fill this out"
VAL_MNEMONICS=("$VAL_MNEMONIC_1","$VAL_MNEMONIC_2","$VAL_MNEMONIC_3","$VAL_MNEMONIC_4","$VAL_MNEMONIC_5")
REV_MNEMONIC="fill this out"

STAYKING_ADMIN_MNEMONIC="fill this out"

RELAYER_GAIA_MNEMONIC="fill this out"

```

Once you've filled in the mnemonics above and run the command below, the chains and a relayer will start running.

```
make start-docker build=sgr
```

1. `s` This will re-build the StayKing binary (default)
2. `g` This will re-build the Gaia binary
3. `e` This will re-build the Evmos binary (Comming Soon)
4. `r` This will re-build the Go-Relayer binary
5. `h` This will re-build the Hermes binary

To bring down the chain, execute:

```
make stop-docker
```

To bring clean the chain, execute:

```
make clean-docker
```
If the docker build completes properly, you can choose to initialize the genesis file as shown below.
```
Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
Done

1) Yes
2) No
초기화 모드를 선택하셨습니다 계속 실행하시겠습니까? : 1
chain init mode
Initializing STAYKING chain...
Node #1 ID: 7d20dc18436f03cba1c9ce34f8e5ef952bc6211a@stayking1:26656
Initializing GAIA chain...
Node #1 ID: 5264f8ef4bb704e51c4f72bbb2e7af1f09043a9e@gaia1:26656
start_chain.sh executed

```
1) Yes, Reset genesis, config and data files
2) No, Keep genesis, config and data files before created


If successful, you can check process running with CLI `docker ps`

<img width="1774" alt="image" src="https://user-images.githubusercontent.com/27390552/223674843-06804de2-9235-4fca-a05e-e2142e1dd2af.png">

And you can check the status of each node through the logs

```
# StayKing log
tail -f dockernet/logs/stayking.log

# Gaia log
tail -f dockernet/logs/gaia.log

# Relayer.log
tail -f dockernet/logs/relayer-gaia.log

```

#### 3. StayKing Build / Run scripts for Testnet and Mainnet

navigate to `scripts/local-to-testnet` for Testnet, `scripts/local-to-mainnet` for Mainnet

To be written, continue...


## Our Project is
Our project is inspired by Stride known as "Liquid Staking Service" in the Cosmos Zone, and most of custom modules such as x/stakeibc, x/records, etc are borrowed from it. Thank Stride Team for being a pioneer and developing this technology.