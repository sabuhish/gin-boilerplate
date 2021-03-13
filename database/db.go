package database

import (
	"fmt"
	"gin-boilerplate/configs"
	"gin-boilerplate/models"
	"gin-boilerplate/pkg/logging"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB connection instance
var (
	DB  *gorm.DB
	err error
)

// initialiazies DB connection
func InitDB() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	DB, err = gorm.Open(postgres.Open(configs.Config.DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	logging.Logger.Info(fmt.Sprintf("Connected to  postgres://%s:**********@%s:%s/%s", configs.Config.DBUser, configs.Config.DBHost, configs.Config.DBPort, configs.Config.DBName))

}

func CloseDB(instance *gorm.DB) {

	db, err := instance.DB()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	if err != nil {
		panic("Failed to close database connection")
	}
	db.Close()

}

func Migrate() {

	var models = []interface{}{&models.User{}}

	err := DB.AutoMigrate(models...)

	if err != nil {
		panic(err)

	}
}
