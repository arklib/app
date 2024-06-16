package model

import "time"

type Model struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt" gorm:"index"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"index"`
}

func All() []any {
	return []any{
		User{},
		UserAddress{},
	}
}
