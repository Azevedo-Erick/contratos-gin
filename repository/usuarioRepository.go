package repository

import (
	"test/database"
	"test/models"
)

func GetUsuarioByEmail(email string) (usuario models.Usuario, err error) {
	err = database.Database.Db.Where("email = ?", email).First(&usuario).Error
	return
}
