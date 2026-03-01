package repository

import (
	"practice4/internal/repository/_postgres"
	"practice4/internal/repository/_postgres/users"
	"practice4/pkg/modules"
)

type UserRepository interface {
	GetUsers(limit, offset int) ([]modules.User, error)
	GetUserByID(id int) (*modules.User, error)
	CreateUser(user modules.User) (int, error)
	UpdateUser(user modules.User) error
	DeleteUser(id int) (int64, error)
}

type Repositories struct {
	UserRepository
}

func NewRepositories(db *_postgres.Dialect) *Repositories {
	return &Repositories{users.NewUserRepository(db)}
}
