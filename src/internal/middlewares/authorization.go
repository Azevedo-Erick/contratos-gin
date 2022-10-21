package middlewares

import (
	"net/http"
	token "test/src/internal/auth"

	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		referer := c.Request.Header.Get("Referer")
		valid, claims := token.VerifyToken(tokenString, referer)
		if !valid {
			c.JSON(http.StatusUnauthorized, "Token inválido")
			c.Abort()
		}
		if len(c.Keys) == 0 {
			c.Keys = make(map[string]interface{})
		}
		c.Keys["claims"] = claims
	}
}
