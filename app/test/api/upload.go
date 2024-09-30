package api

import (
	"fmt"

	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"
)

type (
	UploadIn  struct{}
	UploadOut struct {
		Name string `json:"name"`
	}
)

func (it *Api) Upload(ctx *ark.Ctx, in *UploadIn) (out *UploadOut, err error) {
	file, err := ctx.HttpReq().FormFile("file")
	errx.Assert(err)

	// save file
	uploadPath := fmt.Sprintf("public/storage/%s", file.Filename)
	err = ctx.HttpReq().SaveUploadedFile(file, uploadPath)
	errx.Assert(err)

	out = &UploadOut{file.Filename}
	return
}
