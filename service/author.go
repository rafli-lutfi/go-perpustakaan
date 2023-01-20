package service

import (
	"context"

	"github.com/rafli-lutfi/perpustakaan/database"
	"github.com/rafli-lutfi/perpustakaan/model"
)

type AuthorService interface {
	GetAuthorByID(ctx context.Context, id int) (model.Author, error)
	GetAllAuthor(ctx context.Context) ([]model.Author, error)
	CreateNewAuthor(ctx context.Context, author model.Author) (model.Author, error)
	UpdateAuthor(ctx context.Context, author model.Author) error
	DeleteAuthor(ctx context.Context, id int) error
}

type authorService struct {
	authorDatabase database.AuthorDatabase
}

func NewAuthorService(authorDatabase database.AuthorDatabase) *authorService {
	return &authorService{authorDatabase}
}

func (s *authorService) GetAuthorByID(ctx context.Context, id int) (model.Author, error) {
	author, err := s.authorDatabase.GetAuthorByID(ctx, id)
	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

func (s *authorService) GetAllAuthor(ctx context.Context) ([]model.Author, error) {
	auhtorList, err := s.authorDatabase.GetAllAuthor(ctx)
	if err != nil {
		return []model.Author{}, err
	}

	return auhtorList, nil
}

func (s *authorService) CreateNewAuthor(ctx context.Context, author model.Author) (model.Author, error) {
	newAuthor, err := s.authorDatabase.CreateAuthor(ctx, author)
	if err != nil {
		return model.Author{}, err
	}

	return newAuthor, nil
}

func (s *authorService) UpdateAuthor(ctx context.Context, author model.Author) error {
	err := s.authorDatabase.UpdateAuthor(ctx, author)
	if err != nil {
		return err
	}

	return nil
}

func (s *authorService) DeleteAuthor(ctx context.Context, id int) error {
	err := s.authorDatabase.DeleteAuthor(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
