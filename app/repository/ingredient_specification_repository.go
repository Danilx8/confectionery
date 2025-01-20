package repository

import (
	"app/app/domain"
	"gorm.io/gorm"
)

type ingredientSpecificationRepository struct {
	database *gorm.DB
}

func NewIngredientSpecificationRepository(database *gorm.DB) domain.IngredientSpecificationRepository {
	return &ingredientSpecificationRepository{
		database: database,
	}
}

func (c ingredientSpecificationRepository) FetchByItem(item string) ([]domain.IngredientSpecification, error) {
	var ingredientSpecifications []domain.IngredientSpecification
	result := c.database.Where("item = ?", item).Find(&ingredientSpecifications)
	if result.Error != nil {
		return []domain.IngredientSpecification{}, result.Error
	}
	return ingredientSpecifications, nil
}
