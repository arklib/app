package app

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"demo/etc/query"
)

func (app *App) GetDB() *gorm.DB {
	if app.DB != nil {
		return app.DB
	}

	c := new(struct {
		DSN string
	})
	err := app.BindConfig("db", c)
	if err != nil {
		log.Fatalf("db c: %v", err)
	}

	dbInst, err := gorm.Open(mysql.New(mysql.Config{
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
		log.Fatalf("dbInst connect: %v", err)
	}

	if app.IsDev() {
		dbInst = dbInst.Debug()
	}
	app.Logger.Debug("[app] init db")

	app.DB = dbInst
	app.Query = query.Use(dbInst)
	return dbInst
}
