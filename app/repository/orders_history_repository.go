package repository

import (
	"app/app/domain"
	"gorm.io/gorm"
)

type ordersHistoryRepository struct {
	database *gorm.DB
}

func NewOrdersHistoryRepository(database *gorm.DB) domain.OrdersHistoryRepository {
	return &ordersHistoryRepository{
		database: database,
	}
}

func (o ordersHistoryRepository) Create(history *domain.OrdersHistory) error {
	if result := o.database.Table("orders_history").Create(history); result.Error != nil {
		return result.Error
	}
	return nil
}

func (o ordersHistoryRepository) FetchAll() ([]domain.OrdersHistory, error) {
	var history []domain.OrdersHistory
	if result := o.database.Table("orders_history").Find(&history); result.Error != nil {
		return nil, result.Error
	}
	return history, nil
}
