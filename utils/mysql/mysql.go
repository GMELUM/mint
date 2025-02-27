package mysql

import (
	"fmt"
	"sync"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	ErrMySQLNotInitialized = &MySQLError{
		Number:  45000,
		Message: "mysql is not initialized",
	}
)

// Options struct defines configuration parameters for the database connection.
type Options struct {
	Host           string  // The database host address (e.g., "localhost" or an IP address).
	Username       string  // The username to authenticate with the database.
	Password       string  // The password to authenticate with the database.
	Database       string  // The name of the specific database to connect to.
	Port           int     // The port number on which the database server is listening.
	MaxConnections int     // The maximum number of open database connections.
	Cache          Storage // A custom cache implementation (implements the Storage interface).
	CacheEnabled   bool    // A flag indicating whether caching is enabled.
	Mutex          Mutex   // A custom mutex implementation (implements the Mutex interface).
}

// SQL struct encapsulates the database connection, cache, and synchronization primitives.
type CoreEntity struct {
	DB           *sql.DB              // The underlying SQL database connection.
	prepare      map[string]*sql.Stmt // A map to store prepared SQL statements.
	cache        Storage              // The storage interface for caching query results.
	mutex        Mutex                // The mutex interface for synchronizing access.
	stop         chan bool            // A channel to signal the shutdown of the database connection.
	mx           sync.RWMutex         // A read-write mutex to synchronize internal access.
	CacheEnabled bool                 // Indicates whether caching is enabled.
}

// Storage interface defines methods for a generic key-value storage system.
type Storage interface {
	// Get retrieves the value associated with the given key.
	// Returns `nil, nil` if the key does not exist.
	Get(key string) ([]byte, error)

	// Set stores a key-value pair with an optional expiration duration.
	// A duration of 0 means no expiration.
	Set(key string, val []byte, exp time.Duration) error

	// Delete removes the value associated with the given key.
	Delete(key string) error

	// Reset clears all key-value pairs in the storage.
	Reset() error

	// Close cleans up resources used by the storage (e.g., stopping background workers).
	Close() error
}

// Mutex interface defines methods for locking and unlocking a resource by key.
type Mutex interface {
	// Lock attempts to acquire a lock for the given key.
	Lock(key string) error

	// Unlock releases the lock for the given key.
	Unlock(key string) error
}

// type MySQLError mysql.MySQLError

type MySQLError struct {
	Number   uint16
	SQLState [5]byte
	Message  string
}

func (me *MySQLError) Error() string {
	if me.SQLState != [5]byte{} {
		return fmt.Sprintf("Error %d (%s): %s", me.Number, me.SQLState, me.Message)
	}

	return fmt.Sprintf("Error %d: %s", me.Number, me.Message)
}

func (me *MySQLError) Is(err error) bool {
	if merr, ok := err.(*MySQLError); ok {
		return merr.Number == me.Number
	}
	return false
}

// // Global instance of the SQL struct. It acts as a singleton for the database connection.
var Core *CoreEntity

// New initializes the global MySQL instance with the given options.
func New(opt Options) (*CoreEntity, error) {
	// Open a connection to the MySQL database using the provided options.
	db, err := sql.Open("mysql", connectionString(opt))
	if err != nil {
		return nil, err // Terminate the program if the connection cannot be opened.
	}

	// Set connection pool limits based on the options.
	db.SetMaxOpenConns(opt.MaxConnections) // Maximum number of open connections.
	db.SetMaxIdleConns(opt.MaxConnections) // Maximum number of idle connections.

	// Set the maximum duration for which a connection can be reused.
	db.SetConnMaxLifetime(time.Minute * 5)

	// Ping the database to ensure the connection is valid.
	err = db.Ping()
	if err != nil {
		return nil, err // Terminate the program if the connection is invalid.
	}

	// Initialize the global MySQL instance.
	Core = &CoreEntity{
		DB:           db,
		prepare:      make(map[string]*sql.Stmt), // Initialize the map for prepared statements.
		CacheEnabled: opt.CacheEnabled,           // Enable caching based on the provided option.
	}

	// Set the custom mutex implementation, if provided.
	if opt.Mutex != nil {
		Core.mutex = opt.Mutex
	}

	// Set the custom cache implementation, if provided. Use in-memory storage as a fallback.
	if opt.Cache != nil {
		Core.cache = opt.Cache
	} else {
		Core.cache = NewInMemoryStorage()
	}

	return Core, nil

}

// Close cleans up resources used by the global MySQL instance.
func (c *CoreEntity) Close() {
	// Signal any background processes to stop.
	c.stop <- true

	// Close all prepared SQL statements.
	for _, stmt := range c.prepare {
		if stmt != nil {
			stmt.Close()
		}
	}

	// Close the database connection.
	c.DB.Close()
}

// connectionString constructs the MySQL connection string from the provided options.
func connectionString(opts Options) string {
	// Format: username:password@tcp(host:port)/database?parseTime=true
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", opts.Username, opts.Password, opts.Host, opts.Port, opts.Database)
}
