package user

import (
	"demo/app/base"
)

type Fn struct {
	*base.Base
}

func New(base *base.Base) *Fn {
	fn := &Fn{
		base,
	}
	return fn
}
