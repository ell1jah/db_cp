package repo

import (
	"fmt"

	"github.com/ell1jah/db_cp/internal/models"
	"github.com/ell1jah/db_cp/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type PgItemRepo struct {
	Logger logger.Logger
	DB     *sqlx.DB
}

func (pir *PgItemRepo) Create(item models.Item) (int, error) {
	var id int

	err := pir.DB.QueryRow(
		"insert into Item "+
			"values ((select max(id) from Item) + 1, $1, $2, $3, $4, $5, $6, $7) "+
			"returning id",
		item.Category,
		item.Size,
		item.Price,
		item.Sex,
		item.ImageID,
		item.BrandID,
		item.IsAvailable,
	).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "can`t insert to db")
	}

	return id, nil
}

func (pir *PgItemRepo) Get(id int) (models.Item, error) {
	item := models.Item{}

	err := pir.DB.Get(
		&item,
		"select * "+
			"from Item "+
			"where id = $1",
		id)
	if err != nil {
		return item, errors.Wrap(err, "can`t get from db")
	}

	return item, nil
}

func (pir *PgItemRepo) genGetAllQuery(params models.ItemsParams) string {
	base := "select * from Item"
	conds := []string{}

	if params.Brand >= 0 {
		conds = append(conds, fmt.Sprintf("brand_id = %d", params.Brand))
	}
	if params.Category != models.ItemsParamsAny {
		conds = append(conds, fmt.Sprintf("category = %s", params.Category))
	}
	if params.Sex != models.ItemsParamsAny {
		conds = append(conds, fmt.Sprintf("sex = %s", params.Sex))
	}

	if len(conds) != 0 {
		base += " where"

		for i := 0; i < len(conds)-1; i++ {
			base += " " + conds[i] + ","
		}

		base += " " + conds[len(conds)-1]
	}

	if params.Order != models.ItemsParamsAny {
		base += " order by price"

		if params.Order == models.ItemsOrderDesc {
			base += " desc"
		}
	}

	return base
}

func (pir *PgItemRepo) GetAll(params models.ItemsParams) ([]models.Item, error) {
	query := pir.genGetAllQuery(params)
	rows, err := pir.DB.Queryx(query)
	if err != nil {
		return nil, errors.Wrap(err, "can`t get from db, query: "+query)
	}

	items := []models.Item{}

	for rows.Next() {
		item := models.Item{}

		err := rows.StructScan(&item)
		if err != nil {
			return nil, errors.Wrap(err, "can`t scan struct from db query result")
		}

		items = append(items, item)
	}

	return items, nil
}

func (pir *PgItemRepo) Update(item models.Item) (models.Item, error) {
	_, err := pir.DB.Exec(
		"update Item "+
			"set category = $1, "+
			"size = $2, "+
			"price = $3, "+
			"sex = $4, "+
			"image_id = $5, "+
			"brand_id = $6, "+
			"is_available = $7 "+
			"where id = $8",
		item.Category,
		item.Size,
		item.Price,
		item.Sex,
		item.ImageID,
		item.BrandID,
		item.IsAvailable,
		item.ID)
	if err != nil {
		return item, errors.Wrap(err, "can`t update table in db")
	}

	return item, nil
}

func (pir *PgItemRepo) Delete(id int) error {
	_, err := pir.DB.Exec(
		"delete from Item "+
			"where id = $1",
		id)
	if err != nil {
		return errors.Wrap(err, "can`t delete from db")
	}

	return nil
}
