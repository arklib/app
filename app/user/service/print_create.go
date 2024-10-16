package service

import (
	"context"

	"github.com/arklib/ark/hook"

	"demo/app/user/model"
)

func (it *Service) PrintCreate(ctx context.Context, user *model.User, next hook.Next) error {
	it.Logger.Infof("print user: %#v", user)
	return next()
}
