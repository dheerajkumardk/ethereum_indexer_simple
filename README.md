## Ethereum Block Indexer 🚀

A Golang-based Ethereum Block Indexer that monitors the Ethereum blockchain, listens for newly created blocks, extracts relevant block data, and stores it in a SQLite database. It also provides REST API endpoints for retrieving stored blockchain data.

### 📌 Features
✅ Monitors Ethereum blockchain and listens to newly created blocks.

✅ Stores block data (Block Number, Hash, Miner, Number of Transactions) in a SQLite database using GORM ORM.

✅ Provides REST API to query stored block data.

### ⚙️ How It Works
1️⃣ Listening to Ethereum Blocks
The indexer connects to an Ethereum node via an RPC URL and continuously listens for new blocks.

🔹 Retrieves latest block number and details.

🔹 Extracts & logs Block Number, Block Hash, Txns, and Miner.

2️⃣ Storing Block Data in Database
The indexer saves block data into a SQLite database using GORM ORM.
Stored Fields:

BlockNumber
BlockHash
NumberOfTxns
Miner Address

3️⃣ Querying Data via API
A REST API is built using Fiber to allow users to retrieve indexed blocks.

### 🛠 Tech Stack
- Go (Golang) 🐹
- Ethereum go-ethereum (ethclient)
- SQLite (Database)
- GORM (ORM for database operations)
- Fiber (Fast Golang web framework)
- godotenv (For environment variables)

### 🚀 Setup and Installation

1️⃣ Clone the Repository
```sh
    git clone https://github.com/your-username/indexer-wallet-txns.git
    cd indexer-wallet-txns
```

2️⃣ Install Dependencies
```sh
    go mod tidy
```

3️⃣ Set Up .env File
Create a .env file and add your Ethereum RPC URL:

```ini
    ETH_RPC_URL=wss://mainnet.infura.io/ws/v3/YOUR_INFURA_PROJECT_ID
```

4️⃣ Run the Indexer
```sh
    go run main.go
```

💡 The indexer will start listening for new Ethereum blocks and store them in SQLite.

### 📡 API Endpoints

| **Method** | **Route**                            | **Description**                  |
|-----------|-------------------------------------|----------------------------------|
| `GET`     | `/block/blockNumber/:blockNumber`  | Get block by **Block Number**    |
| `GET`     | `/block/blockHash/:blockHash`      | Get block by **Block Hash**      |
| `GET`     | `/blocks`                          | Get **all blocks**               |
| `POST`    | `/blocks`                          | **Manually post a block** (JSON) |



#### Example API Call
- Get Block by Block Number
```sh
    curl -X GET http://localhost:3000/block/blockNumber/21960419
```

- Get Block by Block Hash
```sh
    curl -X GET http://localhost:3000/block/blockHash/0xdc9091ab54d22bd60185657e9eb6d943b27d66ce02cb75723880bc6c0d2a31cf
```

- Get All Blocks
```sh
    curl -X GET http://localhost:3000/blocks
```

