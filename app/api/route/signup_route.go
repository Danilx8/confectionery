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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(&db)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
