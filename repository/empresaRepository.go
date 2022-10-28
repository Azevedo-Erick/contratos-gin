package repository

import (
	"test/database"
	"test/models"
)

func GetEmpresaById(id int) (models.Empresa, error) {
	var empresa models.Empresa
	err := database.Database.Db.Where("id = ?", id).First(&empresa).Error
	if err != nil {
		return empresa, err
	}
	return empresa, nil
}

func GetAllEmpresas() ([]models.Empresa, error) {
	var empresas []models.Empresa
	err := database.Database.Db.Find(&empresas).Error
	if err != nil {
		return empresas, err
	}
	return empresas, nil
}

func CreateEmpresa(empresa models.Empresa) (models.Empresa, error) {
	err := database.Database.Db.Create(&empresa).Error
	if err != nil {
		return empresa, err
	}
	return empresa, nil
}

func UpdateEmpresa(empresa models.Empresa) (models.Empresa, error) {
	err := database.Database.Db.Save(&empresa).Error
	if err != nil {
		return empresa, err
	}
	return empresa, nil
}

func DeleteEmpresa(empresa models.Empresa) error {
	err := database.Database.Db.Delete(&empresa).Error
	if err != nil {
		return err
	}
	return nil
}
