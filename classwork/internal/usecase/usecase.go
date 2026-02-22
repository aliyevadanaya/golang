package usecase

import "fmt"

type UserUseCase struct{}

func NewUserUseCase() UserUseCase {
	return UserUseCase{}
}

func (u UserUseCase) CreateUser(name string) string {
	fmt.Println("Good Morning!", name)

	result := fmt.Sprintf("Hello %s!", name)
	return result
}
