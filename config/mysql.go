package config

import (
	"time"

	"mint/utils/env"
)

// MySQL configuration
var (
	// MySQLHost specifies the hostname or IP address of the MySQL database server.
	// Environment variable: MYSQL_HOST
	MySQLHost = env.GetEnvString("MYSQL_HOST", "localhost")

	// MySQLUsername defines the username for authenticating to the MySQL database.
	// Environment variable: MYSQL_USERNAME
	MySQLUsername = env.GetEnvString("MYSQL_USERNAME", "root")

	// MySQLPassword specifies the password associated with the MySQLUsername for authentication.
	// Environment variable: MYSQL_PASSWORD
	MySQLPassword = env.GetEnvString("MYSQL_PASSWORD", "")

	// MySQLDatabase sets the name of the MySQL database to connect to.
	// Environment variable: MYSQL_DATABASE
	MySQLDatabase = env.GetEnvString("MYSQL_DATABASE", "")

	// MySQLPort specifies the port on which the MySQL server is listening.
	// Environment variable: MYSQL_PORT
	MySQLPort = env.GetEnvInt("MYSQL_PORT", 3306)

	// MySQLMaxConnections defines the maximum number of simultaneous connections allowed to the MySQL database.
	// Environment variable: MYSQL_MAX_CONNECTIONS
	MySQLMaxConnections = env.GetEnvInt("MYSQL_MAX_CONNECTIONS", 10)

	// CacheEnabled indicates whether query caching should be enabled for the MySQL database.
	// Environment variable: CACHE_ENABLED
	MySQLCacheEnabled = env.GetEnvBool("MYSQL_CACHE_ENABLED", false)

	// MutexEnabled specifies whether to enable the Redis-based mutex to handle cached query data for the MySQL database.
	// Environment variable: MUTEX_ENABLED
	MySQLMutexEnabled = env.GetEnvBool("MYSQL_MUTEX_ENABLED", false)

	// QueryDuration
	// Environment variable: MYSQL_QUERY_DURATION
	MySQLQueryDuration = env.GetEnvDuration("MYSQL_QUERY_DURATION", time.Second)
)
