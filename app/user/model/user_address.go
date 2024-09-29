package model

import "demo/etc/types"

type UserAddress struct {
	types.Model
	UserId uint   `json:"userId"`
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}
