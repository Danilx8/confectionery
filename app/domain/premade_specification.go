package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type PremadeSpecification struct {
	ItemName    string `gorm:"primaryKey;column:item"`
	Item        Item   `gorm:"primaryKey;column:item;foreignKey:ItemName"`
	PremadeName string `gorm:"primaryKey;column:premade"`
	Premade     Item   `gorm:"primaryKey;column:premade;foreignKey:PremadeName"`
	Amount      int    `gorm:"column:amount"`
}

type PremadeSpecificationResponse struct {
	Name           string `json:"name"`
	RequiredAmount int    `json:"required_amount"`
}

type PremadeSpecificationRepository interface {
	FetchByItem(item string) ([]PremadeSpecification, error)
}
