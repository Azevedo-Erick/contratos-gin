package models

import "time"

type Telefone struct {
	Id       int `gorm:"primaryKey"`
	Ddd      string
	Telefone string

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}
