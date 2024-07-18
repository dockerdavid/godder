package main

import (
	"godder/internal/config"
	"godder/internal/database"
	"godder/pkg/disk"
	"godder/pkg/sql"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"
)

func init() {
	if err := config.LoadYmlConfig(); err != nil {
		panic(err)
	}
	database.ConnectDatabases()
}

func main() {

	c := cron.New()

	c.AddFunc("@every 1m", func() {
		disk.CheckDiskUsage()
		sql.CheckSlowQueries()
	})

	c.Start()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
