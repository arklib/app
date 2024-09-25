package gen

import (
	"gorm.io/gen"
	"gorm.io/gorm"

	"demo/app/model"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE id = @id
	Get(id uint) (*gen.T, error)

	// SELECT * FROM @@table WHERE id IN @id
	GetMany(id ...uint) ([]*gen.T, error)
}

func Run(db *gorm.DB, output string) {
	models := model.GetModels()

	g := gen.NewGenerator(gen.Config{
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		OutPath: output,
	})
	g.UseDB(db)

	// Generate models
	g.ApplyBasic(models...)
	g.ApplyInterface(func(Querier) {}, models...)
	g.Execute()
}
