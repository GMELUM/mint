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
  - `SECRET`: A security key required for all API requests, to be sent via the Authorization header or as a query parameter.

### Example `.env` File

Create a `.env` file in the root directory:

```env
PORT=18300
HOST=0.0.0.0
WALLET_WORDS="your seed words here"
WALLET_DESTINATION=your_wallet_destination_address
WALLET_JETTON=your_jetton_identifier
SECRET=your_secret_key
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

### Authentication

All API requests must include the `SECRET` for authentication purposes. This can be provided in one of the following ways:
- **Authorization Header:**

  Include the `SECRET` as a token in the HTTP Authorization header:
  ```http
  Authorization: your_secret_key
  ```

- **Query Parameter:**

  Include the `SECRET` as a query parameter in the request URL:
  ```http
  http://<HOST>:<PORT>/withdraw?secret=your_secret_key
  ```

### Withdraw Request

- **Route:** `POST /withdraw`
- **Request Body:**

  The body of the request should contain the following JSON fields:

  ```json
  {
    "wallet": "recipient_wallet_address",
    "amount": 1000, // the amount to withdraw must be greater than zero
    "message": "Transaction message" // optional message for the transaction
  }
  ```

### Successful Response

- **Format:**

  Upon successful execution of the withdrawal request, the response will be formatted as follows:

  ```json
  {
    "response": {
      "tx_hash": "LdSOGgjcvBuAPmCIEsL8Z48H8LvEiXXRFMxaeYSJeF4="
    }
  }
  ```

### Error Response

When an error occurs, the response will contain an `error` object with the following format:

- **Error Format:**

  ```json
  {
    "error": {
      "code": 123,             // Numeric error code
      "message": "Error message", // Description of the error
      "critical": true            // Indicates if the error is critical; omitted if false
    }
  }
  ```
