package controller

import (
	"app/app/bootstrap"
	domain2 "app/app/domain"
	"app/app/internal/passwordutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain2.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) Signup(c *gin.Context) { // Регистрация только для заказчиков (Client)
	var request domain2.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = sc.SignupUsecase.GetUserByLogin(request.Login)
	if err == nil {
		c.JSON(http.StatusConflict, domain2.ErrorResponse{Message: "User already exists with the given login"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	err = passwordutil.ValidateClientPassword(request.Password, request.Login)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := domain2.User{
		Login:    request.Login,
		Password: request.Password,
		Role:     "Client",
	}

	newUser, err := sc.SignupUsecase.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(newUser, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHours)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(newUser, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHours)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain2.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := domain2.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
