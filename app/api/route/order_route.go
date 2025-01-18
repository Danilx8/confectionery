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

	group.GET("/order/all", middleware.RoleMiddleware([]string{domain.RoleName[domain.Director]}), oc.Get)
	group.GET("/order/own",
		oc.GetOwn)
	group.GET("/order/new", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.Director],
		domain.RoleName[domain.ClientManager],
	}), oc.GetNew)
	group.GET("/order/cancelled", middleware.RoleMiddleware([]string{domain.RoleName[domain.Director]}),
		oc.GetCancelled)
	group.GET("/order/specification", middleware.RoleMiddleware([]string{domain.RoleName[domain.Director]}),
		oc.GetSpecification)
	group.GET("/order/confirmation", middleware.RoleMiddleware([]string{domain.RoleName[domain.Director]}),
		oc.GetConfirmation)
	group.GET("/order/supplement", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.Director],
		domain.RoleName[domain.SupplementManager],
	}), oc.GetSupplement)
	group.GET("/order/production", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.Director],
		domain.RoleName[domain.Master],
	}), oc.GetProduction)
	group.GET("/order/assurance", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.Director],
		domain.RoleName[domain.Master],
	}))
	group.GET("/order/ready", middleware.RoleMiddleware([]string{domain.RoleName[domain.Director]}),
		oc.GetReady)
	group.GET("/order/complete", middleware.RoleMiddleware([]string{domain.RoleName[domain.Director]}),
		oc.GetComplete)

	group.POST("/order/accept", middleware.RoleMiddleware([]string{domain.RoleName[domain.ClientManager]}),
		oc.AcceptNewOrder)
	group.POST("/order/cancel", middleware.RoleMiddleware([]string{
		domain.RoleName[domain.Client],
		domain.RoleName[domain.ClientManager],
	}), oc.Cancel)
	group.POST("/order/specify", middleware.RoleMiddleware([]string{domain.RoleName[domain.ClientManager]}),
		oc.Specify)
	group.POST("/order/supplement", middleware.RoleMiddleware([]string{domain.RoleName[domain.ClientManager]}),
		oc.SetSupplement)
	group.POST("/order/production", middleware.RoleMiddleware([]string{domain.RoleName[domain.SupplementManager]}),
		oc.SetProduction)
	group.POST("/order/assurance", middleware.RoleMiddleware([]string{domain.RoleName[domain.Master]}),
		oc.SetAssurance)
	group.POST("/order/verdict", middleware.RoleMiddleware([]string{domain.RoleName[domain.Master]}),
		oc.AssureQuality)
	group.POST("/order/complete", middleware.RoleMiddleware([]string{domain.RoleName[domain.ClientManager]}),
		oc.SetComplete)

	group.DELETE("/order/delete", middleware.RoleMiddleware([]string{domain.RoleName[domain.Client]}),
		oc.Delete)
}
