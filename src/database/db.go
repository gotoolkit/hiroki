package database

import (
	"database/sql"
	"fmt"
	"ssq/src/config"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/doug-martin/goqu.v4"
)

var DB *goqu.Database
var conn *sql.DB

func Init() {
	var err error
	cfg := config.GetInstance()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	conn, err = sql.Open("mysql", dsn)
	if err != nil {
		panic("Cannot opening database: " + err.Error())
	}

	conn.SetMaxOpenConns(20)
	conn.SetMaxIdleConns(10)

	err = conn.Ping()
	if err != nil {
		panic("Cannot opening database connection: " + err.Error())
	}

	DB = goqu.New("mysql", conn)
}

func Close() {
	defer conn.Close()
}
