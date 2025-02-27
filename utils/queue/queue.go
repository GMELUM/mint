package queue

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"mint/config"
	"mint/shared/models"
	"mint/storage"
	"mint/utils/wallet"
	"net/http"
	"time"
)

func Sheldule() {

	defer func() {

		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r) // Log the panic information.
		}

		time.Sleep(time.Second)
		Sheldule()

	}()

	transaction, errSQL := storage.QUEUE_GET(3)
	if errSQL != nil {
		panic(errSQL)
	}

	messages := []wallet.Transaction{}
	for _, i := range *transaction {
		messages = append(messages, wallet.Transaction{
			Wallet:  i.Wallet,
			Amount:  uint64(i.Amount),
			Message: i.Message,
		})
	}

	if len(messages) == 0 {
		return
	}

	txHash, err := wallet.Core.Withdraw(
		config.WalletJetton,      // Jetton wallet address
		config.WalletDestination, // Source wallet address (from which to withdraw)
		messages,
	)

	if err != nil {
		panic(err)
	}

	for _, i := range *transaction {
		_, errSQL = storage.QUEUE_SUCCESS(i.Transaction, txHash)
		if errSQL != nil {
			panic(errSQL)
		}
	}

}

func Callback() {

	defer func() {

		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r) // Log the panic information.
		}

		time.Sleep(time.Second)
		Sheldule()

	}()

	transaction, errSQL := storage.SUCCESS_GET(10)
	if errSQL != nil {
		panic(errSQL)
	}

	for _, item := range transaction {
		res, err := PostSuccessAndCheckOK(config.CallbackURL, item, time.Second*5)
		if err != nil {
			continue
		}

		if res {
			storage.SUCCESS_DELETE(item.Hash)
		}
	}

}

func PostSuccessAndCheckOK(url string, success *models.Success, timeout time.Duration) (bool, error) {
	// Serialize struct to JSON
	body, err := json.Marshal(success)
	if err != nil {
		return false, err
	}

	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return false, err // Error occurred while sending the request
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err // Error occurred while reading the response
	}

	// Check for response code 200 and body containing "OK"
	if resp.StatusCode == http.StatusOK && string(responseBody) == "OK" {
		return true, nil
	}

	return false, nil // Return false if not 200 or body is not "OK"
}
