package service

import (
	"github.com/ell1jah/db_cp/internal/models"
	"github.com/ell1jah/db_cp/pkg/logger"
	"github.com/pkg/errors"
)

type BrandRepo interface {
	Create(models.Brand) (int, error)
	Get(int) (models.Brand, error)
	Update(models.Brand) (models.Brand, error)
	Delete(int) error
}

type BrandService struct {
	BrandRepo BrandRepo
	Logger    logger.Logger
}

func (br BrandService) Create(brand models.Brand) (int, error) {
	id, err := br.BrandRepo.Create(brand)
	if err != nil {
		return -1, errors.Wrap(err, "can`t add to repo")
	}

	return id, nil
}

func (br BrandService) Get(id int) (models.Brand, error) {
	brand, err := br.BrandRepo.Get(id)
	if err != nil {
		return models.Brand{}, errors.Wrap(err, "can`t get from repo")
	}

	return brand, nil
}

func (br BrandService) Update(brand models.Brand) (models.Brand, error) {
	brand, err := br.BrandRepo.Update(brand)
	if err != nil {
		return brand, errors.Wrap(err, "can`t update repo")
	}

	return brand, nil
}

func (br BrandService) Delete(id int) error {
	err := br.BrandRepo.Delete(id)
	if err != nil {
		return errors.Wrap(err, "can`t delete from repo")
	}

	return nil
}
