package routes

import (
	"errors"
	"net/http"

	"test/database"
	"test/helpers"
	"test/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SetupUsuariosRoutes(r *gin.RouterGroup) {
	usuarios := r.Group("/usuarios")
	usuarios.GET("/", getUsuarios)
	usuarios.POST("/", saveUsuario)
	usuarios.PATCH("/:id", updateUsuario)
	usuarios.DELETE("/:id", deleteUsuario)
	usuarios.GET("/:id", getUsuario)
}

func getUsuarios(c *gin.Context) {
	usuarios := []models.Usuario{}
	err := database.Database.Db.Find(&usuarios).Error
	if err != nil {
		c.JSON(204, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, usuarios)
}

func findUsuario(id string, usuario *models.Usuario) error {
	error := database.Database.Db.Find(&usuario, "id=?", id).Error
	if error != nil {
		return errors.New("o usuario não existe")
	}
	return nil

}

func getUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := findUsuario(c.Param("id"), &usuario); err != nil {
		c.JSON(204, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, usuario)
}

func saveUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := c.BindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	validator := validator.New()

	if err := validator.Struct(usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	senha, err := helpers.HashPassword(usuario.Senha)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	usuario.Senha = senha

	if err := database.Database.Db.Create(&usuario).Error; err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, usuario)
}

func updateUsuario(c *gin.Context) {

	var usuario models.Usuario
	if err := c.BindJSON(&usuario); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	validator := validator.New()
	if err := validator.Struct(usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	var usu models.Usuario
	if err := findUsuario(c.Param("id"), &usu); err != nil {
		c.JSON(204, gin.H{"Error": err.Error()})
		return
	}
	usu.Nome = usuario.Nome
	usu.Email = usuario.Email
	usu.Senha = usuario.Senha

	err := database.Database.Db.Save(&usu).Error
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, usu)
}

func deleteUsuario(c *gin.Context) {
	var usuario models.Usuario
	id := c.Param("id")
	err := database.Database.Db.Where("id = ?", id).Delete(&usuario).Error
	if err != nil {
		c.JSON(204, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(204, gin.H{"message": "Usuario deletado com sucesso"})
}
