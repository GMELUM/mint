package config

import "mint/utils/env"

var (
	// WalletWords holds an array of seed words used for wallet recovery or generation.
	// These words are retrieved from the environment variable "WALLET_WORDS",
	// split by spaces. If the environment variable is not set, it defaults to an empty slice.
	WalletWords = env.GetEnvArrayString("WALLET_WORDS", " ", []string{})

	// WalletDestination specifies the wallet address for transactions.
	// This value is retrieved from the environment variable "WALLET_DESTINATION".
	// If the environment variable is not set, it defaults to an empty string.
	WalletDestination = env.GetEnvString("WALLET_DESTINATION", "")

	// WalletJetton specifies the token or jetton identifier used in transactions.
	// This value is determined from the environment variable "WALLET_JETTON".
	// If the environment variable is not set, it defaults to an empty string.
	WalletJetton = env.GetEnvString("WALLET_JETTON", "")
)
