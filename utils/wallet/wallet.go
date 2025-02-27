package wallet

import (
	"context"
	"encoding/base64"
	"mint/utils/tonlib"

	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

// Wallet represents a TON wallet, including context and current block information.
type Wallet struct {
	*wallet.Wallet                 // Embedded wallet struct from tonutils-go
	Context        context.Context // Execution context for network operations
	Block          *ton.BlockIDExt // Current block information to validate transactions
	Api            ton.APIClientWrapped
}

// New initializes and returns a new Wallet object using the provided seed words and network configuration URL.
// It returns a pointer to a Wallet instance or an error if initialization fails.
func New(words []string, network string) (*Wallet, error) {

	client := liteclient.NewConnectionPool()

	// Retrieve configuration from the URL
	cfg, err := liteclient.GetConfigFromUrl(context.Background(), network)
	if err != nil {
		return nil, err
	}

	// Connect to mainnet lite servers
	err = client.AddConnectionsFromConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	// Initialize API client with proof checking and retry capabilities
	api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()
	api.SetTrustedBlockFromConfig(cfg)

	// Bind all requests to a single TON node
	ctx := client.StickyContext(context.Background())

	// Create a new wallet instance using seed words and specific configuration
	w, err := wallet.FromSeed(api, words, wallet.ConfigV5R1Final{
		NetworkGlobalID: wallet.MainnetGlobalID,
	})
	if err != nil {
		return nil, err
	}

	// Get the current masterchain block information
	block, err := api.CurrentMasterchainInfo(context.Background())
	if err != nil {
		return nil, err
	}

	// Return the configured Wallet instance
	return &Wallet{
		w,
		ctx,
		block,
		api,
	}, nil
}

// Balance retrieves and returns the current balance of the wallet in NanoTON.
// It returns the balance as a uint64 or an error if retrieval fails.
func (w *Wallet) Balance() (uint64, error) {

	balance, err := w.GetBalance(w.Context, w.Block)
	if err != nil {
		return 0, err
	}

	return balance.Nano().Uint64(), nil
}

// Withdraw creates and executes a transaction to transfer Jettons from one address to another.
// Requires jetton details, from and to addresses, amount, and a message.
// Returns the transaction hash as a base64 encoded string or an error if the transaction fails.
func (w *Wallet) Withdraw(
	jetton string,
	fromAddress string,
	toAddress string,
	amount uint64,
	message string,
) (string, error) {

	// Create a transaction message with specific transfer options
	msg, err := tonlib.CreateTransaction(tonlib.JettonTransferOption{
		Jetton:              jetton,      // Identifier for the jetton
		Destination:         fromAddress, // Target wallet for the transfer
		ResponseDestination: toAddress,   // Source wallet for the response
		Message:             message,     // Optional message for the transaction
		Amount:              amount,      // Amount to transfer in NanoTON
	})

	if err != nil {
		return "", err
	}

	// Send the transaction and wait for confirmation
	tx, _, err := w.SendWaitTransaction(context.Background(), &wallet.Message{
		Mode:            3,   // Specifies transaction behavior
		InternalMessage: msg, // Encapsulated internal message for transaction
	})

	if err != nil {
		return "", err
	}

	// Return the transaction hash as a base64 encoded string
	return base64.StdEncoding.EncodeToString(tx.Hash), nil
}
