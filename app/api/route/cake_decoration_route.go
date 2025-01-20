package route

import (
	"app/app/api/controller"
	"app/app/api/middleware"
	"app/app/domain"
	"app/app/repository"
	"app/app/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewCakeDecorationRoute(db gorm.DB, group *gin.RouterGroup) {
	dr := repository.NewCakeDecorationRepository(&db)
	dc := controller.CakeDecorationController{
		CakeDecorationUsecase: usecase.NewCakeDeoctaionUsecase(dr),
	}
	group.GET("/decoration/get", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.Master],
		domain.RoleName[domain.Director],
		domain.RoleName[domain.ClientManager],
		domain.RoleName[domain.SupplementManager],
	}), dc.FindAll)
	group.POST("decoration/edit", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.Director],
		domain.RoleName[domain.SupplementManager],
	}), dc.Edit)
	group.DELETE("decoration/delete", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.Director],
		domain.RoleName[domain.SupplementManager],
	}), dc.Delete)
}
