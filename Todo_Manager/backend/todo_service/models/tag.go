package models
 
import (
	"time"
	"gorm.io/gorm"
)
 
type Tag struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Name      string         `json:"name" gorm:"unique;not null"`
	Todos     []Todo         `json:"-" gorm:"many2many:todo_tags"`
}
 