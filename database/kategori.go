package database

import (
	"context"
	"errors"

	"github.com/rafli-lutfi/perpustakaan/model"
	"gorm.io/gorm"
)

type KategoriDatabase interface {
	GetKategoriByID(ctx context.Context, id int) (model.Kategori, error)
	GetAllKategori(ctx context.Context) ([]model.Kategori, error)
	CreateKategori(ctx context.Context, kategori model.Kategori) (model.Kategori, error)
	UpdateKategori(ctx context.Context, kategori model.Kategori) error
	DeleteKategori(ctx context.Context, id int) error
}

type kategoriDatabase struct {
	db *gorm.DB
}

func NewKategoriDatabase(db *gorm.DB) *kategoriDatabase {
	return &kategoriDatabase{db}
}

func (d *kategoriDatabase) GetKategoriByID(ctx context.Context, id int) (model.Kategori, error) {
	var kategori model.Kategori

	err := d.db.WithContext(ctx).Where("id = ?", id).First(&kategori).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Kategori{}, nil
	}
	if err != nil {
		return model.Kategori{}, err
	}

	return kategori, nil
}

func (d *kategoriDatabase) GetAllKategori(ctx context.Context) ([]model.Kategori, error) {
	var listKategori []model.Kategori

	rows, err := d.db.WithContext(ctx).Model(&model.Kategori{}).Rows()
	if err != nil {
		return []model.Kategori{}, err
	}

	defer rows.Close()

	for rows.Next() {
		d.db.ScanRows(rows, &listKategori)

	}

	return listKategori, nil
}

func (d *kategoriDatabase) CreateKategori(ctx context.Context, kategori model.Kategori) (model.Kategori, error) {
	err := d.db.WithContext(ctx).Create(&kategori).Error
	if err != nil {
		return model.Kategori{}, err
	}

	return kategori, nil
}

func (d *kategoriDatabase) UpdateKategori(ctx context.Context, kategori model.Kategori) error {
	err := d.db.WithContext(ctx).Model(&model.Kategori{}).Where("id = ?", kategori.ID).Updates(kategori).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *kategoriDatabase) DeleteKategori(ctx context.Context, id int) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Kategori{}).Error
	if err != nil {
		return err
	}

	return nil
}
