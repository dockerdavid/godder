package sql

import (
	"database/sql"
	"fmt"
	"godder/internal/config"
	"godder/internal/database"
	"godder/internal/email"
	"log"
	"time"
)

type RowDataPacket struct {
	Id      sql.NullInt64
	Command sql.NullString
	Time    sql.NullInt64
	Info    sql.NullString
}

type SlowQuery struct {
	Id         int
	Time       int
	Info       string
	SendedMail bool
	UnixTime   int64
}

var (
	slowQueries map[int64]SlowQuery = make(map[int64]SlowQuery)
)

func CheckSlowQueries() {
	for _, connection := range database.Databases {
		setSlowQueries(connection)
	}
}

func setSlowQueries(databaseStats database.DatabaseStats) {
	rows, err := databaseStats.Database.Query("SELECT ID, COMMAND, TIME, INFO FROM INFORMATION_SCHEMA.PROCESSLIST")

	if err != nil {
		log.Fatal("Error executing query:", err)
	}

	defer rows.Close()

	for rows.Next() {
		var row RowDataPacket
		err := rows.Scan(&row.Id, &row.Command, &row.Time, &row.Info)

		if err != nil {
			continue
		}

		if row.Command.String == "Query" && row.Info.String != "" {
			if slowQuery, exists := slowQueries[row.Id.Int64]; exists {
				slowQuery.Time = int(row.Time.Int64)
				slowQueries[row.Id.Int64] = slowQuery
			} else {
				slowQueries[row.Id.Int64] = SlowQuery{
					Id:   int(row.Id.Int64),
					Time: int(row.Time.Int64),
					Info: row.Info.String,
				}
			}

		}
	}

	for id, slowQuery := range slowQueries {
		if slowQuery.Time > config.Config.Godder.SQL.SlowQueryTime && !slowQuery.SendedMail {
			email.SendMail(fmt.Sprintf("Slow query detected: %s, DB alias: %s", slowQuery.Info, databaseStats.Name))

			slowQuery.SendedMail = true
			slowQuery.UnixTime = time.Now().Unix()
			slowQueries[id] = slowQuery
		}

		if slowQuery.SendedMail && time.Now().Unix()-slowQuery.UnixTime > 1200 {
			delete(slowQueries, id)
		}
	}
}
