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

type IngredientRepository interface {
	Create(user *User) error
	Fetch() ([]User, error)
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}
