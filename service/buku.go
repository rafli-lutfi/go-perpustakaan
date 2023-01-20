package service

import (
	"context"

	"github.com/rafli-lutfi/perpustakaan/database"
	"github.com/rafli-lutfi/perpustakaan/model"
)

type BukuService interface {
	GetBukuByID(ctx context.Context, idBuku int) (model.BukuInfo, error)
	GetAllBuku(ctx context.Context) ([]model.Buku, error)
	CreateNewBuku(ctx context.Context, buku model.Buku) (model.Buku, error)
	UpdateBuku(ctx context.Context, buku model.Buku) error
	DeleteBuku(ctx context.Context, id int) error
}

type bukuService struct {
	bukuDatabase     database.BukuDatabase
	kategoriDatabase database.KategoriDatabase
	penerbitDatabase database.PenerbitDatabase
	authorDatabase   database.AuthorDatabase
}

func NewBukuService(bukuDatabase database.BukuDatabase, kategoriDatabase database.KategoriDatabase, penerbitDatabase database.PenerbitDatabase, authorDatabase database.AuthorDatabase) *bukuService {
	return &bukuService{bukuDatabase, kategoriDatabase, penerbitDatabase, authorDatabase}
}

func (s *bukuService) GetBukuByID(ctx context.Context, idBuku int) (model.BukuInfo, error) {
	buku, err := s.bukuDatabase.GetBukuByID(ctx, idBuku)
	if err != nil {
		return model.BukuInfo{}, err
	}

	kategori, err := s.kategoriDatabase.GetKategoriByID(ctx, buku.IDKategori)
	if err != nil {
		return model.BukuInfo{}, err
	}

	penerbit, err := s.penerbitDatabase.GetPenerbitByID(ctx, buku.IDPenerbit)
	if err != nil {
		return model.BukuInfo{}, err
	}

	author, err := s.authorDatabase.GetAuthorByID(ctx, buku.IDAuthor)
	if err != nil {
		return model.BukuInfo{}, err
	}

	bukuInfo := model.BukuInfo{
		ID:            buku.ID,
		JudulBuku:     buku.JudulBuku,
		TahunTerbit:   buku.TahunTerbit,
		Stock:         buku.Stock,
		NamaKategori:  kategori.NamaKategori,
		NamaPenerbit:  penerbit.NamaPenerbit,
		NamaPengarang: author.NamaPengarang,
	}

	return bukuInfo, nil
}

func (s *bukuService) GetAllBuku(ctx context.Context) ([]model.Buku, error) {
	listBuku, err := s.bukuDatabase.GetAllBuku(ctx)
	if err != nil {
		return []model.Buku{}, err
	}

	return listBuku, nil
}

func (s *bukuService) CreateNewBuku(ctx context.Context, buku model.Buku) (model.Buku, error) {
	newBuku, err := s.bukuDatabase.CreateBuku(ctx, buku)
	if err != nil {
		return model.Buku{}, err
	}

	return newBuku, nil
}

func (s *bukuService) UpdateBuku(ctx context.Context, buku model.Buku) error {
	err := s.bukuDatabase.UpdateBuku(ctx, buku)
	if err != nil {
		return err
	}

	return nil
}

func (s *bukuService) DeleteBuku(ctx context.Context, id int) error {
	err := s.bukuDatabase.DeleteBuku(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
