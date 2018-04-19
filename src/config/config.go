package config

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/go-ini/ini"
)

var instance *Config

type Config struct {
	AppMode   string
	MaxResult uint

	ServerHost string
	ServerPort uint16

	DBHost     string
	DBPort     uint16
	DBName     string
	DBUser     string
	DBPassword string
}

func (c *Config) Print() {
	fmt.Println("%T", c)
}

func GetInstance() *Config {
	if instance == nil {
		instance = create()
	}
	return instance
}

func create() *Config {
	config := &Config{}

	absPath, _ := filepath.Abs("config.ini")
	cfg, err := ini.Load(absPath)
	if err != nil {
		panic("Fail to read file: " + err.Error())
	}

	// Mode
	rootSection := cfg.Section("")
	config.AppMode = rootSection.Key("app_mode").In("production", []string{"production", "development"})
	config.MaxResult, _ = rootSection.Key("max_result").Uint()

	// DB
	mysqlSection := cfg.Section("mysql")
	config.DBHost = mysqlSection.Key("host").String()
	dbPort, err := strconv.ParseUint(mysqlSection.Key("port").String(), 10, 16)
	if err != nil {
		panic(err)
	}
	config.DBPort = uint16(dbPort)
	config.DBUser = mysqlSection.Key("user").String()
	config.DBPassword = mysqlSection.Key("password").String()
	config.DBName = mysqlSection.Key("db").String()

	// Server
	servSection := cfg.Section("server")
	config.ServerHost = servSection.Key("listen").String()
	sPort, err := strconv.ParseUint(servSection.Key("port").String(), 10, 16)
	if err != nil {
		panic(err)
	}
	config.ServerPort = uint16(sPort)

	return config
}
