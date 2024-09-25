package fn

import (
	"demo/app/base"
	"demo/app/fn/test"
	"demo/app/fn/user"
)

func Define(base *base.Base) {
	base.AddFn("test", test.New(base))
	base.AddFn("user", user.New(base))
}
