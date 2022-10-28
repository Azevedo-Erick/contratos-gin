package models

import "time"

type Cidade struct {
	Id       int `gorm:"primaryKey"`
	Nome     string
	EstadoId int
	Estado   Estado `gorm:"foreignKey:EstadoId"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}
