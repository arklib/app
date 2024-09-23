package model

type User struct {
	Model
	Nickname  string         `json:"nickname"`
	Username  string         `json:"username" gorm:"unique"`
	Password  string         `json:"-"`
	Addresses []*UserAddress `json:"addresses,omitempty"`
}
