package models

import "time"

type Estado struct {
	Id   int `gorm:"primaryKey"`
	Nome string

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}
