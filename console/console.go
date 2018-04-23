package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"hiroki/console/command"
	"hiroki/database"
)

func main() {
	// 初始化数据库
	database.Init()

	app := cli.NewApp()
	app.Name = "console"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		cli.Command{
			Name:     "generate-full-reds-combination",
			Category: "db",
			Action:   command.GenerateFullRedsCombination,
		},
		cli.Command{
			Name:     "db-auto-migrate",
			Category: "db",
			Usage:    "更新数据表结构",
			Action:   command.DBAuthMigrate,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
