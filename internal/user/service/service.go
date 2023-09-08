package service

import (
	"github.com/ell1jah/db_cp/internal/models"
	"github.com/ell1jah/db_cp/pkg/logger"
	"github.com/pkg/errors"
)

type UserRepo interface {
	Create(models.User) (int, error)
	GetByLoginAndPassword(string, string) (models.User, error)
}

type UserService struct {
	UserRepo UserRepo
	Logger   logger.Logger
}

func (us UserService) CreateUser(user models.User) (int, error) {
	id, err := us.UserRepo.Create(user)
	if err != nil {
		return -1, errors.Wrap(err, "can`t add user to repo")
	}

	return id, nil
}

func (us UserService) GetUserByLoginAndPassword(login, password string) (models.User, error) {
	user, err := us.UserRepo.GetByLoginAndPassword(login, password)
	if err != nil {
		return models.User{}, errors.Wrap(err, "can`t get user from repo")
	}

	return user, nil
}
