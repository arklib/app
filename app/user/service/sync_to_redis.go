package service

import (
	"context"

	"demo/app/user/model"
)

func (it *Service) SyncToRedis(ctx context.Context, user *model.User) error {
	it.Logger.Infof("user.sync_to_redis, user: %#v\n", user)
	return nil
}
