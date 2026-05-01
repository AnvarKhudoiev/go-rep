package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string   `json:"title" gorm:"not null"`
	Description string   `json:"description"`
	Completed   bool     `json:"completed" gorm:"default:false"`
	UserID      uint     `json:"user_id"`
	CategoryID  uint     `json:"category_id"` // Внешний ключ на категорию
	Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
	Tags        []Tag    `json:"tags" gorm:"many2many:todo_tags;"`
}