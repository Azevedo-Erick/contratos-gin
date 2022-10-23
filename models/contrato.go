package models

import (
	"time"

	"cloud.google.com/go/civil"
)

type Contrato struct {
	Id          int `gorm:"primaryKey"`
	Contratante Empresa
	Contratado  Empresa
	Valor       float64
	DataInicio  civil.Date
	DataFim     civil.Date

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}
