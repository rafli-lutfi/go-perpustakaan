package utils

import (
	"os"

	"github.com/rafli-lutfi/perpustakaan/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func ConnectToDatabase() {
	var err error
	db, err = gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        os.Getenv("DATABASE_URL"),
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tabel_",
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.User{}, model.Jurusan{})
}

func GetDBConnection() *gorm.DB {
	return db
}
