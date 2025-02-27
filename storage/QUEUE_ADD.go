package storage

import (
	"database/sql"

	"mint/config"
	"mint/utils"
	"mint/utils/mysql"
)

func QUEUE_ADD(transaction, wallet string, amount int64, message string) (*bool, *mysql.MySQLError) {
	return mysql.Query(mysql.Core, mysql.Params{
		Exec:    "QUEUE_ADD",
		Args:    []any{transaction, wallet, amount, message},
		Timeout: config.MySQLQueryDuration,
	}, func(rows *sql.Rows) (*bool, *mysql.MySQLError) {
		// Returning false since no rows are expected in the add operation
		return utils.ToPointer(true), nil
	})
}
