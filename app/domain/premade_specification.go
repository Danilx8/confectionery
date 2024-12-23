package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type PremadeSpecification struct {
	ItemName    int  `gorm:"primaryKey;column:item"`
	Item        Item `gorm:"primaryKey;column:item;foreignKey:ItemName"`
	PremadeName int  `gorm:"primaryKey;column:premade"`
	Premade     Item `gorm:"primaryKey;column:premade;foreignKey:PremadeName"`
	Amount      int  `gorm:"column:amount"`
}

type PremadeSpecificationRepository interface {
	Create(user *User) error
	Fetch() ([]User, error)
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}
