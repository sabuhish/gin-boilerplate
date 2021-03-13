package repository

import (
	"gin-boilerplate/models"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (o *UserRepoMock) Create(u *models.User) (*models.User, error) {

	return u, nil
}

func (o *UserRepoMock) GetOne(u *models.User, id string) error {

	return nil
}

func (o *UserRepoMock) Delete(u *models.User, id string) error {

	return nil

}

func (o *UserRepoMock) Update(u *models.User, id string) error {

	return nil
}
