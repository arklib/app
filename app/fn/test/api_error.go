package test

import (
	"errors"

	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"
)

type (
	ApiErrorIn struct {
		Type string `json:"type"`
	}
	ApiErrorOut struct{}
)

func (fn *Fn) ApiError(at *ark.At, in *ApiErrorIn) (out *ApiErrorOut, err error) {
	stdErr := errors.New("std error")

	switch in.Type {
	case "error":
		err = errx.New("error")
	case "error.code":
		err = errx.New("error code", 403)

	case "assert":
		errx.Assert(stdErr)
	case "assert.error":
		errx.Assert(stdErr, "assert error")
	case "assert.code":
		errx.Assert(stdErr, "assert code", 403)
	case "assert.x":
		errx.AssertX(stdErr, "assert x", 400)

	case "throw":
		errx.Throw("throw error")
	case "throw.ref":
		errx.Throw(stdErr)
	case "throw.code":
		errx.Throw("throw code", 403)
	case "throw.x":
		errx.ThrowX("throw x", 500)
	default:
		panic("服务异常")
	}
	return
}
