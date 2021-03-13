package main

import (
	"fmt"
	"gin-boilerplate/configs"
	"gin-boilerplate/pkg/logging"
	"gin-boilerplate/routers"

	"gin-boilerplate/database"
)

func main() {
	database.InitDB()
	defer database.CloseDB(database.DB)
	database.Migrate()

	logging.Logger.Info(fmt.Sprintf("Starting server at %s:%s", configs.Config.Host, configs.Config.Port))

	routers.Router.Run(configs.Config.Host + ":" + configs.Config.Port)

}
