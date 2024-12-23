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

type CakeDecorationRepository interface {
	Create(user *User) error
	Fetch() ([]User, error)
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}
