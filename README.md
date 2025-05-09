This readme is written by gpt-4o based on the code base.

# GoCoin

GoCoin is a simple blockchain implementation written in Go. It includes features such as a REST API, an HTML explorer, and a basic transaction system.

## Features

- **Blockchain**: Implements a basic blockchain with proof-of-work and adjustable difficulty.
- **Transactions**: Supports creating and validating transactions with unspent transaction outputs (UTxOs).
- **Mempool**: Manages unconfirmed transactions in memory.
- **Persistence**: Stores blockchain data using BoltDB.
- **REST API**: Provides endpoints for interacting with the blockchain.
- **HTML Explorer**: Offers a web-based interface to view and interact with the blockchain.

## Project Structure

```
.
├── blockchain/   # Core blockchain logic
├── cli/          # Command-line interface
├── db/           # Database interactions
├── explorer/     # HTML-based blockchain explorer
├── rest/         # REST API implementation
├── utils/        # Utility functions
└── main.go       # Entry point of the application
```

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/minjaelee0727/gocoin.git
    cd gocoin
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

### Usage

Run the application using the CLI:

```sh
go run main.go -port=4000 -mode=rest
```

#### Flags

- `-port`: Set the port for the server (default: 4000).
- `-mode`: Choose between `html` (HTML explorer) and `rest` (REST API).

### REST API Endpoints

- `GET /`: View API documentation.
- `GET /status`: View blockchain status.
- `GET /blocks`: View all blocks.
- `POST /blocks`: Add a new block.
- `GET /blocks/{hash}`: View a specific block.
- `GET /mempool`: View unconfirmed transactions.
- `GET /balance/{address}`: View balance for an address.

### HTML Explorer

Run the application in `html` mode to start the web-based blockchain explorer:

```sh
go run main.go -port=4000 -mode=html
```

Visit [http://localhost:4000](http://localhost:4000) in your browser.

## License

This project is licensed under the MIT License.
