package repository

import (
	"rp2/internal/repository/_postgres"
	"rp2/internal/repository/_postgres/users"
	"rp2/pkg/modules"
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
