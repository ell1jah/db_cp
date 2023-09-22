package service

import (
	"github.com/ell1jah/db_cp/internal/models"
	"github.com/ell1jah/db_cp/pkg/logger"
	"github.com/pkg/errors"
)

type BasketRepo interface {
	Get(int) (models.Basket, error)
	Commit(int) error
	AddItem(int, int) error
	DecItem(int, int) error
}

type BasketService struct {
	BasketRepo BasketRepo
	Logger     logger.Logger
}

func (bs BasketService) Get(id int) (models.Basket, error) {
	basket, err := bs.BasketRepo.Get(id)
	if err != nil {
		return models.Basket{}, errors.Wrap(err, "can`t get from repo")
	}

	return basket, nil
}

func (bs BasketService) Commit(id int) error {
	err := bs.BasketRepo.Commit(id)
	if err != nil {
		return errors.Wrap(err, "can`t commit in repo")
	}

	return nil
}

func (bs BasketService) AddItem(itemID, userID int) error {
	err := bs.BasketRepo.AddItem(itemID, userID)
	if err != nil {
		return errors.Wrap(err, "can`t add item in repo")
	}

	return nil
}

func (bs BasketService) DecItem(itemID, userID int) error {
	err := bs.BasketRepo.DecItem(itemID, userID)
	if err != nil {
		return errors.Wrap(err, "can`t add item in repo")
	}

	return nil
}
