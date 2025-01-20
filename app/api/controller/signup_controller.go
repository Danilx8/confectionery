package controller

import (
	"app/app/bootstrap"
	"app/app/domain"
	"app/app/internal/passwordutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

// Signup godoc
// @Summary	Signup of clients
// @Tags Authorisation
// @Accept json
// @Produce json
// @Param        data    body   domain.SignupRequest true  "scheme of login"
// @Success 200 {object} domain.SignupResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 409 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /signup [post]
func (sc *SignupController) Signup(c *gin.Context) {
	var request domain.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = sc.SignupUsecase.GetUserByLogin(request.Login)
	if err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given login"})
		return
	}

	err = passwordutil.ValidateClientPassword(request.Password, request.Login)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user := domain.User{
		Login:    request.Login,
		Password: request.Password,
		FullName: request.FullName,
		Role:     domain.RoleName[domain.Client],
	}

	newUser, err := sc.SignupUsecase.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(newUser, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHours)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(newUser, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHours)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
