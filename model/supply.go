package model

import "time"

type Supply struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	IDBuku        int       `json:"id_buku"`
	TanggalTerima time.Time `json:"tanggal_terima"`
}
