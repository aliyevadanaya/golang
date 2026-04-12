package repo

import (
	"practice7/internal/entity"
	"practice7/pkg"
)

type UserRepo struct {
	PG *pkg.Postgres
}

func NewUserRepo(pg *pkg.Postgres) *UserRepo {
	return &UserRepo{PG: pg}
}

func (u *UserRepo) RegisterUser(user *entity.User) (*entity.User, error) {
	err := u.PG.Conn.Create(user).Error
	return user, err
}

func (u *UserRepo) LoginUser(dto *entity.LoginUserDTO) (*entity.User, error) {
	var user entity.User
	err := u.PG.Conn.Where("username = ?", dto.Username).First(&user).Error
	return &user, err
}

func (u *UserRepo) GetByID(id string) (*entity.User, error) {
	var user entity.User
	err := u.PG.Conn.First(&user, "id = ?", id).Error
	return &user, err
}

func (u *UserRepo) PromoteUser(id string) error {
	return u.PG.Conn.Model(&entity.User{}).
		Where("id = ?", id).
		Update("role", "admin").Error
}

func (u *UserRepo) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	err := u.PG.Conn.Find(&users).Error
	return users, err
}
