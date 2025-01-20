package route

import (
	"app/app/api/controller"
	"app/app/api/middleware"
	"app/app/domain"
	"app/app/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewFailureRoute(db gorm.DB, group *gin.RouterGroup) {
	fr := repository.NewFailureRepository(&db)
	fc := controller.NewFailureController(fr)

	group.POST("/failure/register", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.Master],
	}), fc.RegisterFailure)
	group.GET("/failure/list", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.Master],
		domain.RoleName[domain.Director],
	}), fc.ListFailures)
}
