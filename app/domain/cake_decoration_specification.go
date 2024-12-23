package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type CakeDecorationSpecification struct {
	ItemName           string         `gorm:"primaryKey;column:item"`
	Item               Item           `gorm:"primaryKey;column:item;foreignKey:ItemName"`
	CakeDecorationName string         `gorm:"primaryKey;column:cake_decoration"`
	CakeDecoration     CakeDecoration `gorm:"primaryKey;column:cake_decoration;foreignKey:CakeDecorationName"`
	Amount             int            `gorm:"column:amount"`
}

type CakeDecorationSpecificationRepository interface {
	Create(user *User) error
	Fetch() ([]User, error)
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}
