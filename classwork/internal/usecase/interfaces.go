package usecase

type UserInterface interface {
	CreateUser(name string) string
}
