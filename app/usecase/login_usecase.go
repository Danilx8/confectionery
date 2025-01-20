package usecase

import (
	domain2 "app/app/domain"
	"app/app/internal/tokenutil"
	"time"
)

type loginUsecase struct {
	userRepository domain2.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain2.UserRepository, timeout time.Duration) domain2.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByLogin(login string) (*domain2.User, error) {
	return lu.userRepository.FetchByLogin(login)
}

func (lu *loginUsecase) CreateAccessToken(user *domain2.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain2.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
