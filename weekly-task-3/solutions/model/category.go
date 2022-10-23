package model

import "time"

type Category struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	CategoryName string    `gorm:"size:255" json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
