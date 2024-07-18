package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigFile struct {
	Godder Godder `yaml:"godder"`
}

type Godder struct {
	Disk  Disk  `yaml:"disk"`
	SQL   SQL   `yaml:"sql"`
	Email Email `yaml:"email"`
}

type Disk struct {
	DiskUnit       string `yaml:"disk_unit"`
	AlertThreshold int    `yaml:"alert_threshold"`
}

type SQL struct {
	QueryUnit     string      `yaml:"query_unit"`
	SlowQueryTime int         `yaml:"slow_query_time"`
	Databases     []Databases `yaml:"databases"`
}

type Databases struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Email struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	From     string `yaml:"from"`
	Password string `yaml:"password"`
	To       string `yaml:"to"`
}

var (
	Config *ConfigFile
)

const (
	FileName = "config.yml"
)

func LoadYmlConfig() error {
	config := &ConfigFile{}

	file, err := os.Open(FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return err
	}

	Config = config
	return nil
}
