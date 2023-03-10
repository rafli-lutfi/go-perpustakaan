package database

import (
	"context"
	"errors"

	"github.com/rafli-lutfi/perpustakaan/model"
	"gorm.io/gorm"
)

type JurusanDatabase interface {
	GetJurusanByID(ctx context.Context, id int) (model.Jurusan, error)
	GetAllJurusan(ctx context.Context) ([]model.Jurusan, error)
	CreateJurusan(ctx context.Context, jurusan model.Jurusan) (model.Jurusan, error)
	UpdateJurusan(ctx context.Context, jurusan *model.Jurusan) error
	DeleteJurusan(ctx context.Context, id int) error
}

type jurusanDatabase struct {
	db *gorm.DB
}

func NewJurusanDatabase(db *gorm.DB) *jurusanDatabase {
	return &jurusanDatabase{db}
}

func (d *jurusanDatabase) GetJurusanByID(ctx context.Context, id int) (model.Jurusan, error) {
	jurusan := model.Jurusan{}

	err := d.db.WithContext(ctx).Where("id = ?", id).First(&jurusan).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Jurusan{}, nil
	} else if err != nil {
		return model.Jurusan{}, err
	}

	return jurusan, nil
}

func (d *jurusanDatabase) GetAllJurusan(ctx context.Context) ([]model.Jurusan, error) {
	jurusan := []model.Jurusan{}

	rows, err := d.db.WithContext(ctx).Model(&model.Jurusan{}).Rows()
	if err != nil {
		return []model.Jurusan{}, err
	}

	defer rows.Close()

	for rows.Next() {
		d.db.ScanRows(rows, &jurusan)
	}

	return jurusan, nil
}

func (d *jurusanDatabase) CreateJurusan(ctx context.Context, jurusan model.Jurusan) (model.Jurusan, error) {
	err := d.db.WithContext(ctx).Create(&jurusan).Error
	if err != nil {
		return model.Jurusan{}, err
	}

	return jurusan, nil
}

func (d *jurusanDatabase) UpdateJurusan(ctx context.Context, jurusan *model.Jurusan) error {
	err := d.db.WithContext(ctx).Model(&model.Jurusan{}).Where("id = ?", jurusan.ID).Updates(jurusan).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *jurusanDatabase) DeleteJurusan(ctx context.Context, id int) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Jurusan{}).Error
	if err != nil {
		return err
	}

	return nil
}
