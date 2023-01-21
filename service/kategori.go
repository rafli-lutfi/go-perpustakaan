package service

import (
	"context"
	"errors"

	"github.com/rafli-lutfi/perpustakaan/database"
	"github.com/rafli-lutfi/perpustakaan/model"
)

type KategoriService interface {
	GetKategoriByID(ctx context.Context, id int) (model.Kategori, error)
	GetAllKategori(ctx context.Context) ([]model.Kategori, error)
	CreateNewKategori(ctx context.Context, kategori model.Kategori) (model.Kategori, error)
	UpdateKategori(ctx context.Context, kategori model.Kategori) error
	DeleteKategori(ctx context.Context, id int) error
}

type kategoriService struct {
	kategoriDatabase database.KategoriDatabase
}

func NewKategoriService(kategoriDatabase database.KategoriDatabase) *kategoriService {
	return &kategoriService{kategoriDatabase}
}

func (s *kategoriService) GetKategoriByID(ctx context.Context, id int) (model.Kategori, error) {
	kategori, err := s.kategoriDatabase.GetKategoriByID(ctx, id)
	if err != nil {
		return model.Kategori{}, err
	}

	if kategori.ID == 0 || kategori.NamaKategori == "" {
		return model.Kategori{}, errors.New("kategori not found")
	}

	return kategori, nil
}

func (s *kategoriService) GetAllKategori(ctx context.Context) ([]model.Kategori, error) {
	kategoriList, err := s.kategoriDatabase.GetAllKategori(ctx)
	if err != nil {
		return []model.Kategori{}, err
	}

	return kategoriList, nil
}

func (s *kategoriService) CreateNewKategori(ctx context.Context, kategori model.Kategori) (model.Kategori, error) {
	newKategori, err := s.kategoriDatabase.CreateKategori(ctx, kategori)
	if err != nil {
		return model.Kategori{}, err
	}

	return newKategori, nil
}

func (s *kategoriService) UpdateKategori(ctx context.Context, kategori model.Kategori) error {
	err := s.kategoriDatabase.UpdateKategori(ctx, kategori)
	if err != nil {
		return err
	}

	return nil
}

func (s *kategoriService) DeleteKategori(ctx context.Context, id int) error {
	kategori, err := s.kategoriDatabase.GetKategoriByID(ctx, id)
	if err != nil {
		return err
	}

	if kategori.ID == 0 || kategori.NamaKategori == "" {
		return errors.New("kategori not found")
	}

	err = s.kategoriDatabase.DeleteKategori(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
