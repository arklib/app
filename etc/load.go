package etc

import (
	"demo/app"
	"demo/app/test"
	"demo/app/user"
)

func Load(app *app.App) *app.App {
	return app.Use(
		test.Define,
		user.Define,
	)
}
