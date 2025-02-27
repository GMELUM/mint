package env

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvString(t *testing.T) {
	// Set up environment variable
	os.Setenv("TEST_STRING", "value")
	defer os.Unsetenv("TEST_STRING")

	// Test retrieving a valid string
	assert.Equal(t, "value", GetEnvString("TEST_STRING", "default"))

	// Test retrieving a default string when the variable is not set
	assert.Equal(t, "default", GetEnvString("MISSING_STRING", "default"))
}

func TestGetEnvInt(t *testing.T) {
	// Set up environment variable
	os.Setenv("TEST_INT", "123")
	os.Setenv("INVALID_INT", "abc")
	defer func() {
		os.Unsetenv("TEST_INT")
		os.Unsetenv("INVALID_INT")
	}()

	// Test retrieving a valid integer
	assert.Equal(t, 123, GetEnvInt("TEST_INT", 0))

	// Test retrieving a default integer when the variable is not set
	assert.Equal(t, 10, GetEnvInt("MISSING_INT", 10))

	// Test invalid integer value
	assert.PanicsWithValue(t, "Environment variable INVALID_INT is not a valid integer: strconv.Atoi: parsing \"abc\": invalid syntax",
		func() { GetEnvInt("INVALID_INT", 0) })
}

func TestGetEnvDuration(t *testing.T) {
	// Set up environment variable
	os.Setenv("TEST_DURATION", "5s")
	os.Setenv("INVALID_DURATION", "5seconds")
	defer func() {
		os.Unsetenv("TEST_DURATION")
		os.Unsetenv("INVALID_DURATION")
	}()

	// Test retrieving a valid duration
	assert.Equal(t, 5*time.Second, GetEnvDuration("TEST_DURATION", 0))

	// Test retrieving a default duration when the variable is not set
	assert.Equal(t, 10*time.Second, GetEnvDuration("MISSING_DURATION", 10*time.Second))

	// Test invalid duration value
	assert.PanicsWithValue(t, "Environment variable INVALID_DURATION is not a valid duration: time: unknown unit \"seconds\" in duration \"5seconds\"",
		func() { GetEnvDuration("INVALID_DURATION", 0) })
}

func TestGetEnvBool(t *testing.T) {
	// Set up environment variable
	os.Setenv("TEST_BOOL_TRUE", "true")
	os.Setenv("TEST_BOOL_FALSE", "false")
	os.Setenv("INVALID_BOOL", "not_a_bool")
	defer func() {
		os.Unsetenv("TEST_BOOL_TRUE")
		os.Unsetenv("TEST_BOOL_FALSE")
		os.Unsetenv("INVALID_BOOL")
	}()

	// Test retrieving a valid boolean (true)
	assert.True(t, GetEnvBool("TEST_BOOL_TRUE", false))

	// Test retrieving a valid boolean (false)
	assert.False(t, GetEnvBool("TEST_BOOL_FALSE", true))

	// Test retrieving a default boolean when the variable is not set
	assert.True(t, GetEnvBool("MISSING_BOOL", true))

	// Test invalid boolean value
	assert.PanicsWithValue(t, "Environment variable INVALID_BOOL is not a valid boolean: strconv.ParseBool: parsing \"not_a_bool\": invalid syntax",
		func() { GetEnvBool("INVALID_BOOL", false) })
}
