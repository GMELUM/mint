package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// GetEnvString retrieves an environment variable's value or returns the default value if not set.
// No type validation is necessary for string values, as all environment variables are inherently strings.
func GetEnvString(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetEnvArrayString retrieves an environment variable's value or returns the default value if not set.
// No type validation is necessary for string values, as all environment variables are inherently strings.
func GetEnvArrayString(key string, split string, defaultValue []string) []string {
	if value, exists := os.LookupEnv(key); exists {
		return strings.Split(value, split)
	}
	return defaultValue
}

// GetEnvInt retrieves an environment variable's value as an integer or returns the default value if not set.
// Panics if the value exists but cannot be converted to an integer.
func GetEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			panic(fmt.Sprintf("Environment variable %s is not a valid integer: %v", key, err))
		}
		return intValue
	}
	return defaultValue
}

// GetEnvDuration retrieves an environment variable's value as a duration or returns the default value if not set.
// Panics if the value exists but cannot be converted to a valid duration.
func GetEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		durationValue, err := time.ParseDuration(value)
		if err != nil {
			panic(fmt.Sprintf("Environment variable %s is not a valid duration: %v", key, err))
		}
		return durationValue
	}
	return defaultValue
}

// GetEnvBool retrieves an environment variable's value as a boolean or returns the default value if not set.
// Panics if the value exists but cannot be converted to a boolean.
func GetEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			panic(fmt.Sprintf("Environment variable %s is not a valid boolean: %v", key, err))
		}
		return boolValue
	}
	return defaultValue
}
