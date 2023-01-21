package service

import (
	"context"
	"errors"

	"github.com/rafli-lutfi/perpustakaan/database"
	"github.com/rafli-lutfi/perpustakaan/model"
)

type JurusanService interface {
	GetAllJurusan(ctx context.Context) ([]model.Jurusan, error)
	NewJurusan(ctx context.Context, jurusan model.Jurusan) (model.Jurusan, error)
	UpdateJurusan(ctx context.Context, jurusan *model.Jurusan) (model.Jurusan, error)
	DeleteJurusan(ctx context.Context, id int) error
}

type jurusanService struct {
	JurusanDatabase database.JurusanDatabase
}

func NewJurusanService(jurusanDatabase database.JurusanDatabase) *jurusanService {
	return &jurusanService{jurusanDatabase}
}

func (s *jurusanService) GetAllJurusan(ctx context.Context) ([]model.Jurusan, error) {
	dbJurusan, err := s.JurusanDatabase.GetAllJurusan(ctx)
	if err != nil {
		return []model.Jurusan{}, err
	}

	return dbJurusan, nil
}

func (s *jurusanService) NewJurusan(ctx context.Context, jurusan model.Jurusan) (model.Jurusan, error) {
	newJurusan, err := s.JurusanDatabase.CreateJurusan(ctx, jurusan)
	if err != nil {
		return model.Jurusan{}, err
	}

	return newJurusan, nil
}

func (s *jurusanService) UpdateJurusan(ctx context.Context, jurusan *model.Jurusan) (model.Jurusan, error) {
	err := s.JurusanDatabase.UpdateJurusan(ctx, jurusan)
	if err != nil {
		return model.Jurusan{}, err
	}

	return *jurusan, nil
}

func (s *jurusanService) DeleteJurusan(ctx context.Context, id int) error {
	jurusan, err := s.JurusanDatabase.GetJurusanByID(ctx, id)
	if err != nil {
		return err
	}

	if jurusan.ID == 0 || jurusan.NamaJurusan == "" {
		return errors.New("jurusan not found")
	}

	err = s.JurusanDatabase.DeleteJurusan(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
