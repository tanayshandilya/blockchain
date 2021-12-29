# Bit Coverage
Bit Coverage coin **BCX** 

## Blockchain
A simple blockchain to support cryptocurrency transactions that map and verify the data exchange amongst nodes and gateways in an IoT ecosystem. It is comparable and similiar to the Network.

# Premetives

## Blocks
Blocks are used to record the most recent set of transactions. Blocks are mined based on time, as defined in the `block_time` chain variable. The current target block time is `60000` milliseconds (or 60 seconds). During any given `epoch`, the most recent block consists of:

- Block Version `int`
- Block Height `int`
- Previous Block Hash `string`
- Transactions (stored as a Merkle hash) `string`
- Threshold signature from the current consensus group `string`

## Epochs
An `epoch` is the target time period for which a given group of Miners is elected to serve as the consensus group. The target time for an epoch is currently `30 blocks`, as defined in the `election_interval` chain variable. Approximately every 30 blocks mined marks the passing of an epoch, after which a new group of Miners is elected to form the next consensus group. Mining rewards are distributed per epoch (as opposed to per block in most blockchain-based systems). At the conclusion of each epoch, the consensus group will distribute all the $BCX produced in that block via the `rewards` transaction.

## Transactions
All transactions occur on-chain, and all transactions require Data Credits to be submitted and confirmed. The following is a list of the supported transactions:

- `add gateway` - Add a new gateway to the Network. For the purposes of transactions, a “gateway” is the term for a compliant Miner that is mining and providing coverage.
- `assert location` - Assert a gateway’s location on the Network. This happens after a gateway has been added via the add gateway transaction. Once asserted, this location is then used as part of **Proof of Coverage** challenges. A Miner’s location can be asserted more than once but each subsequent assertion will a) cost a fee and b) reset that Miner’s score to neutral (.15)
- `chain vars` - Change a chain variable.
- `coinbase` - Similar to the bitcoin blockchain’s coinbase transaction but used only during testnet phases of the blockchain. The rewards transaction has taken its place.
coinbase data credits - Created the initial 10,000 Data Credits required to bring the first group of Miners online.
- `consensus group` - Marks the election of a new consensus group, responsible for mining during the next epoch.
create hashed timelock - Creates a transaction that can only be redeemed by providing the correct pre-image to the hashlock within the specified timelock.
- `create proof of coverage request` - Submitted by a Miner wishing to initiate a challenge.
- `data credits` - Burn BCX for DCs at the current oracle price and deliver them to the target wallet address.
- `genesis gateway` - Used to define the initial group of Miners that bootstrapped the blockchain.
- `multi-payment` - Used to send $BCX from one wallet to multiple wallets.
- `OUI` - Create a OUI for a new router on the network. In the blockchain, Miners forward packets to Routers that own them based on their OUI as stored in the blockchain.
- `payment` - Used to send $BCX from one wallet to another.
- `proof of coverage receipts` - The result of a POC submitted to the network upon completion.
- `redeem hashed timelock` - Redeem the transaction created using the create hashed timelock transaction.
- `reward` - A token payout for a specific event on the network such as submitting a valid proof of coverage request, participating in a consensus group, etc.
- `rewards` - Bundles multiple reward transactions at the end of each epoch and distributes all $BCX produced in that block to wallets that have earned them.
- `routing` - Update the routing information associated with an OUI.
- `security coinbase` - Distribution of security tokens in the genesis block.
- `security exchange` - The transfer of security tokens from one address to another.
- `state channel open` - Opens a new state channel on a Router
- `state channel close` - Closes a specific state channel on a Router
- `token burn exchange rate` - Change the exchange rate for burning $BCX to DCs.