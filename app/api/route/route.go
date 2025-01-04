package route

import (
	"app/app/api/middleware"
	"app/app/bootstrap"
	"gorm.io/gorm"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	toolingRouter := gin.Group("")
	// Middleware to verify AccessToken
	toolingRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	//toolingRouter.Use(middleware.RoleMiddleware([]string{domain.RoleName[domain.Director]}))
	NewToolingRouter(db, toolingRouter)
}
