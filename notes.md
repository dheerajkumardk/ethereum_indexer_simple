Build an indexer that tracks all the blocks created on the Ethereum chain

What will the Indexer do?
1. Monitor the Ethereum blockchain and listen to newly created blocks

- Connect Golang app with a blockchain node using rpc url
- Try getting latest block number
- Run continuously and print block numbers
- It listens to all blocks
- BlockNumber, BlockHash, Number of Txns, Miner

2. Save the relevant transaction data in the database
- Block data to be stored
- BlockNumber, BlockHash, NumberOfTxns, Miner Address

- DB sqlite
- GORM ORM

Files
- Storage data type, functions


3. Build an API to allow users to retrieve the stored transactions

Routes
- Get block by BlockNumber   /block/blockNumber/:blockNumber
- Get block by BlockHash     /block/blockHash/:blockHash
- Get all blocks             /blocks
- Post a block {blockNumber} /blocks
