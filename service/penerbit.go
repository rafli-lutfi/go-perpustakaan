package service

import (
	"context"
	"errors"

	"github.com/rafli-lutfi/perpustakaan/database"
	"github.com/rafli-lutfi/perpustakaan/model"
)

type PenerbitService interface {
	GetPenerbitByID(ctx context.Context, id int) (model.Penerbit, error)
	GetAllPenerbit(ctx context.Context) ([]model.Penerbit, error)
	CreateNewPenerbit(ctx context.Context, penerbit model.Penerbit) (model.Penerbit, error)
	UpdatePenerbit(ctx context.Context, penerbit model.Penerbit) error
	DeletePenerbit(ctx context.Context, id int) error
}

type penerbitService struct {
	penerbitDatabase database.PenerbitDatabase
}

func NewPenerbitService(penerbitDatabase database.PenerbitDatabase) *penerbitService {
	return &penerbitService{penerbitDatabase}
}

func (s *penerbitService) GetPenerbitByID(ctx context.Context, id int) (model.Penerbit, error) {
	penerbit, err := s.penerbitDatabase.GetPenerbitByID(ctx, id)
	if err != nil {
		return model.Penerbit{}, err
	}

	if penerbit.ID == 0 || penerbit.NamaPenerbit == "" {
		return model.Penerbit{}, errors.New("penerbit not found")
	}

	return penerbit, nil
}

func (s *penerbitService) GetAllPenerbit(ctx context.Context) ([]model.Penerbit, error) {
	penerbitList, err := s.penerbitDatabase.GetAllPenerbit(ctx)
	if err != nil {
		return []model.Penerbit{}, err
	}

	return penerbitList, nil
}

func (s *penerbitService) CreateNewPenerbit(ctx context.Context, penerbit model.Penerbit) (model.Penerbit, error) {
	newPenerbit, err := s.penerbitDatabase.CreatePenerbit(ctx, penerbit)
	if err != nil {
		return model.Penerbit{}, err
	}

	return newPenerbit, nil
}

func (s *penerbitService) UpdatePenerbit(ctx context.Context, penerbit model.Penerbit) error {
	err := s.penerbitDatabase.UpdatePenerbit(ctx, penerbit)
	if err != nil {
		return err
	}

	return nil
}

func (s *penerbitService) DeletePenerbit(ctx context.Context, id int) error {
	penerbit, err := s.penerbitDatabase.GetPenerbitByID(ctx, id)
	if err != nil {
		return err
	}

	if penerbit.ID == 0 || penerbit.NamaPenerbit == "" {
		return errors.New("penerbit not found")
	}

	err = s.penerbitDatabase.DeletePenerbit(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
