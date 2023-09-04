package repo

import (
	"sync"

	"github.com/pkg/errors"
)

type UserRepo struct {
	users []common.User
	mu    sync.RWMutex
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		mu:    sync.RWMutex{},
		users: []common.User{},
	}
}

func (ur *UserRepo) Add(user common.UserFullData) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	userStruct := common.User{
		Username: user.GetUsername(),
		Password: user.GetPassword(),
		ID:       user.GetID(),
	}

	ur.users = append(ur.users, userStruct)
	return nil
}

func (ur *UserRepo) IsExistByUsername(username string) (bool, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	for _, user := range ur.users {
		if user.Username == username {
			return true, nil
		}
	}

	return false, nil
}

func (ur *UserRepo) GetByRegForm(desired common.UserRegForm) (common.UserFullData, error) {
	ur.mu.RLock()
	defer ur.mu.RUnlock()

	for _, user := range ur.users {
		if user.Username == desired.GetUsername() && user.Password == desired.GetPassword() {
			return user, nil
		}
	}

	return nil, errors.Errorf("user not found")
}
