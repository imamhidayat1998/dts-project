package repository

import (
	"database/sql"
	"dts-project/pkg/helpers"
	"dts-project/pkg/model"
	"dts-project/pkg/model/repository/queries"
	"dts-project/pkg/model/request"
	"errors"
	"fmt"

	"github.com/google/uuid"
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
	DataResult, err := Result.RowsAffected()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("CreateUser Result : %d", DataResult), err
}

func (d Repository) UpdateUser(params request.User) error {
	query, err := d.db.Query(queries.QueryUpdateUser, params.Fullname, params.Password, params.Id)
	if err != nil {
		fmt.Sprintf("Detail Error in Queries Params : %s", err)
		return err
	}
	defer query.Close()
	return err
}
func (d Repository) SelectId(params request.User) (*request.User, error) {

	//	query := `select grapari_id,queue_number,grapari_counter_number,queue_status,serving_start_at,agent_userid,agent_name,repeat_call_counter ,queue_delete_at,queue_transit_at ,recall_transit_at from queue where queue_id = ? `
	result, err := d.db.Query(queries.SelectIdUser, params.Fullname, params.Password)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		cek := request.User{}

		err := result.Scan(&cek.Id)

		if err != nil {
			//datalog.Error.Println(err)
			return &request.User{}, err
		}
		return &cek, nil
	}
	return nil, errors.New("queue Number tidak ditemukan")
}
func (d Repository) CreateArtikel(params request.Artikel) (string, error) {
	params.ArtikelId = uuid.NewString()
	GetTime := helpers.GetTimeNow()
	Result, err := d.db.Exec(queries.QueryInsertArtikel,
		params.ArtikelId,
		params.Judul,
		params.Author,
		params.Description,
		GetTime)
	if err != nil {
		fmt.Sprintf("Detail Error in Queries Params : %s", err)
		return "", err
	}
	DataResult, err := Result.RowsAffected()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("CreateUser Result : %d", DataResult), err

}
func (d Repository) GetArtikel() ([]request.Artikel, error) {
	Result, err := d.db.Query(queries.GetArtikel)
	if err != nil {
		fmt.Printf("Detail Error in Queries Params : %s", err)
		return nil, err
	}
	data := []request.Artikel{}
	if Result.Next() {
		var list request.Artikel
		err = Result.Scan(&list.Judul, &list.Author, &list.Description, &list.CreatedAt)
		if err != nil {
			fmt.Printf("Detail Error in Scan Params : %s", err)
			return nil, err
		}
		data = append(data, list)

	}

	err = Result.Err()
	if err != nil {
		fmt.Printf("Detail Error in Rows Params : %s", err)
		return nil, err
	}

	return data, err
}
