package lib

import (
	"github.com/arklib/ark"
	"github.com/arklib/ark/emitter"

	"app/src/dal/model"
)

type Email struct {
	*ark.Server
}

func NewEmail(srv *ark.Server) *Email {
	return &Email{srv}
}

func (s *Email) OnUserCreate(p emitter.Payload[model.User]) error {
	s.Logger.Noticef("[email] create user: %v", p.Data.Username)
	return p.Next()
}
