package domain

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenUsecase interface {
	GetUserByLogin(login string) (*User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	ExtractLoginFromToken(requestToken string, secret string) (string, error)
	ExtractRoleFromToken(requestToken string, secret string) (string, error)
}
