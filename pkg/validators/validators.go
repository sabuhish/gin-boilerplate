package validators

import (
	"github.com/thedevsaddam/govalidator"
)

type UserValidator struct {
	Name        string `json:"name" binding:"required"`
	Age         int    `json:"age" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}



type UserUriValidator struct {
	ID  string `uri:"id" binding:"required,uuid"`

}



func (u *UserValidator) Validate() govalidator.MapData {

	userrules := govalidator.MapData{
		"name":         []string{"required"},
		"age":          []string{"required"},
		"username":     []string{"required"},
		"email":        []string{"required", "email"},
		"phone_number": []string{"required"},
	}
	return userrules

}
