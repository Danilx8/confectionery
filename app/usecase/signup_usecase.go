package usecase

import (
	"app/app/domain"
	"app/app/internal/tokenutil"
	"time"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) Create(user *domain.User) (*domain.User, error) {
	return su.userRepository.Create(user)
}

func (su *signupUsecase) GetUserByLogin(login string) (*domain.User, error) {
	return su.userRepository.FetchByLogin(login)
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
