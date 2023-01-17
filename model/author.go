package model

import "time"

type Author struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	NamaKategori string `json:"nama_kategori" gorm:"type:varchar(50); not null; unique"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
