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

func NewIngredientRoute(db gorm.DB, group *gin.RouterGroup) {
	ir := repository.NewIngredientRepository(&db)
	ic := controller.IngredientController{
		IngredientUsecase: usecase.NewIngredientUsecase(ir),
	}
	group.GET("/ingredient/get", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.Master],
		domain.RoleName[domain.Director],
		domain.RoleName[domain.ClientManager],
		domain.RoleName[domain.SupplementManager],
	}), ic.FindAll)
	group.POST("ingredient/edit", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.Director],
		domain.RoleName[domain.SupplementManager],
	}), ic.Edit)
	group.DELETE("ingredient/delete", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.Director],
		domain.RoleName[domain.SupplementManager],
	}), ic.Delete)
}
