package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type Supplier struct {
	Name         string `gorm:"primaryKey;column:name"`
	Address      string `gorm:"column:address"`
	DeliveryTime int    `gorm:"column:delivery_time"` // в часах
}

type SupplierRepository interface {
	Create(user *User) error
	Fetch() ([]User, error)
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}
