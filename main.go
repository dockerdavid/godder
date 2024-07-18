package main

import (
	"flag"
	"godder/internal/config"
	"godder/internal/database"
	"godder/pkg/disk"
	"godder/pkg/sql"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"
)

const (
	filePath      = "/etc/systemd/system/godder.service"
	godderService = "godder.service"

	cmd1 = "systemctl"
	cmd2 = "enable"
	cmd3 = "godder.service"
)

func init() {
	if err := config.LoadYmlConfig(); err != nil {
		panic(err)
	}
	database.ConnectDatabases()
}

func main() {

	install := flag.Bool("install", false, "Install into systemd")
	start := flag.Bool("start", false, "Start service")

	flag.Parse()

	if *install {
		installService()
	} else if *start {
		startService()
	} else {
		log.Println("Usage: godder -install|-start")
	}

}

func installService() {
	godderServiceFile, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	godderServiceData := `[Unit]
Description=Godder Service
Wants=network.target
After=network.target

[Service]
Type=simple
Restart=always
RestartSec=5
WorkingDirectory=/home/$USER/godder
ExecStart=/home/$USER/godder/godder -start

[Install]
WantedBy=multi-user.target
`
	_, err = godderServiceFile.Write([]byte(godderServiceData))
	if err != nil {
		log.Fatal(err)
	}

	err = godderServiceFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chmod(filePath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(cmd1, cmd2, cmd3)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(stdout))
}

func startService() {
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
