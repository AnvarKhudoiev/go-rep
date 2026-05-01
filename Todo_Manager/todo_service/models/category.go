package models

import "gorm.io/gorm"

type Category struct {
    gorm.Model  `json:"-"`
    Name        string `json:"name" gorm:"not null"` // Удалили unique!
    Color       string `json:"color"`
    UserID      uint   `json:"user_id"`
}