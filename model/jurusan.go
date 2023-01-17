package model

import "time"

type Jurusan struct {
	ID          int    `json:"id" gorm:"primarykey"`
	NamaJurusan string `json:"nama_jurusan" gorm:"type:varchar(255);not null;uniqueIndex"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
