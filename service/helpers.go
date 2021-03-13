package service

import (
	"gin-boilerplate/pkg/validators"
	"reflect"

	"github.com/gin-gonic/gin"
)

const (
	ErrorMessage   = "error"
	SuccessMessage = "success"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HTTPResponse(c *gin.Context, httpstatus int, message string, data interface{}) {

	c.JSON(httpstatus, Response{
		Code:    httpstatus,
		Message: message,
		Data:    data,
	})

}

func UserData(user interface{}) map[string]interface{} {
	data := make(map[string]interface{})

	switch user.(type) {

	case validators.UserValidator:

		value := reflect.ValueOf(user.(validators.UserValidator))

		typeof := value.Type()

		for i := 0; i < value.NumField(); i++ {
			data[typeof.Field(i).Tag.Get("json")] = value.Field(i).Interface()
		}

	}

	return data
}

func CheckUUID(c *gin.Context, uri *validators.UserUriValidator) error {
	if err := c.ShouldBindUri(&uri); err != nil {

		return err
	}
	return nil

}
