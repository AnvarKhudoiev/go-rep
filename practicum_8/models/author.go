package models

import "gorm.io/gorm"

type Author struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `gorm:"unique" json:"email"`
	Todos []Todo `json:"todos"`
	gorm.Model `json:"-"`
}