package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Success defines the structure for the request payload.
// This structure will be used to bind incoming JSON data.
type Success struct {
	ID          int       `json:"id" binding:"required"`
	Transaction string    `json:"transaction" binding:"required"`
	Hash        string    `json:"hash" binding:"required"`
	CreatedAt   time.Time `json:"created_at" binding:"required"`
	UpdatedAt   time.Time `json:"updated_at" binding:"required"`
}

// handlerReceiveSuccess processes incoming requests, prints received data, and responds with "OK".
func handlerReceiveSuccess(ctx *gin.Context) {
	var body Success

	// Bind the incoming JSON to Success and validate the input according to the struct tags
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Print the received data to the console
	fmt.Printf("Received Success: %+v\n", body)

	// Respond with a simple text message "OK"
	ctx.String(200, "OK")
}
