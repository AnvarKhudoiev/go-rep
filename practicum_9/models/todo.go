package models

import "gorm.io/gorm"

type Todo struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"`
	Completed bool `json:"completed"`
	AuthorID uint   `json:"author_id"`    
	Author   Author `gorm:"foreignKey:AuthorID" json:"author"` 
	gorm.Model `json:"-"` 
}