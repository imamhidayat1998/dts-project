package configs

import (
	"dts-project/pkg/database"
	"dts-project/pkg/model/request"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
}

func (c *Config) InitEnv() error {
	err := godotenv.Load("configs/data.env")
	if err != nil {
		log.Println("failed while load env: ", err)
		return err
	}

	return err
}

func (c *Config) GetDBConfig() database.ConfigDB {
	return database.ConfigDB{
		Driver:   os.Getenv("DB_DRIVER"),
		DBName:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PWD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
}
func (c *Config) CatchError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(any(err))
	}
}
func (c *Config) GetURLProject() request.URL {
	return request.URL{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}

}
