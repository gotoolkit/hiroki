package database

import (
	"hiroki/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var DB *gorm.DB

func Init() {
	var err error
	cfg := config.GetInstance()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("Cannot opening database: " + err.Error())
	}

	DB.DB().SetMaxOpenConns(20)
	DB.DB().SetMaxIdleConns(10)
}

func Close() {
	defer DB.Close()
}
