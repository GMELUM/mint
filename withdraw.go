package main

import (
	"mint/config"
	"mint/utils/msg"

	"github.com/gin-gonic/gin"
)

// WithdrawBody defines the structure for the request payload of a withdrawal operation.
// It includes fields for the recipient's wallet address, the amount to transfer, and an optional message.
type WithdrawBody struct {
	Wallet  string `json:"wallet" binding:"required"`      // The recipient wallet address
	Amount  uint64 `json:"amount" binding:"required,gt=0"` // The amount of tokens to withdraw; must be greater than zero
	Message string `json:"message" binding:"required"`     // An optional message or comment for the transaction
}

// WithdrawResponse defines the structure for the response payload after a successful withdrawal.
// It includes the transaction hash which acts as a proof of transaction.
type WithdrawResponse struct {
	TxHash string `json:"tx_hash"` // The hash of the transaction to verify the successful transfer
}

// handlerWithdraw returns a Gin handler function to process withdrawal requests.
func handlerWithdraw(ctx *gin.Context) {
	var body WithdrawBody

	// Bind the incoming JSON to WithdrawBody and validate the input according to the struct tags
	if err := ctx.ShouldBindJSON(&body); err != nil {
		msg.InvalidFields(ctx) // Respond with an error message if validation fails
		return
	}

	// Perform the withdrawal operation using the wallet service, passing the jetton, source, destination, amount, and message
	txHash, err := w.Withdraw(
		config.WalletJetton,      // Jetton wallet address
		body.Wallet,              // Destination wallet address (to which to send)
		config.WalletDestination, // Source wallet address (from which to withdraw)
		body.Amount,              // Amount of tokens to transfer
		body.Message,             // Optional message or comment for the transaction
	)

	// If there's an error during the withdrawal, respond with a bad request message
	if err != nil {
		msg.BadRequest(ctx, err.Error())
		return
	}

	// If the withdrawal is successful, send a response with the transaction hash
	msg.Send(ctx, WithdrawResponse{
		TxHash: txHash, // Transaction hash that confirms the transaction
	})

}
