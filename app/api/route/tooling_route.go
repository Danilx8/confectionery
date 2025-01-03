package route

import (
	"app/app/api/controller"
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
	group.POST("/tooling/create", tc.Create)
	group.GET("/tooling/get", tc.Get)
	group.POST("/tooling/edit", tc.Edit)
	group.DELETE("/tooling/delete", tc.Delete)
}
