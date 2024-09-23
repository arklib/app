package model

type UserAddress struct {
	Model
	UserId uint   `json:"userId"`
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}
