package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type IngredientSpecification struct {
	ItemID int  `gorm:"primaryKey;column:item_id"`
	Item   Item `gorm:"primaryKey;column:item_id;foreignKey:ItemID"`
	Amount int  `gorm:"column:amount"`
}

type IngredientSpecificationRepository interface {
	Create(user *User) error
	Fetch() ([]User, error)
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}
