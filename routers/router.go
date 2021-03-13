package routers

import (
	"gin-boilerplate/repository"
	"gin-boilerplate/service/api"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func UserServiceHandler() *api.UserHandler {
	service := api.UserHandler{
		UserRepo: &repository.DBUserRepositoty{},
	}
	return &service
}

func init() {

	// // Set Gin to production mode
	// gin.SetMode(gin.ReleaseMode)

	Router = gin.Default()

	r := Router.Group("/api")

	{
		r.GET("/ping", api.Ping)

	}
	u := Router.Group("/users")
	{

		u.POST("/", UserServiceHandler().Createuser)
		u.GET("/:id", UserServiceHandler().GetUser)
		u.DELETE("/:id", UserServiceHandler().DeleteUser)
		u.PUT("/:id", UserServiceHandler().UpdateUser)

	}
}
