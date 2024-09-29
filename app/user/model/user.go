package model

import "demo/etc/types"

type User struct {
	types.Model
	Nickname  string         `json:"nickname"`
	Username  string         `json:"username" gorm:"unique"`
	Password  string         `json:"-"`
	Addresses []*UserAddress `json:"addresses,omitempty"`
}
