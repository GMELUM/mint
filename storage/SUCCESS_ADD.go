package storage

import (
	"database/sql"

	"mint/config"
	"mint/utils"
	"mint/utils/mysql"
)

func SUCCESS_ADD(limit int) (*bool, *mysql.MySQLError) {
	return mysql.Query(mysql.Core, mysql.Params{
		Exec:    "SUCCESS_GET",
		Args:    []any{limit},
		Timeout: config.MySQLQueryDuration,
	}, func(rows *sql.Rows) (*bool, *mysql.MySQLError) {
		return utils.ToPointer(true), nil
	})
}
