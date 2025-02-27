package storage

import (
	"database/sql"

	"mint/config"
	"mint/utils"
	"mint/utils/mysql"
)

func SUCCESS_DELETE(hash string) (*bool, *mysql.MySQLError) {
	return mysql.Query(mysql.Core, mysql.Params{
		Exec:    "SUCCESS_DELETE",
		Args:    []any{hash},
		Timeout: config.MySQLQueryDuration,
	}, func(rows *sql.Rows) (*bool, *mysql.MySQLError) {
		// Returning false since no rows are expected in the delete operation
		return utils.ToPointer(true), nil
	})
}
