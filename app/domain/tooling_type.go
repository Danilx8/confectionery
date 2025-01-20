package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type ToolingType struct {
	Name string `gorm:"primaryKey;column:Article"`
}

type ToolingTypeRepository interface {
	Create(user *User) error
	Fetch() ([]User, error)
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}
