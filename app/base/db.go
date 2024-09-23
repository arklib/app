package base

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"demo/app/query"
)

func (base *Base) initDB() {
	base.Logger.Info("[app] Init db")
	config := new(struct {
		DSN string
	})

	err := base.BindConfig("db", config)
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

	if base.IsDev() {
		db = db.Debug()
	}

	base.DB = db
	base.Query = query.Use(base.DB)
}
