package database

import (
	"database/sql"
	"fmt"
	"godder/internal/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseStats struct {
	Name     string
	Host     string
	Port     int
	Database *sql.DB
}

var (
	Databases []DatabaseStats
)

func ConnectDatabases() {
	for _, database := range config.Config.Godder.SQL.Databases {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", database.User, database.Password, database.Host, database.Port)

		db, err := sql.Open("mysql", dsn)

		if err != nil {
			log.Fatalf("Failed to connect to database %s: %v", database.Name, err)
		}

		if err := db.Ping(); err != nil {
			db.Close()
			log.Fatalf("Failed to ping database %s: %v", database.Name, err)
		}

		log.Printf("Connected to database %s\n", database.Name)

		Databases = append(Databases, DatabaseStats{
			Name:     database.Name,
			Host:     database.Host,
			Port:     database.Port,
			Database: db,
		})
	}
}
