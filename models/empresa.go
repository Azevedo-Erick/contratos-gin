package models

import "time"

type Empresa struct {
	Id                 int `gorm:"primaryKey"`
	RazaoSocial        string
	NomeFantasia       string
	InscricaoMunicipal string
	InscricaoEstadual  string
	PorteEmpresa       string
	RamoAtividade      string
	OptantePeloSimples bool
	Cnpj               string

	ResponsavelEmpresa ResponsavelEmpresa
	Endereco           Endereco
	Telefones          []Telefone
	Emails             []Email

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}
