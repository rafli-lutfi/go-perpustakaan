package database

import (
	"context"
	"errors"

	"github.com/rafli-lutfi/perpustakaan/model"
	"gorm.io/gorm"
)

type BukuDatabase interface {
	GetBukuByID(ctx context.Context, id int) (model.Buku, error)
	GetAllBuku(ctx context.Context) ([]model.Buku, error)
	CreateBuku(ctx context.Context, buku model.Buku) (model.Buku, error)
	UpdateBuku(ctx context.Context, buku model.Buku) error
	DeleteBuku(ctx context.Context, id int) error
}

type bukuDatabase struct {
	db *gorm.DB
}

func NewBukuDatabase(db *gorm.DB) *bukuDatabase {
	return &bukuDatabase{db}
}

func (d *bukuDatabase) GetBukuByID(ctx context.Context, id int) (model.Buku, error) {
	var buku model.Buku

	err := d.db.WithContext(ctx).Where("id = ?", id).First(&buku).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Buku{}, nil
	}
	if err != nil {
		return model.Buku{}, err
	}

	return buku, nil
}

func (d *bukuDatabase) GetAllBuku(ctx context.Context) ([]model.Buku, error) {
	var listBuku []model.Buku

	rows, err := d.db.WithContext(ctx).Model(&model.Buku{}).Rows()
	if err != nil {
		return []model.Buku{}, err
	}

	defer rows.Close()

	for rows.Next() {
		d.db.ScanRows(rows, &listBuku)

	}

	return listBuku, nil
}

func (d *bukuDatabase) CreateBuku(ctx context.Context, buku model.Buku) (model.Buku, error) {
	err := d.db.WithContext(ctx).Create(&buku).Error
	if err != nil {
		return model.Buku{}, err
	}

	return buku, nil
}

func (d *bukuDatabase) UpdateBuku(ctx context.Context, buku model.Buku) error {
	err := d.db.WithContext(ctx).Model(&model.Buku{}).Where("id = ?", buku.ID).Updates(buku).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *bukuDatabase) DeleteBuku(ctx context.Context, id int) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Buku{}).Error
	if err != nil {
		return err
	}

	return nil
}
