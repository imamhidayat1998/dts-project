package repository

import (
	"database/sql"
	"dts-project/pkg/model"
	"dts-project/pkg/model/repository/queries"
	"dts-project/pkg/model/request"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewDataRepository(db *sql.DB) model.DataRepository {
	return Repository{
		db: db,
	}
}

func (d Repository) CreateUser(params request.User) (string, error) {
	Result, err := d.db.Exec(queries.CreateUser, params.Id, params.Fullname, params.Password)
	if err != nil {
		fmt.Sprintf("Detail Error in Queries Params : %s", err)
		return "", err
	}

	return fmt.Sprintf("CreateUser Result : %s", Result), err
}

func (d Repository) UpdateUser(params request.User) error {
	query, err := d.db.Query(queries.CreateUser, params.Id, params.Fullname, params.Password)
	if err != nil {
		fmt.Sprintf("Detail Error in Queries Params : %s", err)
		return err
	}
	defer query.Close()
	return err
}
