package repository

import (
	"app/app/domain"
	"fmt"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		database: db,
	}
}

func (u userRepository) Create(user *domain.User) (*domain.User, error) {
	result := u.database.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, result.Error
}

func (u userRepository) FetchByLogin(login string) (*domain.User, error) {
	user := &domain.User{}

	result := u.database.Table("users").Where("login = ?", login).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch user with login %s: %w", login, result.Error)
	}

	return user, nil
}

func (u userRepository) Delete(login string) error {
	var user domain.User

	result := u.database.Where("login = ?", login).First(&user)
	if result.Error != nil {
		return fmt.Errorf("failed to found user with id %s: %w", login, result.Error)
	}
	result = u.database.Delete(&user)
	if result.Error != nil {
		return fmt.Errorf("failed to delete user with id %s: %w", login, result.Error)
	}

	return nil
}
