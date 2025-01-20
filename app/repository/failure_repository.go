package repository

import (
	"app/app/domain"
	"gorm.io/gorm"
)

type failureRepository struct {
	database *gorm.DB
}

func NewFailureRepository(database *gorm.DB) domain.FailureRepository {
	return &failureRepository{
		database: database,
	}
}

func (f failureRepository) Create(failure *domain.Failure) error {
	if result := f.database.Create(failure); result.Error != nil {
		return result.Error
	}
	return nil
}

func (f failureRepository) FetchAll() ([]domain.Failure, error) {
	var failures []domain.Failure
	if result := f.database.Find(&failures); result.Error != nil {
		return nil, result.Error
	}
	return failures, nil
}
