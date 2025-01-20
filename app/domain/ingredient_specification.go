package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type IngredientSpecification struct {
	ItemName       string     `gorm:"primaryKey;column:item"`
	Item           Item       `gorm:"primaryKey;column:item;foreignKey:ItemName"`
	IngredientName string     `gorm:"primaryKey;column:ingredient"`
	Ingredient     Ingredient `gorm:"primaryKey;column:ingredient;foreignKey:IngredientName"`
	Amount         int        `gorm:"column:amount"`
}

type IngredientSpecificationResponse struct {
	Name           string `json:"name"`
	RequiredAmount int    `json:"required_amount"`
}

type IngredientSpecificationRepository interface {
	FetchByItem(item string) ([]IngredientSpecification, error)
}
