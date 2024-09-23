package dal

import (
	"log"

	"demo/app/base"
)

func Migrate(base *base.Base) {
	log.Println("db auto migrate")

	// err := app.DB.Debug().AutoMigrate(app.Models...)
	// if err != nil {
	// 	return
	// }
}
