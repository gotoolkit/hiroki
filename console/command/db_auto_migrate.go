package command

import (
	"github.com/urfave/cli"
	"hiroki/database"
	"hiroki/model"
)

func DBAuthMigrate(c *cli.Context) error {
	db := database.DB

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&model.Number{})

	return nil
}
