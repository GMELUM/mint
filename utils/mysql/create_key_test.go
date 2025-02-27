package mysql

import (
	"testing"
	"time"
)

// TestConcat tests the Concat function.
func TestCreateKey(t *testing.T) {
	tests := []struct {
		query    string
		args     []interface{}
		expected string
	}{
		{
			"SELECT * FROM users WHERE id = ?",
			[]interface{}{42},
			"SELECT * FROM users WHERE id = ?42",
		},
		{
			"INSERT INTO users (name, created_at) VALUES (?, ?)",
			[]interface{}{"John Doe", time.Date(2024, 11, 17, 10, 0, 0, 0, time.UTC)},
			"INSERT INTO users (name, created_at) VALUES (?, ?)John Doe2024-11-17 10:00:00",
		},
		{
			"SELECT name FROM users WHERE age > ? AND country = ?",
			[]interface{}{30, "USA"},
			"SELECT name FROM users WHERE age > ? AND country = ?30USA",
		},
	}

	for _, test := range tests {
		t.Run(test.query, func(t *testing.T) {
			result := CreateKey(test.query, test.args...)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// BenchmarkConcat benchmarks the performance of the Concat function.
func BenchmarkCreateKey(b *testing.B) {
	// Test case for benchmarking
	query := "SELECT * FROM users WHERE id = ? AND name = ?"
	args := []interface{}{42, "John Doe"}

	// Run the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CreateKey(query, args...)
	}
}
