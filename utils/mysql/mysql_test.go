package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectionString(t *testing.T) {
	// Test that the connection string is generated correctly.
	t.Run("Generate Connection String", func(t *testing.T) {
		options := Options{
			Host:     "localhost",
			Username: "user",
			Password: "pass",
			Database: "testdb",
			Port:     3306,
		}

		expected := "user:pass@tcp(localhost:3306)/testdb?parseTime=true"
		actual := connectionString(options)

		assert.Equal(t, expected, actual, "Connection string is not generated correctly")
	})
}
