package models

import "time"

type Email struct {
	Id    int `gorm:"primaryKey"`
	Email string

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}
