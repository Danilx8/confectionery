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

func NewItemRouter(db gorm.DB, group *gin.RouterGroup) {
	ir := repository.NewItemRepository(&db)
	icr := repository.NewIngredientSpecificationRepository(&db)
	cdr := repository.NewCakeDecorationSpecificationRepository(&db)
	psr := repository.NewPremadeSpecificationRepository(&db)
	ic := controller.ItemController{
		ItemUseCase: usecase.NewItemUsecase(ir, icr, cdr, psr),
	}
	group.GET("item/all", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.SupplementManager],
		domain.RoleName[domain.Director],
		domain.RoleName[domain.Master],
		domain.RoleName[domain.ClientManager],
		domain.RoleName[domain.Client],
	}), ic.Get)
	group.POST("item/specifications", middleware.RoleMiddleware([]string{domain.RoleName[domain.Master]}),
		ic.Specifications)
}
