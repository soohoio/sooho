
## StayKing Custom Modules

1. `lendingpool` - A pool where lenders can deposit money to earn a secure return and stakers can leverage that token for staking.
2. `leverstakeibc` - A module that provides stakers with the ability to perform leveraged staking via cross-chain IBC.
3. `stakeibc` - Manages minting and burning of stAssets, staking and unstaking of native assets across chains
4. `icacallbacks` - Callbacks for ICA(a.k.a, InterChain Accounts).
5. `records` - IBC middleware wrapping the transfer module, does record keeping on IBC transfers and ICA calls
6. `claim` - airdrop logic for StayKing's rolling, task-based airdrop
7. `interchainquery` - Issues queries between IBC chains, verifies state proof and executes callbacks.
8. `epochs` - Makes on-chain timers which other modules can execute code during.
9. `mint` - Controls token supply emissions, and what modules they are directed to.


### Attribution

We use the following modules from [Stride](https://github.com/Stride-Labs/stride) provided  under [this Apache V 2.0 License](https://github.com/Stride-Labs/stride/blob/main/LICENSE):

```
x/stakeibc
x/icacallbacks
x/records
x/claim
x/interchainquery
x/epochs
x/mint
```