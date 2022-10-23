package models

import "time"

type ResponsavelEmpresa struct {
	Id    int `json:"id" gorm:"primaryKey"`
	Nome  string
	Email string

	Cpf       string
	Rg        string
	Telefone  string
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}
