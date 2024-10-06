package service

import (
	"context"

	"github.com/arklib/ark/hook"

	"demo/app/user/model"
)

func (it *Service) UserCreatePrint(ctx context.Context, user *model.User, next hook.Next) error {
	it.Logger.Infof("create user: %#v", user)
	return next()
}
