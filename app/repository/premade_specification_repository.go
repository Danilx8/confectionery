package repository

import (
	"app/app/domain"
	"gorm.io/gorm"
)

type premadeSpecificationRepository struct {
	database *gorm.DB
}

func NewPremadeSpecificationRepository(database *gorm.DB) domain.PremadeSpecificationRepository {
	return &premadeSpecificationRepository{
		database: database,
	}
}

func (c premadeSpecificationRepository) FetchByItem(item string) ([]domain.PremadeSpecification, error) {
	var premadeSpecifications []domain.PremadeSpecification
	result := c.database.Find(&premadeSpecifications, "item = ?", item)
	if result.Error != nil {
		return []domain.PremadeSpecification{}, result.Error
	}
	return premadeSpecifications, nil
}
