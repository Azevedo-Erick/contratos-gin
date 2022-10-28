package repository

import (
	"test/database"
	"test/models"
)

func GetResponsavelEmpresaById(id int) (responsavelEmpresa models.ResponsavelEmpresa, err error) {
	err = database.Database.Db.Find(&responsavelEmpresa, "id=", id).Error
	return
}

func CreateResponsavelEmpresa(responsavelEmpresa models.ResponsavelEmpresa) (err error) {
	err = database.Database.Db.Create(&responsavelEmpresa).Error
	return
}

func UpdateResponsavelEmpresa(responsavelEmpresa models.ResponsavelEmpresa) (err error) {
	err = database.Database.Db.Save(&responsavelEmpresa).Error
	return
}

func DeleteResponsavelEmpresa(id int) (err error) {
	err = database.Database.Db.Delete(&models.ResponsavelEmpresa{}, "id=", id).Error
	return
}

func GetAllResponsavelEmpresa() (responsaveisEmpresa []models.ResponsavelEmpresa, err error) {
	err = database.Database.Db.Find(&responsaveisEmpresa).Error
	return
}
