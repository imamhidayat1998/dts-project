package main

import (
	"dts-project/api/router"
	"dts-project/configs"
	"dts-project/pkg/database"
	"goo/pkg/model/repository"
	"kantor/ms-queue-go/service/usecase"
)

func main() {
	Start()

}
func Start()  {
	app := configs.Config{}
	app.CatchError(app.InitEnv())
	dbConfig := app.GetDBConfig()
	mysql, err := database.DBConnection(dbConfig)
	if err != nil {
		return
	}

	DataRepositorys := repository.NewDataRepository(mysql)
	serviceCenterUsecase := usecase.NewServiceCenter(DataRepositorys)
	router.Router(serviceCenterUsecase)

}