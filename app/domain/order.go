package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Order struct {
	ID                     string     `gorm:"primaryKey;column:id"`
	Date                   time.Time  `gorm:"primaryKey;column:date"`
	Name                   string     `gorm:"column:name"`
	ItemName               string     `gorm:"column:item"`
	Item                   Item       `gorm:"column:item;foreignKey:ItemName;references:name"`
	OrdererName            string     `gorm:"column:orderer"`
	Orderer                User       `gorm:"column:orderer;foreignKey:OrdererName;references:login"`
	AssignedManagerName    string     `gorm:"column:assigned_manager"`
	AssignedManager        User       `gorm:"column:assigned_manager;foreignKey:AssignedManagerName;references:login"`
	Price                  float64    `gorm:"column:price"`
	ExpectedFulfilmentDate *time.Time `gorm:"column:expected_fulfilment_date"`
	Examples               string     `gorm:"column:examples"`
	Status                 string     `gorm:"column:status"`
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

type SpecificationRequest struct {
	ID                     string `json:"id"`
	Price                  string `json:"price"`
	ExpectedFulfilmentDate string `json:"expected_fulfilment_date"`
}

type AssuranceRequest struct {
	ID       string          `json:"id"`
	Criteria map[string]bool `json:"criteria"`
}

type OrderResponse struct {
	ID                     string  `json:"id"`
	Date                   string  `json:"date"`
	Name                   string  `json:"name"`
	Status                 string  `json:"status"`
	Price                  float64 `json:"price"`
	OrdererName            string  `json:"orderer"`
	ExpectedFulfilmentDate string  `json:"expected_fulfilment_date"`
	AssignedManagerName    string  `json:"assigned_manager"`
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
	FetchByUser(userLogin string) ([]Order, error)
	FetchByStatus(status string) ([]Order, error)
	CountToday() (int64, error)
	Update(order *Order) error
	Delete(id string) error
}

type OrderUsecase interface {
	Create(order *Order) error
	Fetch() ([]Order, error)
	GetByID(id string) (*Order, error)
	FetchOwn(login string) ([]Order, error)
	FetchByStatus(status string) ([]Order, error)
	Update(order *Order) error
	Delete(id string) error
	MapOrder(order *Order) OrderResponse
}
