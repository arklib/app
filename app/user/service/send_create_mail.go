package service

import (
	"context"

	"demo/app/user/model"
)

func (it *Service) SendCreateMail(ctx context.Context, user *model.User) error {
	it.Logger.Infof("user.send_create_mail, user: %#v\n", user)
	return nil
}
