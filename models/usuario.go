package models

import "time"

type Usuario struct {
	Id        int `gorm:"primaryKey autoIncrement column:id"`
	Nome      string
	Email     string
	Senha     string
	Cargo     int
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}
