package config

import "mint/utils/env"

var (
	// Port defines the port number on which the server will listen.
	// This value is retrieved from the environment variable "PORT".
	// If the environment variable is not set, it defaults to 18300.
	Port = env.GetEnvInt("PORT", 18300)

	// Host specifies the hostname or IP address the server will bind to.
	// It is determined based on the environment variable "HOST".
	// If not set, it defaults to "0.0.0.0".
	Host = env.GetEnvString("HOST", "0.0.0.0")

	// Secret is used for security purposes, such as signing tokens or encrypting sensitive data.
	// This value is retrieved from the environment variable "SECRET".
	// If the environment variable is not set, it defaults to a predefined string.
	Secret = env.GetEnvString("SECRET", "")

	// CallbackURL is used to specify the URL to which callbacks should be sent.
	// This value is retrieved from the environment variable "CALLBACK_URL".
	// If the environment variable is not set, it defaults to an empty string.
	CallbackURL = env.GetEnvString("CALLBACK_URL", "")
)
