package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type Ingredient struct {
	Article        string   `gorm:"primaryKey;column:article"`
	Name           string   `gorm:"column:name"`
	Unit           string   `gorm:"column:unit"`
	Amount         int      `gorm:"column:amount"`
	SupplierName   string   `gorm:"column:supplier"`
	Supplier       Supplier `gorm:"column:supplier;foreignKey:SupplierName"`
	Image          string   `gorm:"column:image"`
	IngredientType string   `gorm:"column:ingredient_type"`
	CostPrice      float32  `gorm:"column:cost_price"`
	Gost           string   `gorm:"column:gost"`
	Packing        string   `gorm:"column:packing"`
	Specs          string   `gorm:"column:specs"`
}

type IngredientResponse struct {
	Article      string  `json:"article"`
	Name         string  `json:"name"`
	Amount       int     `json:"amount"`
	Unit         string  `json:"unit"`
	CostPrice    float32 `json:"cost_price"`
	SupplierName string  `json:"supplier"`
	DeliveryTime int     `json:"delivery_time"`
}

type IngredientRepository interface {
	Create(ingredient *Ingredient) error
	Fetch() ([]Ingredient, error)
	Edit(ingredient *Ingredient) error
	Delete(article string) error
}

type IngredientUsecase interface {
	GetAll() ([]Ingredient, error)
	Edit(ingredient *Ingredient) error
	Delete(article string) error
}
