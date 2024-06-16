package lib

import (
	"github.com/arklib/ark"
	"github.com/arklib/ark/emitter"

	"app/src/dal/model"
)

type SMS struct {
	*ark.Server
}

func NewSMS(srv *ark.Server) *SMS {
	return &SMS{srv}
}

func (s *SMS) OnUserCreate(p emitter.Payload[model.User]) error {
	s.Logger.Noticef("[sms] create user: %v", p.Data.Id)

	return p.Next()
}
