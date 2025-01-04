package middleware

import (
	"app/app/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

func RoleMiddleware(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if setRole, exists := c.Get("x-user-role"); !exists {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
			c.Abort()
			return
		} else if !slices.Contains(roles, setRole.(string)) {
			c.JSON(http.StatusForbidden, domain.ErrorResponse{Message: fmt.Sprintf("Users with role %s are not allowed to access this endpoint", setRole.(string))})
			c.Abort()
			return
		} else {
			c.Next()
			return
		}

	}
}
