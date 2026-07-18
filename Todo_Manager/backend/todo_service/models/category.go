package models

import (
	"time"
	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Name      string         `json:"name" gorm:"not null"`
	Color     string         `json:"color"`
	UserID    uint           `json:"user_id"`
	IsDefault bool           `json:"is_default" gorm:"default:false"`
}