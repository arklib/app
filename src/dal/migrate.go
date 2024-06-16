package dal

import (
	"log"

	"gorm.io/gorm"

	"app/src/dal/model"
)

func Migrate(db *gorm.DB) {
	log.Println("db auto migrate")

	err := db.Debug().AutoMigrate(model.All()...)
	if err != nil {
		return
	}
}
