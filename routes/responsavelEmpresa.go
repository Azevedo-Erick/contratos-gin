package routes

import (
	"test/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupResponsavelEmpresaRoutes(r *gin.RouterGroup) {
	responsavelEmpresa := r.Group("/responsavelEmpresa")
	responsavelEmpresa.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	responsavelEmpresa.POST("/", middlewares.ValidateToken(), middlewares.Authorization([]int{1}), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})
	responsavelEmpresa.PATCH("/", func(ctx *gin.Context) {})
	responsavelEmpresa.DELETE("/", func(ctx *gin.Context) {})
}
