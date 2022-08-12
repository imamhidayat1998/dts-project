package main

import (
	"dts-project/api/router"
	"dts-project/configs"
	"dts-project/pkg/database"
	"dts-project/pkg/model/repository"
	"dts-project/pkg/model/usecase"
)

func main() {
	start()

}
func start() {
	app := configs.Config{}
	app.CatchError(app.InitEnv())
	url := app.GetURLProject()
	dbConfig := app.GetDBConfig()
	mysql, err := database.DBConnection(dbConfig)
	if err != nil {
		return
	}


	DataRepositorys := repository.NewDataRepository(mysql)
	serviceCenterUsecase := usecase.NewServiceCenter(DataRepositorys)

	router.Router(url,serviceCenterUsecase)

}
