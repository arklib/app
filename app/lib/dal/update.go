package dal

import (
	"demo/app/base"
)

func Update(base *base.Base) {
	Migrate(base)
	if base.IsDev() {
		UpdateQuerier(base)
	}
}
