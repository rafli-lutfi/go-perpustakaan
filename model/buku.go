package model

import "time"

type Buku struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	JudulBuku   string `json:"judul_buku" gorm:"type:varchar(100); not null; unique"`
	TahunTerbit int    `json:"tahun_terbit" gorm:"not null"`
	Stock       int    `json:"stock" gorm:"not null"`
	IDKategori  int    `json:"id_kategori" gorm:"not null"`
	IDPenerbit  int    `json:"id_penerbit" gorm:"not null"`
	IDAuthor    int    `json:"id_author" gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BukuInfo struct {
	ID            int    `json:"id" `
	JudulBuku     string `json:"judul_buku"`
	TahunTerbit   int    `json:"tahun_terbit"`
	Stock         int    `json:"stock"`
	NamaKategori  string `json:"nama_kategori"`
	NamaPenerbit  string `json:"nama_penerbit"`
	NamaPengarang string `json:"nama_author"`
}
