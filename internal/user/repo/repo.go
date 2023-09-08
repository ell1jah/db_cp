package repo

import (
	"github.com/ell1jah/db_cp/internal/models"
	"github.com/ell1jah/db_cp/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type PgUserRepo struct {
	Logger logger.Logger
	DB     *sqlx.DB
}

func (pur *PgUserRepo) Create(user models.User) (int, error) {
	var id int

	err := pur.DB.QueryRow(
		"insert into webUser (user_id, user_login, user_password, user_name, user_sex, user_role) "+
			"values ((select max(user_id) from webUser) + 1, $1, $2, $3, $4, $5) "+
			"returning user_id",
		user.Login,
		user.Password,
		user.Name,
		user.Sex,
		user.Role,
	).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "can`t insert to db")
	}

	return id, nil
}

func (pur *PgUserRepo) GetByLoginAndPassword(login, password string) (models.User, error) {
	user := models.User{}

	err := pur.DB.Get(
		&user,
		"select * "+
			"from webUser "+
			"where user_login = $1 and user_password = $2",
		login, password)
	if err != nil {
		return user, errors.Wrap(err, "can`t get from db")
	}

	return user, nil
}
