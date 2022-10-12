package models

import "time"

type Kategori struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
