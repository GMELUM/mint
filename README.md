# TON Mint API

This project provides an API for managing transactions with TON wallets.

## Requirements

- [Go 1.23.2](https://golang.org/dl/)
- Configured environment variables (all mandatory):
  - `PORT`: The port number on which the server will listen.
  - `HOST`: The hostname or IP address to which the server will bind.
  - `WALLET_WORDS`: Seed words for wallet recovery or generation, separated by spaces.
  - `WALLET_DESTINATION`: The wallet address for processing transactions.
  - `WALLET_JETTON`: The identifier for the token or jetton used in transactions.

### Example `.env` File

Create a `.env` file in the root directory:

```env
PORT=18300
HOST=0.0.0.0
WALLET_WORDS="your seed words here"
WALLET_DESTINATION=your_wallet_destination_address
WALLET_JETTON=your_jetton_identifier
```

## Installation and Running

1. Clone the repository and navigate to the project directory:

   ```bash
   git clone https://github.com/your-username/your-repo.git
   cd your-repo
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

The server will be accessible at `http://<HOST>:<PORT>`.

## API Usage

### Withdraw Request

- **Route:** `POST /withdraw`
- **Request Body:**

```json
{
  "wallet": "recipient_wallet_address",
  "amount": 1000, // the amount to withdraw must be greater than zero
  "message": "Transaction message" // optional message for the transaction
}
```

### Successful Response

- **Format:**

```json
{
  "response": {
    "tx_hash": "LdSOGgjcvBuAPmCIEsL8Z48H8LvEiXXRFMxaeYSJeF4="
  }
}
```