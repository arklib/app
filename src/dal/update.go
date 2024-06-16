package dal

import "app/src"

func Update(app *src.App) {

	Migrate(app.DB)
	if app.IsDev() {
		UpdateQuerier(app.DB)
	}
}
