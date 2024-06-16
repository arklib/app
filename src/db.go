package src

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func initDB(app *App) *gorm.DB {
	app.Logger.Info("[app] init db")
	config := new(struct {
		DSN string
	})

	err := app.BindConfig("db", config)
	if err != nil {
		log.Fatalf("db config: %v", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.DSN,
		DefaultStringSize:         64,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}

	if app.IsDev() {
		db = db.Debug()
	}
	return db
}
