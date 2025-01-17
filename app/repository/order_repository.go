package repository

import (
	"app/app/domain"
	"gorm.io/gorm"
	"time"
)

type orderRepository struct {
	database *gorm.DB
}

func NewOrderRepository(database *gorm.DB) domain.OrderRepository {
	return &orderRepository{
		database: database,
	}
}

func (o orderRepository) Create(order *domain.Order) error {
	if result := o.database.Table("orders").Create(order); result.Error != nil {
		return result.Error
	}
	return nil
}

func (o orderRepository) Fetch() ([]domain.Order, error) {
	var orders []domain.Order
	if result := o.database.Table("orders").Find(&orders); result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (o orderRepository) FetchByUser(userLogin string) ([]domain.Order, error) {
	var orders []domain.Order
	if result := o.database.Table("orders").Where("orderer_name = ?", userLogin).Find(&orders); result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (o orderRepository) FetchById(id string) (*domain.Order, error) {
	var order *domain.Order
	if result := o.database.Table("orders").Where("id = ?", id).First(&order); result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (o orderRepository) CountToday() (int64, error) {
	var count int64
	if result := o.database.Table("orders").Where("date = ?", time.Now().Format("YYYY-MM-DD")).
		Count(&count); result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (o orderRepository) Update(order *domain.Order) error {
	if result := o.database.Table("orders").Save(order); result.Error != nil {
		return result.Error
	}
	return nil
}

func (o orderRepository) Delete(id string) error {
	if result := o.database.Table("orders").Where("id = ?", id).Delete(domain.Order{}); result.Error != nil {
		return result.Error
	}
	return nil
}
