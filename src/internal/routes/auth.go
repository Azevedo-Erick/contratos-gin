package routes

import (
	"net/http"
	"test/src/internal/auth"
	"test/src/internal/dtos"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	auth.POST("/login", loginHandler)
}

func loginHandler(c *gin.Context) {
	var loginObj dtos.LoginRequest

	if err := c.ShouldBindJSON(&loginObj); err != nil {
		c.JSON(http.StatusBadRequest, "Requisição inválida")
	}

	claims := &auth.JwtClaims{}

	claims.Username = loginObj.Username
	claims.Roles = []int{1, 2, 3}
	claims.Audience = c.Request.Header.Get("Referer")

	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(30) * time.Minute)
	tokenString, err := auth.GenerateToken(claims, expirationTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Erro ao criar o token")
	}
	c.JSON(http.StatusOK, dtos.LoginResponse{
		AccessToken: tokenString,
	},
	)
}
