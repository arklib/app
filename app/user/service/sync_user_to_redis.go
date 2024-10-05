package service

import (
	"context"
	"fmt"

	"demo/app/user/model"
)

func (it *Service) SyncUserToRedis(ctx context.Context, user *model.User) error {
	fmt.Println("sync:SyncUserToRedis", user)
	return nil
}
