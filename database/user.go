package database

import (
	"context"
	"errors"

	"github.com/rafli-lutfi/perpustakaan/model"
	"gorm.io/gorm"
)

type UserDatabase interface {
	GetUserByID(ctx context.Context, id int) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int) error
}

type userDatabase struct {
	db *gorm.DB
}

func NewUserDatabase(db *gorm.DB) *userDatabase {
	return &userDatabase{db}
}

func (d *userDatabase) GetUserByID(ctx context.Context, id int) (model.User, error) {
	var user = model.User{}

	err := d.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, nil
	} else if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (d *userDatabase) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user = model.User{}

	err := d.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, nil
	} else if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (d *userDatabase) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	err := d.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (d *userDatabase) UpdateUser(ctx context.Context, user *model.User) error {
	err := d.db.WithContext(ctx).Model(&model.User{}).Updates(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *userDatabase) DeleteUser(ctx context.Context, id int) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		return err
	}

	return nil
}
