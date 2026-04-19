package service

import (
	"errors"
	"fmt"
	"practice8/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id int) (*repository.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) CreateUser(user *repository.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) RegisterUser(user *repository.User, email string) error {
	existing, err := s.repo.GetByEmail(email)
	if existing != nil {
		return fmt.Errorf("user already exists")
	}
	if err != nil {
		return err
	}
	return s.repo.CreateUser(user)
}

func (s *UserService) UpdateUserName(id int, newName string) error {
	if newName == "" {
		return fmt.Errorf("name cannot be empty")
	}

	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}

	user.Name = newName
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	if id == 1 {
		return errors.New("cannot delete admin")
	}
	return s.repo.DeleteUser(id)
}
