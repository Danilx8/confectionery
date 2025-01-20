package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Login    string `gorm:"primaryKey;column:login"`
	Password string `gorm:"primaryKey;column:password"`
	Role     string `gorm:"column:role"`
	FullName string `gorm:"column:full_name"`
	PhotoURL string `gorm:"column:photo_url"`
}

type RoleEnum int

const (
	Client RoleEnum = iota
	ClientManager
	SupplementManager
	Master
	Director
)

var RoleName = map[RoleEnum]string{
	Client:            "Заказчик",
	ClientManager:     "Менеджер по работе с клиентами",
	SupplementManager: "Менеджер по закупкам",
	Master:            "Мастер",
	Director:          "Директор",
}

type UserRepository interface {
	Create(user *User) (*User, error)
	FetchByLogin(login string) (*User, error)
	Delete(login string) error
}
