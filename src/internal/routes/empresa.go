package routes

import "github.com/gin-gonic/gin"

func SetupEmpresaRoutes(r *gin.RouterGroup) {
	empresa := r.Group("/empresa")
	empresa.GET("/", func(c *gin.Context) {})
	empresa.POST("/", func(c *gin.Context) {})
	empresa.PUT("/", func(c *gin.Context) {})
	empresa.DELETE("/", func(c *gin.Context) {})
}
