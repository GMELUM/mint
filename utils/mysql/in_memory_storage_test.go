package mysql

import (
	"fmt"
	"testing"
	"time"
)

// TestInMemoryStorage contains unit tests for the InMemoryStorage implementation.
func TestInMemoryStorage(t *testing.T) {
	storage := NewInMemoryStorage() // Initialize a new in-memory storage instance.

	// Test case for setting a value and then retrieving it.
	t.Run("SetAndGet", func(t *testing.T) {
		// Set a value with a 2-second expiration.
		err := storage.Set("key1", []byte("value1"), 2*time.Second)
		if err != nil {
			t.Fatalf("failed to set value: %v", err)
		}

		// Retrieve the value and ensure it matches the expected result.
		val, err := storage.Get("key1")
		if err != nil {
			t.Fatalf("failed to get value: %v", err)
		}

		if string(val) != "value1" {
			t.Errorf("expected value 'value1', got '%s'", val)
		}
	})

	// Test case to ensure keys expire after the specified duration.
	t.Run("KeyExpiration", func(t *testing.T) {
		// Set a value with a 1-second expiration.
		err := storage.Set("key2", []byte("value2"), 1*time.Second)
		if err != nil {
			t.Fatalf("failed to set value: %v", err)
		}

		// Wait 2 seconds to ensure the key has expired.
		time.Sleep(2 * time.Second)

		// Attempt to retrieve the expired key.
		_, err = storage.Get("key2")
		if err == nil {
			t.Fatalf("expected error for expired key, got none")
		}
	})

	// Test case for deleting a key and ensuring it is no longer accessible.
	t.Run("Delete", func(t *testing.T) {
		// Set a value with a long expiration.
		err := storage.Set("key3", []byte("value3"), 10*time.Second)
		if err != nil {
			t.Fatalf("failed to set value: %v", err)
		}

		// Delete the key.
		err = storage.Delete("key3")
		if err != nil {
			t.Fatalf("failed to delete key: %v", err)
		}

		// Ensure the key cannot be retrieved after deletion.
		_, err = storage.Get("key3")
		if err == nil {
			t.Fatalf("expected error for deleted key, got none")
		}
	})

	// Test case for resetting the entire storage.
	t.Run("Reset", func(t *testing.T) {
		// Set a value with a long expiration.
		err := storage.Set("key4", []byte("value4"), 10*time.Second)
		if err != nil {
			t.Fatalf("failed to set value: %v", err)
		}

		// Reset the storage.
		err = storage.Reset()
		if err != nil {
			t.Fatalf("failed to reset storage: %v", err)
		}

		// Ensure no keys can be retrieved after the reset.
		_, err = storage.Get("key4")
		if err == nil {
			t.Fatalf("expected error for reset storage, got none")
		}
	})

	// Test case for manual cleanup of expired keys.
	t.Run("CleanUp", func(t *testing.T) {
		// Set a value with a 1-second expiration.
		err := storage.Set("key5", []byte("value5"), 1*time.Second)
		if err != nil {
			t.Fatalf("failed to set value: %v", err)
		}

		// Wait 2 seconds for the key to expire.
		time.Sleep(2 * time.Second)

		// Manually trigger cleanup to remove expired keys.
		storage.cleanUp()

		// Ensure the expired key is no longer accessible.
		_, err = storage.Get("key5")
		if err == nil {
			t.Fatalf("expected error for cleaned-up key, got none")
		}
	})
}

// BenchmarkInMemoryStorage contains performance benchmarks for the InMemoryStorage implementation.
func BenchmarkInMemoryStorage(b *testing.B) {
	// Benchmark the Set operation.
	b.Run("Set", func(b *testing.B) {
		storage := NewInMemoryStorage()

		// Pre-create a list of keys to avoid repeated allocation during the benchmark.
		keys := make([]string, b.N)
		for i := 0; i < b.N; i++ {
			keys[i] = fmt.Sprintf("key%d", i)
		}

		b.ResetTimer() // Start measuring performance.

		// Perform the Set operation for all keys.
		for i := 0; i < b.N; i++ {
			_ = storage.Set(keys[i], []byte("value"), 10*time.Second)
		}
	})

	// Benchmark the Get operation.
	b.Run("Get", func(b *testing.B) {
		storage := NewInMemoryStorage()

		// Pre-create a list of keys and populate the cache.
		keys := make([]string, b.N)
		for i := 0; i < b.N; i++ {
			keys[i] = fmt.Sprintf("key%d", i)
			_ = storage.Set(keys[i], []byte("value"), 10*time.Second)
		}

		b.ResetTimer() // Start measuring performance.

		// Perform the Get operation for all keys.
		for i := 0; i < b.N; i++ {
			_, _ = storage.Get(keys[i])
		}
	})

	// Benchmark the Delete operation.
	b.Run("Delete", func(b *testing.B) {
		storage := NewInMemoryStorage()

		// Pre-create a list of keys and populate the cache.
		keys := make([]string, b.N)
		for i := 0; i < b.N; i++ {
			keys[i] = fmt.Sprintf("key%d", i)
			_ = storage.Set(keys[i], []byte("value"), 10*time.Second)
		}

		b.ResetTimer() // Start measuring performance.

		// Perform the Delete operation for all keys.
		for i := 0; i < b.N; i++ {
			_ = storage.Delete(keys[i])
		}
	})

	// Benchmark mixed operations (Set, Get, Delete).
	b.Run("MixedOperations", func(b *testing.B) {
		storage := NewInMemoryStorage()

		// Pre-create a list of keys.
		keys := make([]string, b.N)
		for i := 0; i < b.N; i++ {
			keys[i] = fmt.Sprintf("key%d", i)
		}

		b.ResetTimer() // Start measuring performance.

		// Perform a mix of Set, Get, and Delete operations.
		for i := 0; i < b.N; i++ {
			key := keys[i]
			_ = storage.Set(key, []byte("value"), 10*time.Second)
			_, _ = storage.Get(key)
			_ = storage.Delete(key)
		}
	})
}
