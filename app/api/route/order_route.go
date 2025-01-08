package route

import (
	"app/app/api/controller"
	"app/app/api/middleware"
	"app/app/domain"
	"app/app/repository"
	"app/app/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func NewOrderRoute(timeout time.Duration, db gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(&db)
	or := repository.NewOrderRepository(&db)
	oc := controller.OrderController{
		UserUsecase:  usecase.NewLoginUsecase(ur, timeout),
		OrderUsecase: usecase.NewOrderUsecase(or),
	}
	group.POST("/order/create", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.Client],
		domain.RoleName[domain.ClientManager],
	}), oc.Create)
}
