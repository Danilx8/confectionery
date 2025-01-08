package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Order struct {
	ID                     string    `gorm:"primaryKey;column:id"`
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
	Status                 string    `gorm:"column:status"`
}

type OrderRequest struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Orderer     string `json:"orderer,omitempty"`
	Manager     string `json:"manager,omitempty"`
	Description string `json:"description"`
	Size        string `json:"size"`
	Examples    string `json:"examples"`
}

type StatusEnum int

const (
	New StatusEnum = iota
	Cancelled
	Specification
	Confirmation
	Supplement
	Production
	Assurance
	Ready
	Complete
)

var StatusName = map[StatusEnum]string{
	New:           "Новый",
	Cancelled:     "Отменен",
	Specification: "Составление спецификации",
	Confirmation:  "Подтверждение",
	Supplement:    "Закупка",
	Production:    "Производство",
	Assurance:     "Контроль",
	Ready:         "Готов",
	Complete:      "Выполнен",
}

type OrderRepository interface {
	Create(order *Order) error
	Fetch() ([]Order, error)
	FetchById(id string) (*Order, error)
	CountToday() (int64, error)
	Update(order *Order) error
	Delete(id string) error
}

type OrderUsecase interface {
	Create(order *Order) error
	Fetch() ([]Order, error)
	GetByID(id string) (*Order, error)
	Update(order *Order) error
	Delete(id string) error
}
