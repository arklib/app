package base

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"demo/app/model/query"
)

func (base *Base) GetDB() *gorm.DB {
	if base.DB != nil {
		return base.DB
	}

	config := new(struct {
		DSN string
	})
	err := base.BindConfig("db", config)
	if err != nil {
		log.Fatalf("db config: %v", err)
	}

	dbInst, err := gorm.Open(mysql.New(mysql.Config{
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
		log.Fatalf("dbInst connect: %v", err)
	}

	if base.IsDev() {
		dbInst = dbInst.Debug()
	}
	base.Logger.Debug("[app] init db")

	base.DB = dbInst
	base.Query = query.Use(dbInst)
	return dbInst
}
