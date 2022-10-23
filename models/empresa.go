package models

import "time"

type Empresa struct {
	Id                 int `gorm:"primaryKey"`
	Nome               string
	Cnpj               string
	ResponsavelEmpresa ResponsavelEmpresa
	CreatedAt          time.Time `gorm:"column:created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at"`
	DeletedAt          time.Time `gorm:"column:deleted_at"`
}
