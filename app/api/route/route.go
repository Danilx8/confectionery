package route

import (
	"app/app/api/middleware"
	"app/app/bootstrap"
	"gorm.io/gorm"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db gorm.DB, gin *gin.Engine) {
	//gin.Static("/assets", env.StorageLocation)

	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	//protectedRouter.Use(middleware.RoleMiddleware([]string{domain.RoleName[domain.Director]}))
	NewToolingRouter(db, protectedRouter)
	NewIngredientRoute(db, protectedRouter)
	NewCakeDecorationRoute(db, protectedRouter)
}
