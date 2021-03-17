package routers

import (
	"gin-boilerplate/models"
	"gin-boilerplate/pkg/validators"
	"gin-boilerplate/repository"
	"gin-boilerplate/service"
	"gin-boilerplate/service/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func UserServiceHandler() *api.UserHandler {
	service := api.UserHandler{
		UserRepo: &repository.DBUserRepositoty{},
	}
	return &service
}

func Create(c *gin.Context) {

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

	_, err = UserServiceHandler().Createuser(operation)

	if err != nil {
		service.HTTPResponse(c, http.StatusNotFound, "error", err)
		c.Abort()
		return
	}

	service.HTTPResponse(c, http.StatusOK, "success", service.UserData(uservalidate))

	c.Abort()
	return
}

func GetOne(c *gin.Context) {
	var uri validators.UserUriValidator
	var user models.User

	err := service.CheckUUID(c, &uri)

	if err != nil {

		service.HTTPResponse(c, http.StatusNotFound, "error", err.Error())

		c.Abort()
		return
	}

	_, err = UserServiceHandler().GetUser(&user, uri.ID)

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

func DeleteUser(c *gin.Context) {

	var uri validators.UserUriValidator
	var user models.User

	err := service.CheckUUID(c, &uri)

	if err != nil {

		service.HTTPResponse(c, http.StatusNotFound, service.ErrorMessage, err.Error())

		c.Abort()
		return
	}

	err = UserServiceHandler().DeleteUser(&user, uri.ID)

	if err != nil {
		service.HTTPResponse(c, http.StatusNotFound, service.ErrorMessage, err.Error())

		c.Abort()
		return
	}
	service.HTTPResponse(c, http.StatusOK, service.SuccessMessage, "Succefuly deleted")
}

func UpdateUser(c *gin.Context) {
	var uri validators.UserUriValidator

	err := service.CheckUUID(c, &uri)

	if err != nil {

		service.HTTPResponse(c, http.StatusNotFound, service.ErrorMessage, err.Error())

		c.Abort()
		return
	}

	user := models.User{}

	_, err = UserServiceHandler().UpdateUser(&user, uri.ID)

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

func init() {

	Router = gin.Default()

	r := Router.Group("/api")

	{
		r.GET("/ping", api.Ping)

	}
	u := Router.Group("/users")
	{
		u.POST("/", Create)

		u.GET("/:id", GetOne)
		u.DELETE("/:id", DeleteUser)
		u.PUT("/:id", UpdateUser)

		// u.POST("/", UserServiceHandler().Createuser)
	}
}
