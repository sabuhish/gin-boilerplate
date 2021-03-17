package repository

// type UserRepoMock struct {
// 	mock.Mock
// }

// func (m *UserRepoMock) Create(u *models.User) (*models.User, error) {

// 	ret := m.Called(u)

// 	var r0 *models.User

// 	if ret.Get(0) != nil {
// 		r0 = ret.Get(0).(*models.User)
// 	}
// 	var r1 error

// 	if ret.Get(1) != nil {
// 		r1 = ret.Get(1).(error)
// 	}

// 	return r0, r1

// }
// func (m *UserRepoMock) GetOne(u *models.User, id string) error {

// 	return nil
// }

// func (m *UserRepoMock) Delete(u *models.User, id string) error {

// 	return nil

// }

// func (m *UserRepoMock) Update(u *models.User, id string) error {

// 	return nil
// }

// func (m *MockUserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
//     // args that will be passed to "Return" in the tests, when function
//     // is called with a uid. Hence the name "ret"
//   	ret := m.Called(ctx, uid)

//     // first value passed to "Return"
//     var r0 *model.User
//     if ret.Get(0) != nil {
//         // we can just return this if we know we won't be passing function to "Return"
//         r0 = ret.Get(0).(*model.User)
//     }

//     var r1 error

//     if ret.Get(1) != nil {
//         r1 = ret.Get(1).(error)
//     }

//     return r0, r1

// func TestMe(t *testing.T) {
//     // Setup
//     gin.SetMode(gin.TestMode)

//     t.Run("Success", func(t *testing.T) {
//         uid, _ := uuid.NewRandom()

//         mockUserResp := &model.User{
//             UID:   uid,
//             Email: "bob@bob.com",
//             Name:  "Bobby Bobson",
//         }

//         mockUserService := new(mocks.MockUserService)
//         mockUserService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(mockUserResp, nil)

//         // a response recorder for getting written http response
//         rr := httptest.NewRecorder()

//         // use a middleware to set context for test
//         // the only claims we care about in this test
//         // is the UID
//         router := gin.Default()
//         router.Use(func(c *gin.Context) {
//             c.Set("user", &model.User{
//                 UID: uid,
//             },
//             )
//         })

//         NewHandler(&Config{
//             R:           router,
//             UserService: mockUserService,
//         })

//         request, err := http.NewRequest(http.MethodGet, "/me", nil)
//         assert.NoError(t, err)

//         router.ServeHTTP(rr, request)

//         respBody, err := json.Marshal(gin.H{
//             "user": mockUserResp,
//         })
//         assert.NoError(t, err)

//         assert.Equal(t, 200, rr.Code)
//         assert.Equal(t, respBody, rr.Body.Bytes())
//         mockUserService.AssertExpectations(t) // assert that UserService.Get was called
//     })

//     t.Run("NoContextUser", func(t *testing.T) {
//         mockUserService := new(mocks.MockUserService)
//         mockUserService.On("Get", mock.Anything, mock.Anything).Return(nil, nil)

//         // a response recorder for getting written http response
//         rr := httptest.NewRecorder()

//         // do not append user to context
//         router := gin.Default()
//         NewHandler(&Config{
//             R:           router,
//             UserService: mockUserService,
//         })

//         request, err := http.NewRequest(http.MethodGet, "/me", nil)
//         assert.NoError(t, err)

//         router.ServeHTTP(rr, request)

//         assert.Equal(t, 500, rr.Code)
//         mockUserService.AssertNotCalled(t, "Get", mock.Anything)
//     })

//     t.Run("NotFound", func(t *testing.T) {
//         uid, _ := uuid.NewRandom()
//         mockUserService := new(mocks.MockUserService)
//         mockUserService.On("Get", mock.Anything, uid).Return(nil, fmt.Errorf("Some error down call chain"))

//         // a response recorder for getting written http response
//         rr := httptest.NewRecorder()

//         router := gin.Default()
//         router.Use(func(c *gin.Context) {
//             c.Set("user", &model.User{
//                 UID: uid,
//             },
//             )
//         })

//         NewHandler(&Config{
//             R:           router,
//             UserService: mockUserService,
//         })

//         request, err := http.NewRequest(http.MethodGet, "/me", nil)
//         assert.NoError(t, err)

//         router.ServeHTTP(rr, request)

//         respErr := apperrors.NewNotFound("user", uid.String())

//         respBody, err := json.Marshal(gin.H{
//             "error": respErr,
//         })
//         assert.NoError(t, err)

//         assert.Equal(t, respErr.Status(), rr.Code)
//         assert.Equal(t, respBody, rr.Body.Bytes())
//         mockUserService.AssertExpectations(t) // assert that UserService.Get was called
//     })
// }
