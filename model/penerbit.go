package model

import "time"

type Penerbit struct {
	ID              int    `json:"id" gorm:"primaryKey"`
	NamaPenerbit    string `json:"nama_penerbit" gorm:"type:varchar(50); not null; unique"`
	AlamatPenerbit  string `json:"alamat_penerbit" type:"varchar(100); unique"`
	TeleponPenerbit string `json:"telepon_penerbit" type:"varchar(25); unique"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
