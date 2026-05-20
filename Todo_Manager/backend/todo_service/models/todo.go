package models
 
import (
	"time"
	"gorm.io/gorm"
)
 
type Priority string
 
const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)
 
type Todo struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description"`
	Completed   bool           `json:"completed" gorm:"default:false"`
	UserID      uint           `json:"user_id"`
	CategoryID  *uint          `json:"category_id"`
	Category    Category       `json:"category" gorm:"foreignKey:CategoryID"`
	Tags        []Tag          `json:"tags" gorm:"many2many:todo_tags;"`
	Priority    Priority       `json:"priority" gorm:"default:'medium'"`
	DueDate     *time.Time     `json:"due_date"`
}
 