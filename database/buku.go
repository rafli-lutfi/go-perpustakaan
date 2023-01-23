package database

import (
	"context"
	"errors"

	"github.com/rafli-lutfi/perpustakaan/model"
	"gorm.io/gorm"
)

type BukuDatabase interface {
	GetBukuByID(ctx context.Context, id int) (model.Buku, error)
	GetAllBuku(ctx context.Context, offset int, limit int) ([]model.Buku, int, error)
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

func (d *bukuDatabase) GetAllBuku(ctx context.Context, offset int, limit int) ([]model.Buku, int, error) {
	var listBuku []model.Buku

	// count all data
	data := d.db.WithContext(ctx).Find(&[]model.Buku{})
	count := int(data.RowsAffected)

	// get data by offset and limit
	err := d.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&listBuku).Error
	if err != nil {
		return []model.Buku{}, 0, err
	}

	return listBuku, count, nil
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
