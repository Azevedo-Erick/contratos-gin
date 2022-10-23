package middlewares

import (
	"net/http"
	"test/auth"

	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		referer := c.Request.Header.Get("Referer")
		valid, claims := auth.VerifyToken(tokenString, referer)
		if !valid {
			c.JSON(http.StatusUnauthorized, "Token inválido")
			c.Abort()
		}
		if len(c.Keys) == 0 {
			c.Keys = make(map[string]interface{})
		}
		c.Keys["Username"] = claims.Username
		c.Keys["Roles"] = claims.Roles
	}
}

func Authorization(validRoles []int) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Keys) == 0 {
			c.JSON(http.StatusUnauthorized, "Erro interno")
			c.Abort()
		}
		rolesVal := c.Keys["Roles"]
		if rolesVal == nil {
			c.JSON(http.StatusUnauthorized, "Erro ao recuperar cargos")
			c.Abort()
		}
		roles := rolesVal.([]int)
		validation := make(map[int]int)
		for _, value := range roles {
			validation[value] = 0
		}
		for _, value := range validRoles {
			if _, ok := validation[value]; !ok {
				c.JSON(http.StatusUnauthorized, "Você não tem permissão para acessar este recurso")
				c.Abort()
			}
		}
	}
}
