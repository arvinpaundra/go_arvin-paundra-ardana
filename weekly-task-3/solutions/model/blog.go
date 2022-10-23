package model

import "time"

type Blog struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	UserId     string    `gorm:"size:255" json:"user_id"`
	CategoryId string    `gorm:"size:255" json:"category_id"`
	Title      string    `gorm:"size:255" json:"title"`
	Content    string    `json:"content"`
	User       User      `json:"user"`
	Category   Category  `json:"category"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
