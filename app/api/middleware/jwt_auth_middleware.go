package middleware

import (
	"app/app/domain"
	"app/app/internal/tokenutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) != 2 {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
			c.Abort()
			return
		}

		authToken := t[1]
		authorized, err := tokenutil.IsAuthorized(authToken, secret)
		if !authorized {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}

		userRole, err := tokenutil.ExtractRoleFromToken(authToken, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}

		r := strings.NewReplacer("\r", "")
		c.Set("x-user-role", r.Replace(userRole))
		c.Next()
		return
	}
}
