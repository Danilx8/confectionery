package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type OperationSpecification struct {
	ItemName        string      `gorm:"primaryKey;column:item"`
	Item            Item        `gorm:"primaryKey;column:item;foreignKey:ItemName"`
	Operation       string      `gorm:"primaryKey;column:operation"`
	SequenceNumber  int         `gorm:"primaryKey;column:sequence_number"`
	ToolingTypeName string      `gorm:"column:tooling_type"`
	ToolingType     ToolingType `gorm:"column:tooling_type;foreignKey:ToolingTypeName"`
	RequiredTime    int         `gorm:"column:required_time"` //в часах
}

type OperationSpecificationRepository interface {
	Create(user *User) error
	Fetch() ([]User, error)
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}
