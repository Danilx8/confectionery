package route

import (
	"app/app/api/controller"
	"app/app/api/middleware"
	"app/app/domain"
	"app/app/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewOrdersHistoryRoute(db gorm.DB, group *gin.RouterGroup) {
	hr := repository.NewOrdersHistoryRepository(&db)
	hc := controller.NewOrdersHistoryController(hr)
	group.POST("/order/history", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.Director],
		domain.RoleName[domain.ClientManager],
	}), hc.Get)
}
