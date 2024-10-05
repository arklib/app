package app

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"demo/etc/query"
)

func (app *App) initDB() {
	var c struct {
		DSN string
	}

	err := app.BindConfig("db", &c)
	if err != nil {
		log.Fatalf("db config: %v", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.DSN,
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
		// db = dbInst.Debug()
	}

	app.DB = db
	app.Query = query.Use(db)
}
