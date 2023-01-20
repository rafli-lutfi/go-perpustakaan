package database

import (
	"context"
	"errors"

	"github.com/rafli-lutfi/perpustakaan/model"
	"gorm.io/gorm"
)

type PenerbitDatabase interface {
	GetPenerbitByID(ctx context.Context, id int) (model.Penerbit, error)
	GetAllPenerbit(ctx context.Context) ([]model.Penerbit, error)
	CreatePenerbit(ctx context.Context, penerbit model.Penerbit) (model.Penerbit, error)
	UpdatePenerbit(ctx context.Context, penerbit model.Penerbit) error
	DeletePenerbit(ctx context.Context, id int) error
}

type penerbitDatabase struct {
	db *gorm.DB
}

func NewPenerbitDatabase(db *gorm.DB) *penerbitDatabase {
	return &penerbitDatabase{db}
}

func (d *penerbitDatabase) GetPenerbitByID(ctx context.Context, id int) (model.Penerbit, error) {
	var penerbit model.Penerbit

	err := d.db.WithContext(ctx).Where("id = ?", id).First(&penerbit).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Penerbit{}, nil
	}
	if err != nil {
		return model.Penerbit{}, err
	}

	return penerbit, nil
}

func (d *penerbitDatabase) GetAllPenerbit(ctx context.Context) ([]model.Penerbit, error) {
	var penerbit []model.Penerbit

	rows, err := d.db.WithContext(ctx).Model(&model.Penerbit{}).Rows()
	if err != nil {
		return []model.Penerbit{}, err
	}

	defer rows.Close()

	for rows.Next() {
		d.db.ScanRows(rows, &penerbit)

	}

	return penerbit, nil
}

func (d *penerbitDatabase) CreatePenerbit(ctx context.Context, penerbit model.Penerbit) (model.Penerbit, error) {
	err := d.db.WithContext(ctx).Create(&penerbit).Error
	if err != nil {
		return model.Penerbit{}, err
	}

	return penerbit, nil
}

func (d *penerbitDatabase) UpdatePenerbit(ctx context.Context, penerbit model.Penerbit) error {
	err := d.db.WithContext(ctx).Model(&model.Penerbit{}).Where("id = ?", penerbit.ID).Updates(penerbit).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *penerbitDatabase) DeletePenerbit(ctx context.Context, id int) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Penerbit{}).Error
	if err != nil {
		return err
	}

	return nil
}
