package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model `json:"-"` // This will hide ID, CreatedAt, UpdatedAt, DeletedAt
	ID         uint       `json:"id" gorm:"primarykey"` // Explicitly include ID
	Title      string     `json:"title"`
	AuthorID   uint       `json:"author_id"`
	Author     Author     `json:"-" gorm:"foreignKey:AuthorID"` // Hide from JSON
	CategoryID uint       `json:"category_id"`
	Category   Category   `json:"-" gorm:"foreignKey:CategoryID"` // Hide from JSON
}
