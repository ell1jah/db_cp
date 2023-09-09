package repo

import (
	"github.com/ell1jah/db_cp/internal/models"
	"github.com/ell1jah/db_cp/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type PgBrandRepo struct {
	Logger logger.Logger
	DB     *sqlx.DB
}

func (pbr *PgBrandRepo) Create(brand models.Brand) (int, error) {
	var id int

	err := pbr.DB.QueryRow(
		"insert into Brand "+
			"values ((select max(id) from Brand) + 1, $1, $2, $3, $4) "+
			"returning id",
		brand.Name,
		brand.Year,
		brand.Logo,
		brand.Owner,
	).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "can`t insert to db")
	}

	return id, nil
}

func (pbr *PgBrandRepo) Get(id int) (models.Brand, error) {
	brand := models.Brand{}

	err := pbr.DB.Get(
		&brand,
		"select * "+
			"from Brand "+
			"where id = $1",
		id)
	if err != nil {
		return brand, errors.Wrap(err, "can`t get from db")
	}

	return brand, nil
}

func (pbr *PgBrandRepo) Update(brand models.Brand) (models.Brand, error) {
	_, err := pbr.DB.Exec(
		"update Brand "+
			"set brand_name = $1, "+
			"founding_year = $2, "+
			"logo_id = $3, "+
			"brand_owner = $4"+
			"where id = $5",
		brand.Name,
		brand.Year,
		brand.Logo,
		brand.Owner,
		brand.ID)
	if err != nil {
		return brand, errors.Wrap(err, "can`t update table in db")
	}

	return brand, nil
}

func (pbr *PgBrandRepo) Delete(id int) error {
	_, err := pbr.DB.Exec(
		"delete from Brand "+
			"where id = $1",
		id)
	if err != nil {
		return errors.Wrap(err, "can`t delete from db")
	}

	return nil
}
