package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type Item struct {
	Name string `gorm:"primaryKey;column:name"`
	Size string `gorm:"column:size"`
}

type ItemRepository interface {
	Create(user *User) error
	Fetch() ([]User, error)
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}
