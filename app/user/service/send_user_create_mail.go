package service

import (
	"context"
	"fmt"

	"demo/app/user/model"
)

func (it *Service) SendUserCreateMail(ctx context.Context, user *model.User) error {
	fmt.Println("user.SendUserCreateMail", user)
	return fmt.Errorf("123")
}
