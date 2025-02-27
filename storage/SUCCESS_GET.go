package storage

import (
	"database/sql"

	"mint/config"
	"mint/shared/models"
	"mint/utils/mysql"
)

func SUCCESS_GET(limit int) ([]*models.Success, *mysql.MySQLError) {
	successes, err := mysql.Query(mysql.Core, mysql.Params{
		Exec:    "SUCCESS_GET",
		Args:    []any{limit},
		Timeout: config.MySQLQueryDuration,
	}, func(rows *sql.Rows) (*[]*models.Success, *mysql.MySQLError) {
		var result []*models.Success
		for rows.Next() {
			success := &models.Success{}
			err := rows.Scan(
				&success.ID,
				&success.Transaction,
				&success.Hash,
				&success.CreatedAt,
				&success.UpdatedAt,
			)
			if err != nil {
				return nil, mysql.NewError(err)
			}
			result = append(result, success)
		}
		return &result, nil
	})
	if err != nil {
		return nil, err
	}
	return *successes, nil
}
