package repository

import (
	"app/app/domain"
	"gorm.io/gorm"
)

type cakeDecorationSpecificationRepository struct {
	database *gorm.DB
}

func NewCakeDecorationSpecificationRepository(database *gorm.DB) domain.CakeDecorationSpecificationRepository {
	return &cakeDecorationSpecificationRepository{
		database: database,
	}
}

func (c cakeDecorationSpecificationRepository) FetchByItem(item string) ([]domain.CakeDecorationSpecification, error) {
	var cakeDecorationSpecifications []domain.CakeDecorationSpecification
	result := c.database.Where("item = ?", item).Find(&cakeDecorationSpecifications)
	if result.Error != nil {
		return []domain.CakeDecorationSpecification{}, result.Error
	}
	return cakeDecorationSpecifications, nil
}
