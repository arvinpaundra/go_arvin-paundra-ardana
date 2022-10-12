package models

import "time"

type Barang struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	Nama       string    `gorm:"size:255" json:"nama" form:"nama"`
	Deskripsi  string    `json:"deskripsi" form:"deskripsi"`
	JumlahStok int       `json:"jumlah_stok" form:"jumlah_stok"`
	Harga      int       `json:"harga" form:"harga"`
	IdKategori string    `json:"id_kategori"`
	Kategori   Kategori  `gorm:"Foreignkey:IdKategori;association_foreignkey:ID" json:"kategori"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
