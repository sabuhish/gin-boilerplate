

## Gin-BoilerPlate

This simple ready development and production microservice boilerplate is designed for gin framework. Throught this library I have used, go-arg for argument parsing and getting variables from env, logrus for logging operation either within json or txt format (also can write to the file as well), gorm for database operations, testify for testing.

Features:
swagger documentation will be added, testing will be enhanced.


Structure of the App

```bash

├── configs
│   └── config.go
├── database
│   └── db.go
├── Dockerfile
├── docs
├── go.mod
├── go.sum
├── logging.txt
├── main.go
├── middleware
│   └── middleware.go
├── models
│   └── models.go
├── pkg
│   ├── errors
│   ├── logging
│   │   └── logging.go
│   └── validators
│       └── validators.go
├── README.md
├── repository
│   ├── userrepo.go
│   └── userrepo_mock.go
├── routers
│   └── router.go
├── scripts
├── service
│   ├── api
│   │   ├── user.go
│   │   ├── user_mock.go
│   │   └── user_test.go
│   └── helpers.go


14 directories, 20 files
```


## 🌱 Contributing
If you want to contribute and some features feel free to open an issue and send a pull request.
