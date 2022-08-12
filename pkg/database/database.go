package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type ConfigDB struct {
	Driver   string
	DBName   string
	Username string
	Password string
	Host     string
	Port     string
}

func DBConnection(c ConfigDB) (*sql.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.Username, c.Password, c.Host, c.Port, c.DBName)

	db, err := sql.Open(c.Driver, connString)
	if err != nil {
		return nil, err
	}

	return db, err
}
