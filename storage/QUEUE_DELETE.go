package storage

import (
	"database/sql"

	"mint/config"
	"mint/utils"
	"mint/utils/mysql"
)

func QUEUE_DELETE(transaction string) (*bool, *mysql.MySQLError) {
	return mysql.Query(mysql.Core, mysql.Params{
		Exec:    "QUEUE_DELETE",
		Args:    []any{transaction},
		Timeout: config.MySQLQueryDuration,
	}, func(rows *sql.Rows) (*bool, *mysql.MySQLError) {
		// Returning false since no rows are expected in the delete operation
		return utils.ToPointer(true), nil
	})
}
