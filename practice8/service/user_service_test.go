package service

import (
	"fmt"
	"testing"

	"practice8/repository"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	user := &repository.User{ID: 1, Name: "Test"}

	mockRepo.EXPECT().GetByEmail("test@mail.com").Return(user, nil)
	err := service.RegisterUser(user, "test@mail.com")
	assert.Error(t, err)

	mockRepo.EXPECT().GetByEmail("new@mail.com").Return(nil, nil)
	mockRepo.EXPECT().CreateUser(user).Return(nil)
	err = service.RegisterUser(user, "new@mail.com")
	assert.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	err := service.DeleteUser(1)
	assert.Error(t, err)

	mockRepo.EXPECT().DeleteUser(2).Return(nil)
	err = service.DeleteUser(2)
	assert.NoError(t, err)
}

func TestRegisterUser_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	user := &repository.User{ID: 1, Name: "Test"}

	mockRepo.EXPECT().GetByEmail("err@mail.com").Return(nil, fmt.Errorf("db error"))

	err := service.RegisterUser(user, "err@mail.com")

	assert.Error(t, err)
}

func TestUpdateUserName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	err := service.UpdateUserName(1, "")
	assert.Error(t, err)

	mockRepo.EXPECT().GetUserByID(2).Return(nil, fmt.Errorf("not found"))
	err = service.UpdateUserName(2, "NewName")
	assert.Error(t, err)

	user := &repository.User{ID: 3, Name: "Old"}
	mockRepo.EXPECT().GetUserByID(3).Return(user, nil)
	mockRepo.EXPECT().UpdateUser(user).Return(nil)
	err = service.UpdateUserName(3, "New")
	assert.NoError(t, err)

	user2 := &repository.User{ID: 4, Name: "Old"}
	mockRepo.EXPECT().GetUserByID(4).Return(user2, nil)
	mockRepo.EXPECT().UpdateUser(user2).Return(fmt.Errorf("update error"))
	err = service.UpdateUserName(4, "New")
	assert.Error(t, err)
}

func TestDeleteUser_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	service := NewUserService(mockRepo)

	mockRepo.EXPECT().DeleteUser(5).Return(fmt.Errorf("db error"))

	err := service.DeleteUser(5)

	assert.Error(t, err)
}
