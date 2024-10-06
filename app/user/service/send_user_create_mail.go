package service

import (
	"context"

	"demo/app/user/model"
)

func (it *Service) SendUserCreateMail(ctx context.Context, user *model.User) error {
	it.Logger.Infof("user.send_user_create_mail, user: %#v\n", user)
	return nil
}
