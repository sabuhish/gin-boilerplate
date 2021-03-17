package api

import (
	"encoding/json"
	"fmt"
	"gin-boilerplate/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Request(r http.Handler, method, path string) *httptest.ResponseRecorder {

	request, _ := http.NewRequest(method, path, nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	return w
}

func TestPingsRoute(t *testing.T) {
	body := gin.H{
		"hello": "world",
	}

	router := gin.Default()

	router.GET("/api/ping", Ping)

	// Make HTTP Request to the testing endpoint
	w := Request(router, http.MethodGet, "/api/ping")

	var response map[string]string

	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Checking statusCode
	assert.Equal(t, http.StatusOK, w.Code)

	value, exists := response["hello"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["hello"], value)
}

func TestCreate(t *testing.T) {
	mockrepo := new(MockUserRepository)
	service := UserHandler{UserRepo: mockrepo}

	data := models.User{
		Name:        "testName",
		Age:         11,
		Username:    "testusername",
		Email:       "test@gmail.com",
		PhoneNumber: "+994555555555",
	}

	user := models.User{
		Name:        "testName",
		Age:         11,
		Username:    "testusername",
		Email:       "test@gmail.com",
		PhoneNumber: "+994555555555",
	}

	mockrepo.On("Create", &user).Return(&data, nil)

	m, err := service.Createuser(user)

	mockrepo.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, user.Name, m.Name)
	assert.Equal(t, user.Username, m.Username)
	assert.Equal(t, user.Age, 11)
}

func TestGetOne(t *testing.T) {
	mockrepo := new(MockUserRepository)
	service := UserHandler{UserRepo: mockrepo}

	user := models.User{
		Name:        "testName",
		Age:         11,
		Username:    "testusername",
		Email:       "test@gmail.com",
		PhoneNumber: "+994555555555",
	}

	mockrepo.On("GetOne", &user, user.ID.String()).Return(nil)

	data, err := service.GetUser(&user, user.ID.String())

	assert.Equal(t, user.Name, data.Name)
	assert.Equal(t, user.Username, data.Username)
	assert.Equal(t, err, nil)
}

func TestDelete(t *testing.T) {

	mockrepo := new(MockUserRepository)

	service := UserHandler{UserRepo: mockrepo}

	user := models.User{
		Name:        "testName",
		Age:         11,
		Username:    "testusername",
		Email:       "test@gmail.com",
		PhoneNumber: "+994555555555",
	}

	mockrepo.On("Delete", &user, user.ID.String()).Return(nil)

	err := service.DeleteUser(&user, user.ID.String())

	assert.Equal(t, err, nil)

}

func TestUpdate(t *testing.T) {

	mockrepo := new(MockUserRepository)
	service := UserHandler{UserRepo: mockrepo}
	fmt.Println(service)

	user := models.User{
		Name:        "testName",
		Age:         11,
		Username:    "testusername",
		Email:       "test@gmail.com",
		PhoneNumber: "+994555555555",
	}
	updated_user := models.User{
		Name:        "updated_testName",
		Age:         11,
		Username:    "updated_testusername",
		Email:       "updated_test@gmail.com",
		PhoneNumber: "555555555",
	}
	mockrepo.On("GetOne", &user, user.ID.String()).Return(nil)

	data, _ := service.GetUser(&user, user.ID.String())

	assert.Equal(t, user.Name, data.Name)

	mockrepo.On("Update", &updated_user, user.ID.String()).Return(nil)

	m, err := service.UpdateUser(&updated_user, user.ID.String())

	assert.Equal(t, updated_user.Name, m.Name)
	assert.Equal(t, updated_user.Username, m.Username)
	assert.Equal(t, err, nil)
}
