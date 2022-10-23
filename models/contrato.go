package models

import "time"

type Contrato struct {
	Id          int `gorm:"primaryKey"`
	Contratante Empresa
	Contratado  Empresa
	Valor       float64
	DataInicio  string
	DataFim     string
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at"`
}
