package route

import (
	"app/app/api/controller"
	"app/app/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewItemRouter(db gorm.DB, group *gin.RouterGroup) {
	ir := repository.NewItemRepository(&db)
	ic := controller.ItemController{ItemRepository: ir}
	group.GET("item/all", ic.Get)
}
