package api

import (
	"encoding/json"
	"gin-boilerplate/models"
	"gin-boilerplate/repository"
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

func UserServiceHandler() (*UserHandler, *repository.UserRepoMock) {

	service := UserHandler{
		UserRepo: &repository.UserRepoMock{},
	}
	return &service, &repository.UserRepoMock{}
}

func TestCreate(t *testing.T) {

	service, mockrepo := UserServiceHandler()
	user := models.User{
		Name:        "testName",
		Age:         11,
		Username:    "testusername",
		Email:       "test@gmail.com",
		PhoneNumber: "+994555555555",
	}
	mockrepo.On("Create", &user).Return(nil)

	u, err := service.UserRepo.Create(&user)

	assert.Nil(t, err)
	assert.Equal(t, user.Name, u.Name)
	assert.Equal(t, user.Username, u.Username)
	assert.Equal(t, user.Age, 11)

}
