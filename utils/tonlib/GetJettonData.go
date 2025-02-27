package tonlib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type JettonWalletResponse struct {
	Type         string       `json:"type"`
	Readme       string       `json:"readme"`
	JettonWallet JettonWallet `json:"jetton_wallet"`
	Jetton       Jetton       `json:"jetton"`
}

type JettonWallet struct {
	Address           string `json:"address"`
	Balance           string `json:"balance"`
	OwnerAddress      string `json:"owner_address"`
	JettonAddress     string `json:"jetton_address"`
	Owner             string `json:"owner"`
	Jetton            string `json:"jetton"`
	LastTransactionLT string `json:"last_transaction_lt"`
	CodeHash          string `json:"code_hash"`
	DataHash          string `json:"data_hash"`
}

type Jetton struct {
	TotalSupply        string   `json:"total_supply"`
	IsMutable          bool     `json:"is_mutable"`
	HasOnchainMetadata bool     `json:"has_onchain_metadata"`
	AdminAddress       *string  `json:"admin_address"`
	MetadataURL        *string  `json:"metadata_url"`
	Metadata           Metadata `json:"metadata"`
}

type Metadata struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Symbol      string                 `json:"symbol"`
	Image       map[string]interface{} `json:"image"`
	ImageData   *string                `json:"image_data"`
	Decimals    int                    `json:"decimals"`
	URI         string                 `json:"uri"`
}

var (
	cache     = make(map[string]JettonWalletResponse)
	cacheLock sync.Mutex
)

func GetJettonData(contractAddress string) (*JettonWalletResponse, error) {
	baseURL := "https://api.ton.cat/v2/contracts/jetton/"
	fullURL := baseURL + contractAddress

	cacheLock.Lock()
	if data, found := cache[fullURL]; found {
		cacheLock.Unlock()
		return &data, nil
	}
	cacheLock.Unlock()

	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var jettonData JettonWalletResponse
	err = json.Unmarshal(body, &jettonData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %v", err)
	}

	cacheLock.Lock()
	cache[fullURL] = jettonData
	cacheLock.Unlock()

	return &jettonData, nil
}
