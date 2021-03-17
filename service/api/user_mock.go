package api

import (
	"gin-boilerplate/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(u *models.User) (*models.User, error) {

	args := m.Called(u)

	return args.Get(0).(*models.User), args.Error(1)

}

func (m *MockUserRepository) GetOne(user *models.User, id string)  error  {

	args := m.Called(user, id)

	return args.Error(0)

}

func (m *MockUserRepository) Delete(user *models.User, id string) error {

	args := m.Called(user, id)

	return args.Error(0)

}

func (m *MockUserRepository) Update(user *models.User, id string) error {

	args := m.Called(user, id)

	return args.Error(0)

}
