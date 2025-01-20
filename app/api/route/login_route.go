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

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(&db)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
