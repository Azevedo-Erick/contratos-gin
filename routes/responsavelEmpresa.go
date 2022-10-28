package routes

import (
	"test/middlewares"
	"test/repository"

	"github.com/gin-gonic/gin"
)

func SetupResponsavelEmpresaRoutes(r *gin.RouterGroup) {
	responsavelEmpresa := r.Group("/responsavelEmpresa")
	responsavelEmpresa.GET("/", findAll)
	responsavelEmpresa.POST("/", middlewares.ValidateToken(), middlewares.Authorization([]int{1}), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})
	responsavelEmpresa.PATCH("/", func(ctx *gin.Context) {})
	responsavelEmpresa.DELETE("/", func(ctx *gin.Context) {})
}

func findAll(c *gin.Context) {
	data, err := repository.GetAllResponsavelEmpresa()
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, data)
}
