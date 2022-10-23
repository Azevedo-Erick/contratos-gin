package models

import (
	"time"

	"cloud.google.com/go/civil"
)

type ResponsavelEmpresa struct {
	Id             int `json:"id" gorm:"primaryKey"`
	Nome           string
	Email          string
	DataNascimento civil.Date

	Cpf string
	Rg  string

	Telefones []Telefone
	Emails    []Email

	Endereco Endereco

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}
