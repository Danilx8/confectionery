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

func NewToolingRouter(db gorm.DB, group *gin.RouterGroup) {
	tr := repository.NewToolingRepository(&db)
	tc := controller.ToolingController{
		ToolingUsecase: usecase.NewToolingUsecase(tr),
	}
	group.POST("/tooling/create", middleware.RoleMiddleware([]string{domain.RoleName[domain.Director]}), tc.Create)
	group.GET("/tooling/get", middleware.RoleMiddleware([]string{domain.RoleName[domain.Director]}), tc.Get)
	group.POST("/tooling/edit", middleware.RoleMiddleware([]string{domain.RoleName[domain.Director]}), tc.Edit)
	group.DELETE("/tooling/delete", middleware.RoleMiddleware([]string{domain.RoleName[domain.Director]}), tc.Delete)

}
