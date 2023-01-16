package service

import (
	"context"
	"errors"
	"time"

	"github.com/rafli-lutfi/perpustakaan/database"
	"github.com/rafli-lutfi/perpustakaan/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, user *model.User) (id int, err error)
	Register(ctx context.Context, user *model.User) (model.User, error)

	Delete(ctx context.Context, id int) error
	GetUserData(ctx context.Context, idUser int) (model.UserInfo, error)
}

type userService struct {
	userDatabase    database.UserDatabase
	jurusanDatabase database.JurusanDatabase
}

func NewUserService(userDatabase database.UserDatabase, jurusanDatabase database.JurusanDatabase) UserService {
	return &userService{userDatabase, jurusanDatabase}
}

func (s *userService) Login(ctx context.Context, user *model.User) (id int, err error) {
	dbUser, err := s.userDatabase.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return 0, err
	}

	if dbUser.Email == "" || dbUser.ID == 0 {
		return 0, errors.New("user not found")
	}

	// compare sent in password with saved password hash
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return 0, errors.New("wrong email or password")
	}

	return dbUser.ID, nil
}

func (s *userService) Register(ctx context.Context, user *model.User) (model.User, error) {
	dbUser, err := s.userDatabase.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return *user, err
	}

	if dbUser.Email != "" || dbUser.ID != 0 {
		return *user, err
	}

	user.CreatedAt = time.Now()

	newUser, err := s.userDatabase.CreateUser(ctx, *user)
	if err != nil {
		return *user, err
	}

	return newUser, nil
}

func (s *userService) Delete(ctx context.Context, id int) error {
	return s.userDatabase.DeleteUser(ctx, id)
}

func (s *userService) GetUserData(ctx context.Context, idUser int) (model.UserInfo, error) {
	dbUser, err := s.userDatabase.GetUserByID(ctx, idUser)
	if err != nil {
		return model.UserInfo{}, err
	}

	idJurusan := dbUser.IDJurusan

	dbJurusan, err := s.jurusanDatabase.GetJurusanByID(ctx, idJurusan)
	if err != nil {
		return model.UserInfo{}, err
	}

	userData := model.UserInfo{
		ID:          dbUser.ID,
		Fullname:    dbUser.Fullname,
		Address:     dbUser.Address,
		NPM:         dbUser.NPM,
		Email:       dbUser.Email,
		PhoneNumber: dbUser.PhoneNumber,
		NamaJurusan: dbJurusan.NamaJurusan,
	}

	return userData, nil
}

func (s *userService) Update(ctx context.Context, user *model.User) (model.User, error) {
	err := s.userDatabase.UpdateUser(ctx, user)
	if err != nil {
		return model.User{}, err
	}

	return *user, nil
}
