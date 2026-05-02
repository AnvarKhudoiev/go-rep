package models

import "time"

type FavoriteBook struct {
	ID        uint      `json:"-" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"index;not null"`
	BookID    uint      `json:"book_id" gorm:"index;not null"`
	CreatedAt time.Time  `json:"created_at"`
}