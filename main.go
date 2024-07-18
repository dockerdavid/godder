package main

import (
	"godder/internal/config"
	"godder/internal/database"
	"godder/pkg/sql"
	"time"
)

func init() {
	if err := config.LoadYmlConfig(); err != nil {
		panic(err)
	}
	database.ConnectDatabases()
}

func main() {
	idk()
}

func idk() {
	sql.CheckSlowQueries()
	time.Sleep(5 * time.Second)
	idk()
}
