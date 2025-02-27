package tonlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type AddressBookEntry struct {
	UserFriendly string `json:"user_friendly"`
}

type Response struct {
	JettonWallets []JettonWallet              `json:"jetton_wallets"`
	AddressBook   map[string]AddressBookEntry `json:"address_book"`
}

var caches sync.Map
var cacheExpiration = time.Hour * 24 // Set cache expiration time

type CacheEntry struct {
	Data      *string
	ExpiresAt time.Time
}

func GetJettonWallet(ownerAddress, jettonAddress string) (*string, error) {
	// Generate a cache key based on ownerAddress and jettonAddress
	cacheKey := fmt.Sprintf("%s_%s", ownerAddress, jettonAddress)

	// Check if the wallet is in the cache
	if entry, found := caches.Load(cacheKey); found {
		cacheEntry := entry.(CacheEntry)
		if time.Now().Before(cacheEntry.ExpiresAt) {
			return cacheEntry.Data, nil // Return cached data if not expired
		}
	}

	baseURL := "https://toncenter.com/api/v3/jetton/wallets"
	params := url.Values{}
	params.Add("owner_address", ownerAddress)
	params.Add("jetton_address", jettonAddress)
	params.Add("limit", "1")
	params.Add("offset", "0")
	params.Add("api_key", "bfbf3e38b4ba9e9de21e6e425b942163a8d58f7ec93fa31b1cfb4f3e33821750")

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("неудачный запрос: статус код %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения тела ответа: %v", err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("ошибка десериализации JSON: %v", err)
	}

	if len(response.JettonWallets) == 0 {
		return nil, errors.New("WALLET_IS_NOT_EXISTS")
	}

	if item, ok := response.AddressBook[response.JettonWallets[0].Address]; ok {
		// Cache the result with expiration
		caches.Store(cacheKey, CacheEntry{
			Data:      &item.UserFriendly,
			ExpiresAt: time.Now().Add(cacheExpiration),
		})
		return &item.UserFriendly, nil
	}

	return nil, errors.New("WALLET_IS_NOT_EXISTS")
}
