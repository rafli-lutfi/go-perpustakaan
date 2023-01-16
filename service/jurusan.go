package service

import (
	"context"

	"github.com/rafli-lutfi/perpustakaan/database"
	"github.com/rafli-lutfi/perpustakaan/model"
)

type JurusanService interface {
	GetAllJurusan(ctx context.Context) ([]model.Jurusan, error)
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
