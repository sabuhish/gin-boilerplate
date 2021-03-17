package api

import (
	"gin-boilerplate/models"
	"gin-boilerplate/repository"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

type UserService interface {
	Createuser(m models.User) (models.User, error)
	GetUser(m models.User) (*models.User, error)
	DeleteUser(m models.User) error
	UpdateUser(m models.User) (*models.User, error)
}

type UserHandler struct {
	UserRepo repository.UserRepository
}

func (u *UserHandler) Createuser(m models.User) (models.User, error) {

	_, err := u.UserRepo.Create(&m)

	return m, err
}

func (u *UserHandler) GetUser(m *models.User, uri string) (*models.User, error) {

	err := u.UserRepo.GetOne(m, uri)

	return m, err
}

func (u *UserHandler) DeleteUser(m *models.User, uri string) error {

	err := u.UserRepo.Delete(m, uri)

	return err
}

func (u *UserHandler) UpdateUser(m *models.User, uri string)  (*models.User, error){

	err := u.UserRepo.Update(m, uri)

	return m, err

}
