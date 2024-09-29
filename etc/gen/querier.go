package gen

import (
	"gorm.io/gen"

	"demo/app"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE id = @id
	Get(id uint) (*gen.T, error)

	// SELECT * FROM @@table WHERE id IN @id
	GetMany(id ...uint) ([]*gen.T, error)
}

func BuildQuerier(app *app.App, output string) {
	g := gen.NewGenerator(gen.Config{
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		OutPath: output,
	})
	g.UseDB(app.UseDB())
	models := app.GetModels()

	// Generate models
	g.ApplyBasic(models...)
	g.ApplyInterface(func(Querier) {}, models...)
	g.Execute()
}
