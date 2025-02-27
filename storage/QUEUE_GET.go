package storage

import (
	"database/sql"

	"mint/config"
	"mint/shared/models"
	"mint/utils/mysql"
)

func QUEUE_GET(limit int) (*[]models.Queue, *mysql.MySQLError) {
	return mysql.Query(mysql.Core, mysql.Params{
		Exec:    "QUEUE_GET",
		Args:    []interface{}{limit},
		Timeout: config.MySQLQueryDuration,
	},
		func(rows *sql.Rows) (*[]models.Queue, *mysql.MySQLError) {
			queues := []models.Queue{}
			for rows.Next() {
				queue := models.Queue{}
				err := rows.Scan(
					&queue.ID,
					&queue.Transaction,
					&queue.Wallet,
					&queue.Amount,
					&queue.Message,
					&queue.CreatedAt,
					&queue.UpdatedAt,
				)
				if err != nil {
					return nil, mysql.NewError(err)
				}
				queues = append(queues, queue)
			}
			return &queues, nil
		})
}
