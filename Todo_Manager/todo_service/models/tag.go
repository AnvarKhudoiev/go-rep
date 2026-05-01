package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model `json:"-"`
	Name string `gorm:"unique;not null" json:"name"`
	Todos []Todo `gorm:"many2many:todo_tags" json:"todos"`
}

