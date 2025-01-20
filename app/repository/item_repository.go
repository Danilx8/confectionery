package repository

import (
	"app/app/domain"
	"gorm.io/gorm"
)

type itemRepository struct {
	database *gorm.DB
}

func NewItemRepository(database *gorm.DB) domain.ItemRepository {
	return &itemRepository{
		database: database,
	}
}

func (i itemRepository) Fetch() ([]domain.Item, error) {
	var items []domain.Item
	result := i.database.Find(&items)
	if result.Error != nil {
		return []domain.Item{}, result.Error
	}

	return items, nil
}
