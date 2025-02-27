package storage

import (
	"database/sql"

	"mint/config"
	"mint/utils"
	"mint/utils/mysql"
)

// QUEUE_SUCCESS executes a transaction where a record is added to the success table
// and removed from the queue table based on the transaction identifier.
func QUEUE_SUCCESS(transaction, hash string) (*bool, *mysql.MySQLError) {
	return mysql.Query(mysql.Core, mysql.Params{
		Exec:    "QUEUE_SUCCESS",
		Args:    []any{transaction, hash},
		Timeout: config.MySQLQueryDuration,
	}, func(rows *sql.Rows) (*bool, *mysql.MySQLError) {
		// Returning false since no rows are expected in the transaction operation
		return utils.ToPointer(true), nil
	})
}
