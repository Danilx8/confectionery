package usecase

import (
	domain2 "app/app/domain"
	"app/app/internal/tokenutil"
	"time"
)

type refreshTokenUsecase struct {
	userRepository domain2.UserRepository
	contextTimeout time.Duration
}

func NewRefreshTokenUsecase(userRepository domain2.UserRepository, timeout time.Duration) domain2.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (rtu *refreshTokenUsecase) GetUserByLogin(email string) (*domain2.User, error) {
	return rtu.userRepository.FetchByLogin(email)
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *domain2.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *domain2.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractLoginFromToken(requestToken string, secret string) (string, error) {
	return tokenutil.ExtractLoginFromToken(requestToken, secret)
}
