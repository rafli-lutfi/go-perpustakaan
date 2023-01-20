package database

import (
	"context"
	"errors"

	"github.com/rafli-lutfi/perpustakaan/model"
	"gorm.io/gorm"
)

type AuthorDatabase interface {
	GetAuthorByID(ctx context.Context, id int) (model.Author, error)
	GetAllAuthor(ctx context.Context) ([]model.Author, error)
	CreateAuthor(ctx context.Context, author model.Author) (model.Author, error)
	UpdateAuthor(ctx context.Context, author model.Author) error
	DeleteAuthor(ctx context.Context, id int) error
}

type authorDatabase struct {
	db *gorm.DB
}

func NewAuthorDatabase(db *gorm.DB) *authorDatabase {
	return &authorDatabase{db}
}

func (d *authorDatabase) GetAuthorByID(ctx context.Context, id int) (model.Author, error) {
	var author model.Author

	err := d.db.WithContext(ctx).Where("id = ?", id).First(&author).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Author{}, nil
	}
	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

func (d *authorDatabase) GetAllAuthor(ctx context.Context) ([]model.Author, error) {
	var author []model.Author

	rows, err := d.db.WithContext(ctx).Model(&model.Author{}).Rows()
	if err != nil {
		return []model.Author{}, err
	}

	defer rows.Close()

	for rows.Next() {
		d.db.ScanRows(rows, &author)

	}

	return author, nil
}

func (d *authorDatabase) CreateAuthor(ctx context.Context, author model.Author) (model.Author, error) {
	err := d.db.WithContext(ctx).Create(&author).Error
	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

func (d *authorDatabase) UpdateAuthor(ctx context.Context, author model.Author) error {
	err := d.db.WithContext(ctx).Model(&model.Author{}).Where("id = ?", author.ID).Updates(author).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *authorDatabase) DeleteAuthor(ctx context.Context, id int) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Author{}).Error
	if err != nil {
		return err
	}

	return nil
}
