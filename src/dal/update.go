package dal

import (
	. "app/src"
)

func Update(app *App) {

	Migrate(app.DB)
	if app.IsDev() {
		UpdateQuerier(app.DB)
	}
}
