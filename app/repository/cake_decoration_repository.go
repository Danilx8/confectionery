package repository

import (
	"app/app/domain"
	"gorm.io/gorm"
)

type cakeDecorationRepository struct {
	database *gorm.DB
}

func NewCakeDecorationRepository(database *gorm.DB) domain.CakeDecorationRepository {
	return &cakeDecorationRepository{
		database: database,
	}
}

func (cd cakeDecorationRepository) Create(cakeDecoration *domain.CakeDecoration) error {
	if result := cd.database.Create(cakeDecoration); result.Error != nil {
		return result.Error
	}
	return nil
}

func (cd cakeDecorationRepository) FetchAll() ([]domain.CakeDecoration, error) {
	var cakeDecorations []domain.CakeDecoration
	if result := cd.database.Table("cake_decorations").
		Joins("suppliers on cake_decorations.supplier = suppliers.id").
		Find(&cakeDecorations); result.Error != nil {
		return nil, result.Error
	}
	return cakeDecorations, nil
}

func (cd cakeDecorationRepository) FetchByID(id string) (*domain.CakeDecoration, error) {
	var cakeDecoration domain.CakeDecoration
	if result := cd.database.Where("article = ?", id).First(&cakeDecoration); result.Error != nil {
		return nil, result.Error
	}
	return &cakeDecoration, nil
}

func (cd cakeDecorationRepository) Edit(cakeDecoration *domain.CakeDecoration) error {
	if result := cd.database.Table("cake_decorations").Save(cakeDecoration); result.Error != nil {
		return result.Error
	}
	return nil
}

func (cd cakeDecorationRepository) Delete(article string) error {
	if result := cd.database.Table("cake_decorations").
		Delete(&domain.CakeDecoration{Article: article}); result.Error != nil {
		return result.Error
	}
	return nil
}
