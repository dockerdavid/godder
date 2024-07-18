package disk

import (
	"fmt"
	"godder/internal/config"
	"godder/shared"
	"syscall"
)

type DiskUsage struct {
	Free  string
	IsLow bool
}

func CheckDiskUsage() (DiskUsage, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/", &fs)
	if err != nil {
		return DiskUsage{}, err
	}

	var Free = fs.Bfree * uint64(fs.Bsize)

	return DiskUsage{
		Free:  fmt.Sprintf("%.2f GB", float64(Free)/float64(shared.GB)),
		IsLow: Free < uint64(config.Config.Godder.Disk.AlertThreshold)*shared.GB,
	}, nil
}
