package routes

import (
	"net/http"
	"test/auth"
	"test/dtos"
	"test/helpers"
	"test/repository"
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

	usuario, err := repository.GetUsuarioByEmail(loginObj.Email)

	if err != nil || !helpers.CheckPasswordHash(loginObj.Senha, usuario.Senha) {

		c.JSON(http.StatusUnauthorized, "Email ou senha inválidos")
		return
	}

	claims := &auth.JwtClaims{}

	claims.Username = usuario.Email
	claims.Roles = []int{usuario.Cargo}
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
