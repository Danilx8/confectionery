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

		_, err = tokenutil.ExtractIDFromToken(authToken, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		//c.Set("x-user-id", userID) тратит ну слишком много места
		c.Next()
		return
	}
}
