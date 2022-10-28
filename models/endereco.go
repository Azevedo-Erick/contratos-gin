package models

import "time"

type Endereco struct {
	Id          int `gorm:"primaryKey"`
	Cep         string
	Cidade      Cidade
	CidadeID    int `gorm:"foreignKey:Cidade"`
	Bairro      string
	Rua         string
	Numero      string
	Complemento string

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}
