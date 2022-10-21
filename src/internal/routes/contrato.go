package routes

import "github.com/gin-gonic/gin"

func SetupContratosRoutes(r *gin.RouterGroup) {
	contratos := r.Group("/contratos")
	contratos.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	contratos.POST("/", func(ctx *gin.Context) {})
	contratos.PATCH("/", func(ctx *gin.Context) {})
	contratos.DELETE("/", func(ctx *gin.Context) {})

}
