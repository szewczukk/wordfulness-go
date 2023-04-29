package storage

import (
	"errors"
	"wordfulness/types"
)

type UserMemoryStorage struct {
	users []*types.User
}

func NewUserMemoryStorage(users []*types.User) *UserMemoryStorage {
	return &UserMemoryStorage{
		users: users,
	}
}

func (s *UserMemoryStorage) GetUserByUserName(username string) (*types.User, error) {
	for _, user := range s.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("not found")
}

func (s *UserMemoryStorage) CreateUser(
	username string,
	password string,
) error {
	for _, user := range s.users {
		if user.Username == username {
			return errors.New("duplicate")
		}
	}

	s.users = append(s.users, &types.User{
		Id:       len(s.users),
		Username: username,
		Password: password,
	})

	return nil
}
