package main

import (
	"fmt"
	"log"
	"mint/config"
	"mint/shared/middleware"
	"mint/utils/wallet"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var w *wallet.Wallet

func main() {

	// Load environment variables from a .env file if it exists, skip if not found
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found") // Log the absence as informational, not as an error
	}

	// Use defer and recover to handle panics gracefully.
	// Defer ensures the function is called at the end of main, recovering from any panic that might occur during runtime.
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r) // Log the panic information.
		}
	}()

	entity, err := wallet.New(config.WalletWords, "https://ton.org/global.config.json")
	if err != nil {
		panic(err) // Log any error that occurs during wallet initialization
	}

	w = entity

	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin engine instance with default middleware: logger and recovery.
	engine := gin.New()

	// Configure CORS (Cross-Origin Resource Sharing) to manage requests from different domains.
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // Allow requests from any origin. In production, it's better to specify allowed origins.
		AllowMethods:     []string{"GET", "POST"},                             // Allow only GET and POST requests to come through.
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Specify which headers are allowed in requests.
		ExposeHeaders:    []string{"Content-Length"},                          // Headers that can be exposed to the client.
		AllowCredentials: false,                                               // Disable credentials support for security.
		MaxAge:           12 * time.Hour,                                      // Set preflight request cache duration.
	}))

	// Define a POST route to handle withdrawal requests.
	engine.POST("withdraw", middleware.Secret, handlerWithdraw)

	// Attempt to run the server on the specified host and port.
	// fmt.Sprintf is used to create a formatted string for the address.
	if err := engine.Run(fmt.Sprintf("%v:%v", config.Host, config.Port)); err != nil {
		log.Fatalf("Failed to run server: %v", err) // Log any error that occurs while starting the server and exits the application.
	}
}
