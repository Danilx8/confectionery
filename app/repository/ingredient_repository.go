package repository

import (
	"app/app/domain"
	"gorm.io/gorm"
)

type ingredientRepository struct {
	database *gorm.DB
}

func NewIngredientRepository(database *gorm.DB) domain.IngredientRepository {
	return &ingredientRepository{
		database: database,
	}
}

func (i ingredientRepository) Create(ingredient *domain.Ingredient) error {
	if result := i.database.Create(ingredient); result.Error != nil {
		return result.Error
	}
	return nil
}

func (i ingredientRepository) Fetch() ([]domain.Ingredient, error) {
	var ingredients []domain.Ingredient
	if result := i.database.Table("ingredients").
		Joins("suppliers on ingredients.supplier = suppliers.name").Find(&ingredients); result.Error != nil {
		return nil, result.Error
	}
	return ingredients, nil
}

func (i ingredientRepository) FetchById(id string) (*domain.Ingredient, error) {
	var ingredient domain.Ingredient
	result := i.database.Find(&ingredient, "article = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ingredient, nil
}

func (i ingredientRepository) Edit(ingredient *domain.Ingredient) error {
	if result := i.database.Table("ingredients").Save(ingredient); result.Error != nil {
		return result.Error
	}
	return nil
}

func (i ingredientRepository) Delete(article string) error {
	if result := i.database.Table("ingredients").
		Delete(&domain.Ingredient{Article: article}); result.Error != nil {
		return result.Error
	}
	return nil
}
