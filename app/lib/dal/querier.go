package dal

import (
	"gorm.io/gen"

	"demo/app/base"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE id = @id
	Get(id uint) (*gen.T, error)

	// SELECT * FROM @@table WHERE id IN @id
	GetMany(id ...uint) ([]*gen.T, error)
}

func UpdateQuerier(base *base.Base) {
	g := gen.NewGenerator(gen.Config{
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		OutPath: "dal/query",
	})
	g.UseDB(base.DB.Debug())

	// Generate models
	// g.ApplyBasic(base.Models...)
	// g.ApplyInterface(func(Querier) {}, base.Models...)

	g.Execute()
}
