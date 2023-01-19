package model

import "time"

type Author struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	NamaPengarang string `json:"nama_pengarang" gorm:"type:varchar(50); not null; unique"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
