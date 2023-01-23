package service

import (
	"context"
	"errors"

	"github.com/rafli-lutfi/perpustakaan/database"
	"github.com/rafli-lutfi/perpustakaan/model"
)

type BukuService interface {
	GetBukuByID(ctx context.Context, idBuku int) (model.BukuInfo, error)
	GetAllBuku(ctx context.Context, offset int, limit int) ([]model.BukuInfo, int, error)
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

	if buku.ID == 0 || buku.JudulBuku == "" {
		return model.BukuInfo{}, errors.New("buku not found")
	}

	kategori, err := s.kategoriDatabase.GetKategoriByID(ctx, buku.IDKategori)
	if err != nil {
		return model.BukuInfo{}, err
	}

	if kategori.ID == 0 || kategori.NamaKategori == "" {
		return model.BukuInfo{}, errors.New("kategori not found")
	}

	penerbit, err := s.penerbitDatabase.GetPenerbitByID(ctx, buku.IDPenerbit)
	if err != nil {
		return model.BukuInfo{}, err
	}

	if penerbit.ID == 0 || penerbit.NamaPenerbit == "" {
		return model.BukuInfo{}, errors.New("penerbit not found")
	}

	author, err := s.authorDatabase.GetAuthorByID(ctx, buku.IDAuthor)
	if err != nil {
		return model.BukuInfo{}, err
	}

	if author.ID == 0 || author.NamaPengarang == "" {
		return model.BukuInfo{}, errors.New("author not found")
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

func (s *bukuService) GetAllBuku(ctx context.Context, offset int, limit int) ([]model.BukuInfo, int, error) {
	listBuku, count, err := s.bukuDatabase.GetAllBuku(ctx, offset, limit)
	if err != nil {
		return []model.BukuInfo{}, 0, err
	}

	var listBukuInfo []model.BukuInfo

	for _, buku := range listBuku {
		kategori, err := s.kategoriDatabase.GetKategoriByID(ctx, buku.IDKategori)
		if err != nil {
			return []model.BukuInfo{}, 0, err
		}
		if kategori.ID == 0 || kategori.NamaKategori == "" {
			return []model.BukuInfo{}, 0, errors.New("kategori not found")
		}

		penerbit, err := s.penerbitDatabase.GetPenerbitByID(ctx, buku.IDPenerbit)
		if err != nil {
			return []model.BukuInfo{}, 0, err
		}
		if penerbit.ID == 0 || penerbit.NamaPenerbit == "" {
			return []model.BukuInfo{}, 0, errors.New("penerbit not found")
		}

		author, err := s.authorDatabase.GetAuthorByID(ctx, buku.IDAuthor)
		if err != nil {
			return []model.BukuInfo{}, 0, err
		}
		if author.ID == 0 || author.NamaPengarang == "" {
			return []model.BukuInfo{}, 0, errors.New("author not found")
		}

		bukuinfo := model.BukuInfo{
			ID:            buku.ID,
			JudulBuku:     buku.JudulBuku,
			TahunTerbit:   buku.TahunTerbit,
			Stock:         buku.Stock,
			NamaKategori:  kategori.NamaKategori,
			NamaPenerbit:  penerbit.NamaPenerbit,
			NamaPengarang: author.NamaPengarang,
		}

		listBukuInfo = append(listBukuInfo, bukuinfo)
	}

	return listBukuInfo, count, nil
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
	buku, err := s.bukuDatabase.GetBukuByID(ctx, id)
	if err != nil {
		return err
	}

	if buku.ID == 0 || buku.JudulBuku == "" {
		return errors.New("buku not found")
	}

	err = s.bukuDatabase.DeleteBuku(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
