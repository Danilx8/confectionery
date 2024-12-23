package route

import (
	"app/app/api/controller"
	"app/app/bootstrap"
	"app/app/repository"
	"app/app/usecase"
	"gorm.io/gorm"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(&db)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
