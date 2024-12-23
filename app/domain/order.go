package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Order struct {
	ID                     int       `gorm:"primaryKey;column:id"`
	Date                   time.Time `gorm:"primaryKey;column:date"`
	Name                   string    `gorm:"column:name"`
	ItemName               string    `gorm:"column:item"`
	Item                   Item      `gorm:"column:item;foreignKey:ItemName"`
	OrdererName            string    `gorm:"column:orderer"`
	Orderer                User      `gorm:"column:orderer;foreignKey:OrdererName"`
	AssignedManagerName    string    `gorm:"column:assigned_manager"`
	AssignedManager        User      `gorm:"column:assigned_manager;foreignKey:AssignedManagerName"`
	Price                  float32   `gorm:"column:price"`
	ExpectedFulfilmentDate time.Time `gorm:"column:expected_fulfilment_date"`
	Examples               string    `gorm:"column:examples"`
}

type OrderRepository interface {
	Create(user *User) error
	Fetch() ([]User, error)
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}
