package user

import (
	"github.com/arklib/ark/event"

	"demo/app/model"
)

func (fn *Fn) OnCreateSendEmail(p event.Payload[model.User]) error {
	fn.Logger.Noticef("[sms] create user: %v", p.Data.Id)
	return p.Next()
}
