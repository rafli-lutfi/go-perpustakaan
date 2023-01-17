package model

import "time"

type User struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Fullname    string    `json:"fullname" gorm:"type:varchar(50);not null;unique"`
	Address     string    `json:"address" gorm:"type:varchar(100);not null"`
	Password    string    `json:"-" gorm:"type:varchar(100);not null"`
	NPM         string    `json:"npm" gorm:"type:varchar(25);not null;unique"`
	Email       string    `json:"email" gorm:"type:varchar(50);not null;unique"`
	IDJurusan   int       `json:"id_jurusan" gorm:"not null"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(25); not null;"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserInfo struct {
	ID          int    `json:"id"`
	Fullname    string `json:"fullname"`
	Address     string `json:"address"`
	NPM         string `json:"npm"`
	Email       string `json:"email"`
	NamaJurusan string `json:"nama_jurusan"`
	PhoneNumber string `json:"phone_number"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Fullname    string `json:"fullname" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Address     string `json:"address" binding:"required"`
	IDJurusan   int    `json:"id_jurusan" binding:"required"`
	NPM         string `json:"npm" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}
