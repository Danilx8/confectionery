package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type Tooling struct {
	Marking    string      `gorm:"primaryKey;column:marking"`
	TypeName   string      `gorm:"column:type"`
	Type       ToolingType `gorm:"column:type;foreignKey:TypeName"`
	Properties string      `gorm:"column:properties"`
}

type ToolingRepository interface {
	Create(user *User) error
	Fetch() ([]User, error)
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}
