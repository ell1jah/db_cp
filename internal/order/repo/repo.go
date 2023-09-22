package repo

import (
	"github.com/ell1jah/db_cp/internal/models"
	"github.com/ell1jah/db_cp/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type PgOrderRepo struct {
	Logger logger.Logger
	DB     *sqlx.DB
}

func (por *PgOrderRepo) Get(id int) (models.Order, error) {
	order := models.Order{}

	err := por.DB.Get(
		&order,
		"select * "+
			"from Ordering "+
			"where id = $1",
		id)
	if err != nil {
		return order, errors.Wrap(err, "can`t get from db")
	}

	rows, err := por.DB.Queryx(`SELECT 
	i.id, 
	i.category, 
	i.size, 
	i.price, 
	i.sex, 
	i.image_id, 
	i.brand_id, 
	i.is_available, 
	o.amount 
  FROM 
	OrderItems o 
	JOIN Item i ON o.item_id = i.id 
  where 
	o.order_id = $1;`, id)
	if err != nil {
		return order, errors.Wrap(err, "can`t get from db")
	}

	items := []models.OrderItem{}

	for rows.Next() {
		item := models.OrderItem{}

		err := rows.StructScan(&item)
		if err != nil {
			return order, errors.Wrap(err, "can`t scan struct from db query result")
		}

		items = append(items, item)
	}

	order.Items = items

	return order, nil
}

func (por *PgOrderRepo) GetAll() ([]models.Order, error) {
	orders := []models.Order{}

	rows, err := por.DB.Queryx("select * " +
		"from Ordering where current_status != 'корзина'")
	if err != nil {
		return orders, errors.Wrap(err, "can`t get from db")
	}

	for rows.Next() {
		order := models.Order{}

		err := rows.StructScan(&order)
		if err != nil {
			return orders, errors.Wrap(err, "can`t scan struct from db query result")
		}

		orders = append(orders, order)
	}

	for i := range orders {
		rows, err := por.DB.Queryx(`SELECT 
		i.id, 
		i.category, 
		i.size, 
		i.price, 
		i.sex, 
		i.image_id, 
		i.brand_id, 
		i.is_available, 
		o.amount 
	  FROM 
		OrderItems o 
		JOIN Item i ON o.item_id = i.id 
	  where 
		o.order_id = $1;`, orders[i].ID)
		if err != nil {
			return orders, errors.Wrap(err, "can`t get from db")
		}

		items := []models.OrderItem{}

		for rows.Next() {
			item := models.OrderItem{}

			err := rows.StructScan(&item)
			if err != nil {
				return orders, errors.Wrap(err, "can`t scan struct from db query result")
			}

			items = append(items, item)
		}

		orders[i].Items = items
	}

	return orders, nil
}

func (por *PgOrderRepo) GetUsersAll(user int) ([]models.Order, error) {
	orders := []models.Order{}

	rows, err := por.DB.Queryx("select * "+
		"from Ordering where user_id = $1 and current_status != 'корзина'", user)
	if err != nil {
		return orders, errors.Wrap(err, "can`t get from db")
	}

	for rows.Next() {
		order := models.Order{}

		err := rows.StructScan(&order)
		if err != nil {
			return orders, errors.Wrap(err, "can`t scan struct from db query result")
		}

		orders = append(orders, order)
	}

	for i := range orders {
		rows, err := por.DB.Queryx(`SELECT 
		i.id, 
		i.category, 
		i.size, 
		i.price, 
		i.sex, 
		i.image_id, 
		i.brand_id, 
		i.is_available, 
		o.amount 
	  FROM 
		OrderItems o 
		JOIN Item i ON o.item_id = i.id 
	  where 
		o.order_id = $1;`, orders[i].ID)
		if err != nil {
			return orders, errors.Wrap(err, "can`t get from db")
		}

		items := []models.OrderItem{}

		for rows.Next() {
			item := models.OrderItem{}

			err := rows.StructScan(&item)
			if err != nil {
				return orders, errors.Wrap(err, "can`t scan struct from db query result")
			}

			items = append(items, item)
		}

		orders[i].Items = items
	}

	return orders, nil
}

func (por *PgOrderRepo) Update(order models.Order) (models.Order, error) {
	_, err := por.DB.Exec(
		"update Ordering "+
			"set commit_date = $1, "+
			"user_id = $2, "+
			"price = $3, "+
			"current_status = $4 "+
			"where id = $5",
		order.Date,
		order.UserID,
		order.Price,
		order.Status,
		order.ID)
	if err != nil {
		return order, errors.Wrap(err, "can`t update table in db")
	}

	return order, nil
}
