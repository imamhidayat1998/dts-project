package configs

import (
	"log"
	"dts-project/pkg/database"
	"os"

	"github.com/joho/godotenv"
)
type Config struct {
}

func (c *Config) InitEnv() error {
	err := godotenv.Load()
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
