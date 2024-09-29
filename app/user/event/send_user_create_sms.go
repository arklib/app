package event

import (
	"github.com/arklib/ark/event"

	"demo/app/user/model"
)

func (it *Event) SendUserCreateSMS(p event.Payload[model.User]) error {
	it.Logger.Noticef("[sms] create user: %v", p.Data.Id)
	return p.Next()
}
