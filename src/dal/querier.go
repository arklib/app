package dal

import (
	"gorm.io/gen"
	"gorm.io/gorm"

	"app/src/dal/model"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE id = @id
	Get(id uint) (*gen.T, error)

	// SELECT * FROM @@table WHERE id IN @id
	GetMany(id ...uint) ([]*gen.T, error)
}

func UpdateQuerier(db *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		OutPath: "app/dal/query",
	})
	g.UseDB(db.Debug())

	// Generate models
	models := model.All()
	g.ApplyBasic(models...)
	g.ApplyInterface(func(Querier) {}, models...)

	g.Execute()
}
