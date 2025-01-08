package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type CakeDecoration struct {
	Article      string   `gorm:"primaryKey;column:article"`
	Name         string   `gorm:"column:name"`
	Unit         string   `gorm:"column:unit"`
	Amount       int      `gorm:"column:amount"`
	SupplierName string   `gorm:"column:supplier"`
	Supplier     Supplier `gorm:"column:supplier;foreignKey:SupplierName"`
	Image        string   `gorm:"column:image"`
	Type         string   `gorm:"column:type"`
	CostPrice    float32  `gorm:"column:cost_price"`
	Weight       float32  `gorm:"column:weight"`
}

type CakeDecorationResponse struct {
	Article      string  `json:"article"`
	Name         string  `json:"name"`
	Amount       int     `json:"amount"`
	Unit         string  `json:"unit"`
	CostPrice    float32 `json:"cost_price"`
	SupplierName string  `json:"supplier"`
	DeliveryTime int     `json:"delivery_time"`
}

type CakeDecorationRepository interface {
	Create(decoration *CakeDecoration) error
	FetchAll() ([]CakeDecoration, error)
	Edit(decoration *CakeDecoration) error
	Delete(article string) error
}

type CakeDecorationUsecase interface {
	GetAll() ([]CakeDecoration, error)
	Edit(cakeDecoration *CakeDecoration) error
	Delete(article string) error
}
