package dal

import (
	"demo/app/base"
)

func Update(base *base.Base) {
	Migrate(base.DB)
	if base.IsDev() {
		UpdateQuerier(base.DB)
	}
}
