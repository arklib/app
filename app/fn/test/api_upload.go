package test

import (
	"fmt"

	"github.com/arklib/ark"
	"github.com/arklib/ark/errx"
)

type ApiUploadIn struct{}
type ApiUploadOut struct {
	Name string `json:"name"`
}

func (fn *Fn) ApiUpload(at *ark.At, in *ApiUploadIn) (out *ApiUploadOut, err error) {
	file, err := at.Http().FormFile("file")
	errx.Assert(err)

	// save file
	uploadPath := fmt.Sprintf("public/storage/%s", file.Filename)
	err = at.Http().SaveUploadedFile(file, uploadPath)
	errx.Assert(err)

	out = &ApiUploadOut{file.Filename}
	return
}
