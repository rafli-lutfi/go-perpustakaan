package model

type Jurusan struct {
	ID          int    `json:"id" gorm:"primarykey"`
	NamaJurusan string `json:"nama_jurusan" gorm:"type:varchar(255);not null;uniqueIndex"`
}
