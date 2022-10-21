package routes

import (
	"test/src/internal/middlewares"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func Run() {
	gin.ForceConsoleColor()
	router.Use(gin.Logger(), gin.Recovery(), middlewares.BasicAuth())

	getRoutes()
	router.Run(":3000")
}
func getRoutes() {
	v1 := router.Group("/v1")
	SetupContratosRoutes(v1)
	SetupUsuariosRoutes(v1)
}
