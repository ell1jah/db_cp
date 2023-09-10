package service

import (
	"github.com/ell1jah/db_cp/internal/models"
	"github.com/ell1jah/db_cp/pkg/logger"
	"github.com/pkg/errors"
)

type ItemRepo interface {
	Create(models.Item) (int, error)
	Get(int) (models.Item, error)
	GetAll(models.ItemsParams) ([]models.Item, error)
	Update(models.Item) (models.Item, error)
	Delete(int) error
}

type ItemService struct {
	ItemRepo ItemRepo
	Logger   logger.Logger
}

func (is ItemService) Create(item models.Item) (int, error) {
	id, err := is.ItemRepo.Create(item)
	if err != nil {
		return -1, errors.Wrap(err, "can`t add to repo")
	}

	return id, nil
}

func (is ItemService) Get(id int) (models.Item, error) {
	item, err := is.ItemRepo.Get(id)
	if err != nil {
		return models.Item{}, errors.Wrap(err, "can`t get from repo")
	}

	return item, nil
}

func (is ItemService) GetAll(params models.ItemsParams) ([]models.Item, error) {
	items, err := is.ItemRepo.GetAll(params)
	if err != nil {
		return nil, errors.Wrap(err, "can`t get from repo")
	}

	return items, nil
}

func (is ItemService) Update(item models.Item) (models.Item, error) {
	item, err := is.ItemRepo.Update(item)
	if err != nil {
		return item, errors.Wrap(err, "can`t update repo")
	}

	return item, nil
}

func (is ItemService) Delete(id int) error {
	err := is.ItemRepo.Delete(id)
	if err != nil {
		return errors.Wrap(err, "can`t delete from repo")
	}

	return nil
}
