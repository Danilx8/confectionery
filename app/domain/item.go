package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type Item struct {
	Name string `gorm:"primaryKey;column:name"`
	Size string `gorm:"column:size"`
}

type ItemSpecificationsResponse struct {
	Ingredients []IngredientSpecificationResponse     `json:"ingredients"`
	Decorations []CakeDecorationSpecificationResponse `json:"decorations"`
	Premades    []PremadeSpecificationResponse        `json:"premades"`
	Steps       string                                `json:"steps"`
	Description string                                `json:"description"`
}

type ItemRepository interface {
	Fetch() ([]Item, error)
}

type ItemUseCase interface {
	FetchAll() ([]Item, error)
	FetchRequired(name string) (ItemSpecificationsResponse, error)
}
