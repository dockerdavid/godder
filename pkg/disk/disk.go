package disk

import (
	"fmt"
	"godder/internal/config"
	"godder/internal/email"
	"godder/shared"
	"syscall"
)

type DiskUsage struct {
	Free  string
	IsLow bool
}

func CheckDiskUsage() {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/", &fs)
	if err != nil {
		return
	}

	var Free = fs.Bfree * uint64(fs.Bsize)

	if Free < uint64(config.Config.Godder.Disk.AlertThreshold)*shared.GB {
		email.SendMail(fmt.Sprintf("Disk space is low: %.2f GB", float64(Free)/float64(shared.GB)))
	}
}
