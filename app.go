package main

import (
	"gin-test-go/configs"
	"gin-test-go/db"
	"gin-test-go/routers"
)

func main() {
	//init config
	config := configs.InitConfig()

	//init db
	db.InitDB(config)

	//init router
	router := routers.InitRouter(config)

	//run server
	router.Run(config.Server.Domain + ":" + config.Server.Port)
}
