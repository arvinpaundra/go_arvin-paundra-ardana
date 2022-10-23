package model

import (
	"time"
)

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Fullname  string    `gorm:"size:255" json:"fullname"`
	Email     string    `gorm:"size:255" json:"email"`
	Password  string    `gorm:"size:255" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
