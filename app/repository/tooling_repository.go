package repository

import (
	"app/app/domain"
	"gorm.io/gorm"
)

type toolingRepository struct {
	database *gorm.DB
}

func NewToolingRepository(database *gorm.DB) domain.ToolingRepository {
	return &toolingRepository{
		database: database,
	}
}

func (t toolingRepository) Create(tooling *domain.Tooling) error {
	result := t.database.Create(tooling)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t toolingRepository) Fetch(conditions string) ([]domain.Tooling, error) {
	var toolings []domain.Tooling
	result := t.database.Table("toolings").Where(conditions).Find(&toolings)
	if result.Error != nil {
		return []domain.Tooling{}, result.Error
	}

	return toolings, nil
}

func (t toolingRepository) Edit(tooling *domain.Tooling) error {
	if result := t.database.Table("toolings").Save(tooling); result.Error != nil {
		return result.Error
	}
	return nil
}

func (t toolingRepository) Remove(marking string) error {
	result := t.database.Table("toolings").Where("marking = ?", marking).Delete(&domain.Tooling{Marking: marking})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
