package tokenutil

import (
	"app/app/domain"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	claims := &domain.JwtCustomClaims{
		Name: user.Login,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "https://localhost:8080/",
			Subject:   "Authentication",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	claims := &domain.JwtCustomRefreshClaims{
		Name: user.Login,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Кондитерская",
			Subject:   "Authentication",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	return claims["id"].(string), nil
}
