package api

import (
	"context"
	"gin-boilerplate/models"
	"gin-boilerplate/pkg/validators"
	"gin-boilerplate/repository"
	"gin-boilerplate/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

type UserService interface {
	Createuser(ctx context.Context)
	GetUser(ctx context.Context)
	DeleteUser(ctx context.Context)
	UpdateUser(ctx context.Context)
}

type UserHandler struct {
	UserRepo repository.UserRepository
}

func (u *UserHandler) Createuser(c *gin.Context) {

	var uservalidate validators.UserValidator

	err := c.ShouldBindJSON(&uservalidate)

	if err != nil {
		service.HTTPResponse(c, http.StatusNotFound, "error", err.Error())
		c.Abort()
		return
	}

	operation := models.User{
		Age:         uservalidate.Age,
		Name:        uservalidate.Name,
		Username:    uservalidate.Username,
		Email:       uservalidate.Email,
		PhoneNumber: uservalidate.PhoneNumber}

	_, somerror := u.UserRepo.Create(&operation)

	if somerror != nil {
		service.HTTPResponse(c, http.StatusNotFound, "error", somerror)
		c.Abort()
		return
	}

	service.HTTPResponse(c, http.StatusOK, "success", service.UserData(uservalidate))

	c.Abort()
	return
}

func (u *UserHandler) GetUser(c *gin.Context) {
	var uri validators.UserUriValidator
	var user models.User

	err := service.CheckUUID(c, &uri)

	if err != nil {

		service.HTTPResponse(c, http.StatusNotFound, "error", err.Error())

		c.Abort()
		return
	}

	err = u.UserRepo.GetOne(&user, uri.ID)

	if err != nil {
		service.HTTPResponse(c, http.StatusNotFound, "error", err.Error())
		c.Abort()
		return
	}

	service.HTTPResponse(c, http.StatusOK, "success", map[string]interface{}{
		"id":          user.ID,
		"name":        user.Name,
		"age":         user.Age,
		"email":       user.Email,
		"phone_numer": user.PhoneNumber,
		"username":    user.Username})

}

func (u *UserHandler) DeleteUser(c *gin.Context) {

	var uri validators.UserUriValidator
	var user models.User

	err := service.CheckUUID(c, &uri)

	if err != nil {

		service.HTTPResponse(c, http.StatusNotFound, "error", err.Error())

		c.Abort()
		return
	}

	err = u.UserRepo.Delete(&user, uri.ID)

	if err != nil {
		service.HTTPResponse(c, http.StatusNotFound, "error", err.Error())

		c.Abort()
		return
	}
	service.HTTPResponse(c, http.StatusOK, service.SuccessMessage, "Succefuly deleted")
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	var uri validators.UserUriValidator

	err := service.CheckUUID(c, &uri)

	if err != nil {

		service.HTTPResponse(c, http.StatusNotFound, service.ErrorMessage, err.Error())

		c.Abort()
		return
	}

	user := models.User{}

	err = u.UserRepo.Update(&user, uri.ID)

	if err != nil {

		service.HTTPResponse(c, http.StatusNotFound, service.ErrorMessage, err.Error())
		c.Abort()
		return
	}
	c.BindJSON(&user)

	service.HTTPResponse(c, http.StatusOK, service.SuccessMessage, map[string]interface{}{
		"id":          user.ID,
		"name":        user.Name,
		"age":         user.Age,
		"email":       user.Email,
		"phone_numer": user.PhoneNumber,
		"username":    user.Username})

}
